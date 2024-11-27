package config

import (
	"os"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

func ConnectMidtrans() *snap.Client {
	var client snap.Client

	serverKey := os.Getenv("MIDTRANS_SERVER_KEY")

	client.New(serverKey, midtrans.Sandbox)

	return &client
}