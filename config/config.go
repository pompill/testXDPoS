package config

import (
	"github.com/XinFinOrg/XDPoSChain/ethclient"
	//"os"
)

var EthClient *ethclient.Client

func init() {
	//netUrl := os.Getenv("ethMainNetUrl")
	netUrl := "https://arbitrum-sepolia.blockpi.network/v1/rpc/public"
	client, err := ethclient.Dial(netUrl)
	if err != nil{
		panic("ethclient regist fail")
	}
	EthClient = client
}
