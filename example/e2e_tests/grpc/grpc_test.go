package grpc

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestGRPCService(t *testing.T) {
	//Arrange
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatal("failed to connect to gRPC server:", err)
	}
	defer conn.Close()

	client := NewExampleClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	//Act
	resp, err := client.Register(ctx, &RegisterRequest{
		Username: "testuser",
		Password: "test12345",
	})

	//Assert
	require.NoError(t, err, "failed to register user:", err)
	assert.NotEmpty(t, resp.Success, "user registered successfully")
}
