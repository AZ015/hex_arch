package main

import (
	"hex/internal/adapters/app/api"
	"hex/internal/adapters/core/arithmetic"
	gRPC "hex/internal/adapters/framework/left/grpc"
	"hex/internal/adapters/framework/right/db"
	"hex/internal/ports"
	"log"
	"os"
)

func main() {
	var err error

	// ports
	var dbasAdapter ports.DBPort
	var core ports.ArithmeticPort
	var appAdapter ports.APIPort
	var gRPCAdapter ports.GRPCPort

	dbaseDriver := os.Getenv("DB_DRIVER")
	dsourceName := os.Getenv("DS_NAME")

	dbasAdapter, err = db.NewAdapter(dbaseDriver, dsourceName)
	if err != nil {
		log.Fatalf("failed to initiate dbase connection: %v", err)
	}
	defer dbasAdapter.CloseDBConnection()

	core = arithmetic.NewAdapter()

	appAdapter = api.NewAdapter(dbasAdapter, core)

	gRPCAdapter = gRPC.NewAdapter(appAdapter)
	gRPCAdapter.Run()
}
