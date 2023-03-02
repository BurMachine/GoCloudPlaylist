package main

import (
	api "GoCloudPlaylist/pkg/api"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/metadata"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	conn, err := grpc.Dial(
		"localhost:9090",
		grpc.WithInsecure(),
		grpc.WithKeepaliveParams(keepalive.ClientParameters{
			Time:                10 * time.Second,
			Timeout:             time.Second,
			PermitWithoutStream: true,
		}),
	)
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	client := api.NewGoCloudPlaylistClient(conn)
	req := &api.Empty{}
	ctx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs("key", "value")) // можно добавить метаданные
	resp, err := client.PlaySong(ctx, req)
	if err != nil {
		log.Fatalf("Failed to get response: %v", err)
	}
	log.Printf("Response(grpc): %v\n", resp)

	// http

	clientHttp := &http.Client{Timeout: 5 * time.Second}

	respHttp, err := clientHttp.Get("http://localhost:8080/pause")
	if err != nil {
		log.Fatalf("Failed to get response: %v", err)
	}
	defer respHttp.Body.Close()

	body, err := io.ReadAll(respHttp.Body)
	if err != nil {
		log.Fatalf("Failed to get response: %v", err)
	}
	log.Printf("\n\n\nResponse(http): %v\n", string(body))

	// Добавление
	addReq := api.AddRequest{
		Name: "Song2 - Blur",
		Time: "00:00:12",
	}
	respAdd, err := client.AddSong(ctx, &addReq)
	if err != nil {
		log.Fatalf("Failed to get response: %v", err)
	}
	log.Printf("Response(grpc): %v\n", respAdd)

	// Удаление
	respHttpDel, err := clientHttp.Get("http://localhost:8080/pause")
	if err != nil {
		log.Fatalf("Failed to get response: %v", err)
	}
	defer respHttpDel.Body.Close()
	log.Printf("\nResponse(http): %v\n", respHttpDel.StatusCode)
}
