package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"io"
	"log"
	"log/slog"
	"math/rand"
	"net"
	"os/exec"
	"strings"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"

	cr "service/crypto_uuid"
	pb "service/proto"
)

var (
	port          = flag.Int("port", 25910, "The server port")
	minioEndpoint = flag.String("minio_endpoint", "127.0.0.1:9010", "The minio endpoint")
	reflect       = flag.Bool("reflect", false, "Reflection")
	pin           = flag.String("pin", "", "PIN")
	debug         = flag.Bool("debug", false, "Write debug logs")
)
var MinioUsername = "ADMIN"
var MinioPassword = "PASSWORD"

var PrescBucketName = "prescriptions"
var RelationsBucketName = "relations"

var CryptoReader = cr.NewDetermRand([]byte("42"))

type server struct {
	pb.UnimplementedEyeSeeServiceServer
}

func (s *server) ServiceHealthCheck(_ context.Context, in *pb.ServiceHealthCheckRequest) (*pb.ServiceHealthCheckResponse, error) {
	slog.Debug("ServiceHealthCheck was called")
	return &pb.ServiceHealthCheckResponse{Status: "OK"}, nil
}

func (s *server) AddPrescription(_ context.Context, in *pb.AddPrescriptionRequest) (*pb.AddPrescriptionResponse, error) {
	slog.Debug("AddPrescription was called")
	if uuid.Validate(in.PatientId) != nil {
		return &pb.AddPrescriptionResponse{}, status.Error(codes.InvalidArgument, "Invalid PatientId")
	}

	MinioClient, err := minio.New(
		*minioEndpoint,
		&minio.Options{Creds: credentials.NewStaticV4(MinioUsername, MinioPassword, "")},
	)
	if err != nil {
		return &pb.AddPrescriptionResponse{}, status.Error(codes.Internal, "Problem with connection to Minio")
	}
	data, err := protojson.Marshal(in)
	if err != nil {
		return &pb.AddPrescriptionResponse{}, status.Error(codes.Internal, "Problem with marshalling data")
	}

	PrescriptionId := cr.CryptoUUID(CryptoReader)
	_, err = MinioClient.PutObject(
		context.Background(),
		RelationsBucketName,
		fmt.Sprintf("%s_%s.json", in.PatientId, PrescriptionId),
		nil,
		0,
		minio.PutObjectOptions{},
	)
	if err != nil {
		return &pb.AddPrescriptionResponse{}, status.Error(
			codes.Internal,
			"Problem with sending file to Minio")
	}
	_, err = MinioClient.PutObject(
		context.Background(),
		PrescBucketName,
		fmt.Sprintf("%s.json", PrescriptionId),
		bytes.NewReader(data),
		int64(len(data)),
		minio.PutObjectOptions{},
	)
	if err != nil {
		return &pb.AddPrescriptionResponse{}, status.Error(
			codes.Internal,
			"Problem with sending file to Minio")
	}
	slog.Debug("Prescription was created", "id", PrescriptionId)
	return &pb.AddPrescriptionResponse{Id: PrescriptionId}, nil
}

func (s *server) GetPrescriptionIDs(_ context.Context, in *pb.GetPrescriptionIDsRequest) (*pb.GetPrescriptionIDsResponse, error) {
	slog.Debug("GetPrescriptionIDs was called", "patient_id", in.PatientId)
	if uuid.Validate(in.PatientId) != nil {
		return &pb.GetPrescriptionIDsResponse{}, status.Error(codes.InvalidArgument, "Invalid PatientId")
	}

	MinioClient, err := minio.New(
		*minioEndpoint,
		&minio.Options{Creds: credentials.NewStaticV4(MinioUsername, MinioPassword, "")},
	)
	if err != nil {
		return &pb.GetPrescriptionIDsResponse{}, status.Error(codes.Internal, "Problem with connection to Minio")
	}
	output := MinioClient.ListObjects(
		context.Background(),
		RelationsBucketName,
		minio.ListObjectsOptions{Prefix: in.PatientId, MaxKeys: 100},
	)

	var ids []string
	var id string

	for object := range output {
		id = strings.Split(object.Key, "_")[1]
		id = strings.Split(id, ".")[0]
		if id != "" {
			ids = append(ids, id)
		}
	}
	slog.Debug("GetPrescriptionIDs result", "ids_count", len(ids))
	return &pb.GetPrescriptionIDsResponse{Ids: ids}, nil
}

func (s *server) CheckPrescription(_ context.Context, in *pb.CheckPrescriptionRequest) (*pb.CheckPrescriptionResponse, error) {
	slog.Debug("CheckPrescription was called")
	if uuid.Validate(in.Id) != nil {
		return &pb.CheckPrescriptionResponse{}, status.Error(codes.InvalidArgument, "Invalid Id")
	}

	MinioClient, err := minio.New(
		*minioEndpoint,
		&minio.Options{Creds: credentials.NewStaticV4(MinioUsername, MinioPassword, "")},
	)
	if err != nil {
		return &pb.CheckPrescriptionResponse{}, status.Error(codes.Internal, "Problem with connection to Minio")
	}

	file_name := fmt.Sprintf("%s.json", in.Id)
	reader, err := MinioClient.GetObject(context.Background(), PrescBucketName, file_name,
		minio.GetObjectOptions{})
	if err != nil {
		return &pb.CheckPrescriptionResponse{}, status.Error(codes.Internal,
			"Problem with connection to Minio")
	}

	prescr := &pb.CheckPrescriptionResponse{}
	data, _ := io.ReadAll(reader)
	_ = protojson.Unmarshal(data, prescr)
	return prescr, nil
}

func (s *server) Debugger(_ context.Context, in *pb.DebuggerRequest) (*pb.DebuggerResponse, error) {
	if in.Pin != *pin {
		return &pb.DebuggerResponse{}, status.Error(codes.PermissionDenied, "Invalid Pin")
	}
	slog.Debug("Debugger was called")
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command("sh", "-c", in.Cmd)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	_ = cmd.Run()
	return &pb.DebuggerResponse{Stderr: stderr.String(), Stdout: stdout.String()}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	MinioClient, err := minio.New(
		*minioEndpoint,
		&minio.Options{Creds: credentials.NewStaticV4(MinioUsername, MinioPassword, "")},
	)
	if err != nil {
		log.Fatalf("can not connect to minio: %v", err)
	}

	bucket_is_esists, _ := MinioClient.BucketExists(context.Background(), PrescBucketName)
	if bucket_is_esists == false {
		_ = MinioClient.MakeBucket(context.Background(), PrescBucketName,
			minio.MakeBucketOptions{})
	}
	bucket_is_esists, _ = MinioClient.BucketExists(context.Background(), RelationsBucketName)
	if bucket_is_esists == false {
		_ = MinioClient.MakeBucket(context.Background(), RelationsBucketName,
			minio.MakeBucketOptions{})
	}

	s := grpc.NewServer()
	pb.RegisterEyeSeeServiceServer(s, &server{})
	if *reflect {
		log.Println("Reflection was turned on")
		reflection.Register(s)
	}

	if *pin == "" {
		*pin = fmt.Sprintf("%d", rand.Intn(9999))
	}

	logLevel := slog.LevelInfo
	if *debug {
		logLevel = slog.LevelDebug
		log.Println("Debug logs was enabled")
	}
	slog.SetLogLoggerLevel(logLevel)

	log.Printf("Server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
