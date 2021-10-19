package main

import (
	"fmt"
	"os"

	"go-tracing-demo/global"
	"go-tracing-demo/traces/elasticapm/api"
)

func main() {
	dbcfg := &global.DBConfig{Type: "gorm", ConnInfo: "USER:PASSWORD@tcp(127.0.0.1:3306)/trace_test"}
	if err := global.MustInitDBManager(dbcfg); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	server := api.NewServer()
	if err := server.Run(":8080"); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
