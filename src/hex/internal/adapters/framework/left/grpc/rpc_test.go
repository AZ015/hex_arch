package rpc

import (
	"context"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"hex/internal/adapters/app/api"
	"hex/internal/adapters/core/arithmetic"
	"hex/internal/adapters/framework/left/grpc/pb"
	"hex/internal/adapters/framework/right/db"
	"hex/internal/ports"
	"log"
	"net"
	"os"
	"testing"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {
	lis = bufconn.Listen(bufSize)
	grpcServer := grpc.NewServer()

	// ports
	var dbasAdapter ports.DBPort
	var core ports.ArithmeticPort
	var appAdapter ports.APIPort
	var gRPCAdapter ports.GRPCPort

	dbaseDriver := os.Getenv("DB_DRIVER")
	dsourceName := os.Getenv("DS_NAME")

	dbasAdapter, err := db.NewAdapter(dbaseDriver, dsourceName)
	if err != nil {
		log.Fatalf("failed to initiate dbase connection: %v", err)
	}

	core = arithmetic.NewAdapter()

	appAdapter = api.NewAdapter(dbasAdapter, core)

	gRPCAdapter = NewAdapter(appAdapter)

	pb.RegisterArithmeticServiceServer(grpcServer, gRPCAdapter)

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("test server start error: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func getGRPCConnection(ctx context.Context, t *testing.T) *grpc.ClientConn {
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial bufnet: %v", err)
	}

	return conn
}

func TestAdapter_GetAddition(t *testing.T) {
	ctx := context.Background()
	conn := getGRPCConnection(ctx, t)
	defer conn.Close()
	client := pb.NewArithmeticServiceClient(conn)

	tests := []struct {
		name    string
		req     *pb.OperationParameters
		want    *pb.Answer
		wantErr bool
	}{
		{
			"get addition success",
			&pb.OperationParameters{
				A: 1,
				B: 2,
			},
			&pb.Answer{
				Value: 3,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := client.GetAddition(ctx, tt.req)
			if (err != nil) != tt.wantErr {
				require.Error(t, err)
			}
			require.Equal(t, got.Value, tt.want.Value)
		})
	}
}

func TestAdapter_GetSubtraction(t *testing.T) {
	ctx := context.Background()
	conn := getGRPCConnection(ctx, t)
	defer conn.Close()
	client := pb.NewArithmeticServiceClient(conn)

	tests := []struct {
		name    string
		req     *pb.OperationParameters
		want    *pb.Answer
		wantErr bool
	}{
		{
			"get subtraction success",
			&pb.OperationParameters{
				A: 10,
				B: 5,
			},
			&pb.Answer{
				Value: 5,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := client.GetSubtraction(ctx, tt.req)
			if (err != nil) != tt.wantErr {
				require.Error(t, err)
			}
			require.Equal(t, got.Value, tt.want.Value)
		})
	}
}

func TestAdapter_GetMultiplication(t *testing.T) {
	ctx := context.Background()
	conn := getGRPCConnection(ctx, t)
	defer conn.Close()
	client := pb.NewArithmeticServiceClient(conn)

	tests := []struct {
		name    string
		req     *pb.OperationParameters
		want    *pb.Answer
		wantErr bool
	}{
		{
			"get multiplication success",
			&pb.OperationParameters{
				A: 2,
				B: 2,
			},
			&pb.Answer{
				Value: 4,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := client.GetMultiplication(ctx, tt.req)
			if (err != nil) != tt.wantErr {
				require.Error(t, err)
			}
			require.Equal(t, got.Value, tt.want.Value)
		})
	}
}

func TestAdapter_GetDivision(t *testing.T) {
	ctx := context.Background()
	conn := getGRPCConnection(ctx, t)
	defer conn.Close()
	client := pb.NewArithmeticServiceClient(conn)

	tests := []struct {
		name    string
		req     *pb.OperationParameters
		want    *pb.Answer
		wantErr bool
	}{
		{
			"get division success",
			&pb.OperationParameters{
				A: 6,
				B: 3,
			},
			&pb.Answer{
				Value: 2,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := client.GetDivision(ctx, tt.req)
			if (err != nil) != tt.wantErr {
				require.Error(t, err)
			}
			require.Equal(t, got.Value, tt.want.Value)
		})
	}
}
