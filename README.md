1. Pkg install git
2. git clone https://github.com/fauziii1/0g-auto-upload.git
3. pkg install golang
4. mkdir 0g-storage-test && cd 0g-storage-test
5. go mod init 0g-auto-upload
6. go get github.com/0glabs/0g-storage-client
7. touch file.txt
8. nano file.txt (isi file bebas untuk upload contoh :test) ctrl+x untuk keluar dan save
9. go run -ldflags=-checklinkname=0 fmain.go
