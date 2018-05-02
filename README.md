# gRPC echo service

## 1. Proto file
Service definition stored in `api/echo_service.proto` file.
To re-generate golang code from this proto, run `scripts/bootstrap.sh` and `scripts/generate_grpc_code.sh`

## 2. Run server
`go run cmd/main.go`


## 3. Run client
`go run scripts/test_client.go 'ping-pong king-kong'`
