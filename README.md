1. pkg install git
2. git clone https://github.com/fauziii1/0g-auto-upload.git
3. cd 0g-auto-upload
4. pkg install golang
5. go mod init 0g-auto-upload
6. go get github.com/0glabs/0g-storage-client
7. touch file.txt
8. nano file.txt (isi file bebas untuk upload contoh :test) ctrl+x untuk keluar dan save
9. pilih jenis upload (1x atau berulang kali upload)
10. go run -ldflags=-checklinkname=0 fmain.go (1x upload)
11. go run -ldflags=-checklinkname=0 loopmain.go (berulang kali upload) ctrl+c untuk berhenti
