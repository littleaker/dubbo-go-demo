package main

import (
	"context"
	"dubbo.apache.org/dubbo-go/v3/common/logger"
	"dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
	"go-dubbo-demo/api"
)

func main(){
	if err := config.Load(); err != nil{
		panic(err)
	}
	var i int32 = 1
	//发起调用
	user,err := api.UserProviderClient.GetUser(context.TODO(), i)
	if err != nil{
		panic(err)
	}
	logger.Infof("response result: %+v", user)
}
