package route

import (
	beego "github.com/beego/beego/v2/server/web"
	"testXDPoS/controller"
)

func init(){
	ns := beego.NewNamespace("backend",
			beego.NSNamespace("eth",
				beego.NSRouter("getBlockReceipts", &controller.EthController{}, "Post:GetBlockReceipts"),
		),
	)
	beego.AddNamespace(ns)
}
