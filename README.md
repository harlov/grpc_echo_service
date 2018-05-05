# gRPC echo service

## 1. Proto file
Service definition stored in `api/echo_service.proto` file.
To re-generate golang code from this proto, run `scripts/bootstrap.sh` and `scripts/generate_grpc_code.sh`

## 2. Run server
`go run cmd/main.go server`


## 3. Run client
`go run cmd/main.go client client_1 'ping-pong king-kong'`
