package controller

import (
	"context"
	"encoding/json"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"math/big"
	"testXDPoS/config"
	"testXDPoS/utils"
)

type EthController struct {
	beego.Controller
}

type EthBlockReq struct {
	BlockNum int64  `json:"blockNum"`
}

func (c *EthController) GetBlockReceipts() {
	req := &EthBlockReq{}
	json.Unmarshal(c.Ctx.Input.RequestBody, req)
	fmt.Printf("req:%+v",req)
	blockNum := big.NewInt(req.BlockNum)
	receipts, err := config.EthClient.GetBlockReceipts(context.Background(), blockNum)
	if err != nil{
		utils.Error(c.Ctx, 400, err.Error())
		return
	}
	fmt.Printf("res:%v",receipts)
	utils.Success(c.Ctx, receipts, 0)
	return
}