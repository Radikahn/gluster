dev:
	go build -o gluster ./cmd
	./gluster ls

api:
    go run api/main.go
