package beefy

import (
	"os"
	"testing"

	"github.com/Cerebellum-Network/go-substrate-rpc-client/v8/client"
	"github.com/Cerebellum-Network/go-substrate-rpc-client/v8/config"
)

var testBeefy Beefy

func TestMain(m *testing.M) {
	cl, err := client.Connect(config.Default().RPCURL)
	if err != nil {
		panic(err)
	}
	testBeefy = NewBeefy(cl)
	os.Exit(m.Run())
}
