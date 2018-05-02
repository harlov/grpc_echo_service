#!/bin/sh

protoc -I ../api ../api/echo_service.proto --go_out=plugins=grpc:../pkg/proto
