package main

import (
	"fmt"
	"os"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"go-tracing-demo/global"
	"go-tracing-demo/pkg/ginx"
	"go.elastic.co/apm/module/apmgin"
)


func newServer() *gin.Engine {
	r := gin.Default()
	pprof.Register(r)

	// add elastic apm
	r.Use(apmgin.Middleware(r))

	apiR := r.Group("api")
	{
		eventR := apiR.Group("event")
		{
			eventR.POST("/create", ginx.CreateEvent)
			eventR.GET("list", ginx.ListEvent)
			eventR.GET("/findByEventId", ginx.FindEventByEventId)
		}
	}
	return r
}


func main() {
	dbcfg := &global.DBConfig{Type: "gorm", ConnInfo: "xxxx:xxxx@tcp(10.1.1.1:3306)/trace_test"}
	if err := global.MustInitDBManager(dbcfg); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	server := newServer()
	if err := server.Run(":8080"); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
