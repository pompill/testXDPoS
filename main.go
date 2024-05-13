package main

import (
	beego "github.com/beego/beego/v2/server/web"
	_ "testXDPoS/config"
	_ "testXDPoS/route"
)

func main(){
	beego.Run()
}
