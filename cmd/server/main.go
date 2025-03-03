package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"cloud.google.com/go/storage"
	"github.com/gorilla/mux"
	"google.golang.org/api/iterator"
)

type HealthResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type AssetListResponse struct {
	Assets []string `json:"assets"`
}

func main() {
	r := mux.NewRouter()

	// Health check endpoint
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		response := HealthResponse{
			Status:  "ok",
			Message: "DCSS Roguelike service is running",
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}).Methods("GET")

	// Asset list endpoint
	r.HandleFunc("/assets", listAssets).Methods("GET")

	// Get port from environment variable or default to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

// listAssets lists all assets in the GCP bucket
func listAssets(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	// Get bucket name from environment variable
	bucketName := os.Getenv("ASSET_BUCKET")
	if bucketName == "" {
		bucketName = "malathon-roguelike-asset" // Default fallback
	}

	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Printf("Failed to create storage client: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer client.Close()

	// Create a bucket handle
	bucket := client.Bucket(bucketName)

	// List objects in the bucket
	query := &storage.Query{Prefix: ""}
	it := bucket.Objects(ctx, query)

	var assets []string
	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Printf("Error listing bucket objects: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		assets = append(assets, attrs.Name)
	}

	// Return the list of assets
	response := AssetListResponse{
		Assets: assets,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
