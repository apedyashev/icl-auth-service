package main

import (
	"icl-auth/pkg/adapter/datastore"
	"icl-auth/pkg/adapter/grpc"
)

func main() {
	db := datastore.NewDB()
	datastore.Migrate(db)

	grpc.StartAllGrpcServers(db)
}
