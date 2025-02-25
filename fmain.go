package main

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"github.com/0glabs/0g-storage-client/common/blockchain"
	"github.com/0glabs/0g-storage-client/indexer"
	"github.com/0glabs/0g-storage-client/transfer"
)

func main() {
	fmt.Println("Setting up Go environment...")
	if err := setupGoEnvironment(); err != nil {
		fmt.Println("Failed to set up Go environment:", err)
		return
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Ethereum RPC URL: ")
	evnRpc, _ := reader.ReadString('\n')
	evnRpc = evnRpc[:len(evnRpc)-1]

	fmt.Print("Enter your Ethereum private key: ")
	privateKey, _ := reader.ReadString('\n')
	privateKey = privateKey[:len(privateKey)-1]

	fmt.Print("Enter Indexer RPC URL: ")
	indRpc, _ := reader.ReadString('\n')
	indRpc = indRpc[:len(indRpc)-1]

	fmt.Print("Enter file path to upload: ")
	filePath, _ := reader.ReadString('\n')
	filePath = filePath[:len(filePath)-1]

	ctx := context.Background()

	w3client := blockchain.MustNewWeb3(evnRpc, privateKey)
	defer w3client.Close()

	indexerClient, err := indexer.NewClient(indRpc)
	if err != nil {
		fmt.Println("Failed to create indexer client:", err)
		return
	}

	nodes, err := indexerClient.SelectNodes(ctx, 1, 1, nil)
	if err != nil {
		fmt.Println("Failed to select storage nodes:", err)
		return
	}

	uploader, err := transfer.NewUploader(ctx, w3client, nodes)
	if err != nil {
		fmt.Println("Failed to create uploader:", err)
		return
	}

	txHash, err := uploader.UploadFile(ctx, filePath)
	if err != nil {
		fmt.Println("Failed to upload file:", err)
		return
	}

	fmt.Println("File uploaded successfully! Transaction Hash:", txHash)
}

func setupGoEnvironment() error {
	cmds := []string{
		"go mod init 0g-storage-uploader",
		"go get github.com/0glabs/0g-storage-client",
	}

	for _, cmd := range cmds {
		fmt.Println("Running:", cmd)
		if err := executeCommand(cmd); err != nil {
			return err
		}
	}
	return nil
}

func executeCommand(command string) error {
	cmd := exec.Command("sh", "-c", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
