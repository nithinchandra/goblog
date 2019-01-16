package main

import(
	"fmt"
	"github.com/nithinchandra/goblog/accountservice/dbclient"
	"github.com/nithinchandra/goblog/accountservice/service"
)

var appName = "accountservice"

func main(){
	fmt.Printf("Starting %v\n",appName)
	initializeBoltClient()
	service.StartWebServer("6767")
}

// Creates instance and calls the OpenBoltDb and Seed funcs
func initializeBoltClient(){
	service.DBClient = &dbclient.BoltClient{}
	service.DBClient.OpenBoltDb()
	service.DBClient.Seed()

}