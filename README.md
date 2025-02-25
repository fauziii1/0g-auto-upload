Pkg install git
git clone https://github.com/fauziii1/0g-auto-upload.git
pkg install golang
mkdir 0g-storage-test && cd 0g-storage-test
go mod init 0g-auto-upload
go get github.com/0glabs/0g-storage-client
touch file.txt
nano file.txt (isi file bebas untuk upload contoh :test) ctrl+x untuk keluar dan save
go run -ldflags=-checklinkname=0 fmain.go
