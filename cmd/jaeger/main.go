package main

import (
	"fmt"
	"os"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	"go-tracing-demo/global"
	"go-tracing-demo/pkg/ginx"
)

func newServer() *gin.Engine {
	r := gin.Default()
	pprof.Register(r)

	cfg, err := jaegercfg.FromEnv()
	if err != nil {
		// parsing errors might happen here, such as when we get a string where we expect a number
		fmt.Printf("Could not parse Jaeger env vars: %s", err.Error())
		os.Exit(1)
	}

	tracer, closer, err := cfg.NewTracer()
	if err != nil {
		// parsing errors might happen here, such as when we get a string where we expect a number
		fmt.Printf("NewTracer err: %s", err.Error())
		os.Exit(1)
	}
	defer closer.Close()

	r.Use(ginx.Opentracing(tracer))

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
	dbCfg := &global.DBConfig{Type: "gorm", ConnInfo: "xxxx:xxxx@tcp(10.1.1.1:3306)/trace_test"}
	if err := global.MustInitDBManager(dbCfg); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	server := newServer()
	if err := server.Run(":8081"); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
