package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/0glabs/0g-storage-client/common/blockchain"
	"github.com/0glabs/0g-storage-client/indexer"
	"github.com/0glabs/0g-storage-client/transfer"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Input manual langsung
	evmRpc := getInput(reader, "Enter Ethereum RPC URL: ")
	privateKey := getInput(reader, "Enter your Ethereum private key: ")
	indRpc := getInput(reader, "Enter Indexer RPC URL: ")
	filePath := getInput(reader, "Enter file path to upload: ")

	// Konteks untuk request
	ctx := context.Background()

	// Inisialisasi Web3 client
	w3client := blockchain.MustNewWeb3(evmRpc, privateKey)
	defer w3client.Close()

	// Inisialisasi Indexer client
	indexerClient, err := indexer.NewClient(indRpc)
	if err != nil {
		log.Fatalf("Failed to create indexer client: %v", err)
	}

	// Memilih node storage
	nodes, err := indexerClient.SelectNodes(ctx, 1, 1, nil)
	if err != nil {
		log.Fatalf("Failed to select storage nodes: %v", err)
	}

	// Membuat uploader
	uploader, err := transfer.NewUploader(ctx, w3client, nodes)
	if err != nil {
		log.Fatalf("Failed to create uploader: %v", err)
	}

	// Upload file
	txHash, _, err := uploader.UploadFile(ctx, filePath)
	if err != nil {
		log.Fatalf("Failed to upload file: %v", err)
	}

	fmt.Println("File uploaded successfully! Transaction Hash:", txHash)
}

// Fungsi untuk membaca input satu baris penuh
func getInput(reader *bufio.Reader, prompt string) string {
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input) // Hapus newline
}
