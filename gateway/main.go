package main

import (
    "context"
    "encoding/json"
    "log"
    "net/http"

    example "github.com/hamidehalaee/proto/github.com/hamidehalaee/proto/example" 
    "google.golang.org/grpc"
)

func main() {
    grpcEndpoint := "localhost:50051" // gRPC service1 address
    httpEndpoint := ":8080"           // RESTful API gateway address

    // Setup gRPC client connection to service1
    conn, err := grpc.DialContext(
        context.Background(),
        grpcEndpoint,
        grpc.WithInsecure(),
        grpc.WithBlock(),
    )
    if err != nil {
        log.Fatalf("Failed to connect to gRPC server: %v", err)
    }
    defer conn.Close()

    // gRPC client
    client := example.NewExampleServiceClient(conn)

    // Custom HTTP handler for the /v1/hello endpoint
    http.HandleFunc("/v1/hello", func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPost {
            http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
            return
        }

        var req example.HelloRequest
        if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
            http.Error(w, "Invalid request payload", http.StatusBadRequest)
            return
        }

        // Call the gRPC SayHello method
        grpcResp, err := client.SayHello(context.Background(), &req)
        if err != nil {
            http.Error(w, "gRPC error: "+err.Error(), http.StatusInternalServerError)
            return
        }

        // Respond with JSON
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(grpcResp)
    })

    // Start the HTTP server for REST requests
    log.Printf("Gateway is listening on %s", httpEndpoint)
    if err := http.ListenAndServe(httpEndpoint, nil); err != nil {
        log.Fatalf("Failed to start HTTP server: %v", err)
    }
}
