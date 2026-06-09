package main

import (
	"github.com/gin-gonic/gin"
	"go-study/routers"
)

// go run github.com/pilu/fresh
func main() {
	router := gin.Default()
	routers.AdminRoutersInit(router)
	router.Run() // listens on 0.0.0.0:8080 by default

}

//  goctl rpc -o greet.proto
//  goctl rpc protoc greet.proto --go_out=./grpc-server --go-grpc_out=./grpc-server --zrpc_out=./grpc-server
