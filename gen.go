package main

//go:generate go run gen/goa/main.go
//go:generate protoc -I grpc/ grpc/job.proto --go_out=plugins=grpc:grpc/gen
