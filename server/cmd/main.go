package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/JHeimbach/nfc-cash-system/server"
	"github.com/JHeimbach/nfc-cash-system/server/models/mysql"
	"google.golang.org/grpc"
)

func main() {
	url := flag.String("grpc-host", "localhost", "Host for grpc server")
	port := flag.String("grpc-port", "50051", "Port for grpc server")
	flag.Parse()

	if err := startGrpcServer(*port); err != nil {
		log.Fatal(err)
	}
}

func startGrpcServer(port string) error {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	server.RegisterAccountServer(s, mysql.NewAccountModel(&sql.DB{}))

	if err := s.Serve(lis); err != nil {
		return fmt.Errorf("failed to serve: %v", err)
	}

	return nil
}
