package api

import (
	"go-tracing-demo/global"
	"go-tracing-demo/pkg/httpx"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/module/apmgin"
)


func NewServer() *gin.Engine {
	r := gin.Default()
	pprof.Register(r)

	// add elastic apm
	r.Use(apmgin.Middleware(r))

	apiR := r.Group("api")
	{
		eventR := apiR.Group("event")
		{
			eventR.POST("/create", CreateEvent)
			eventR.GET("list", ListEvent)
			eventR.GET("/findByEventId", FindEventByEventId)
		}
	}
	return r
}

func CreateEvent(c *gin.Context) {
	err := global.GetDBManager().EventDao().NewEvent(c.Request.Context(), "")
	if err != nil {
		httpx.GinResponseError(c, httpx.EventErrCode, err.Error())
		return
	}
	httpx.GinResponseSuccess(c, httpx.SuccessCode, nil)
}

func ListEvent(c *gin.Context) {
	events, err := global.GetDBManager().EventDao().List(c.Request.Context())
	if err != nil {
		httpx.GinResponseError(c, httpx.EventErrCode, err.Error())
		return
	}
	httpx.GinResponseSuccess(c, httpx.SuccessCode, events)
}

func FindEventByEventId(c *gin.Context) {
	eventId := c.Query("event_id")
	if len(eventId) == 0 {
		httpx.GinResponseError(c, httpx.EventErrCode, "invalid param")
		return
	}
	event, err := global.GetDBManager().EventDao().GetByEventId(c.Request.Context(), eventId)
	if err != nil {
		httpx.GinResponseError(c, httpx.EventErrCode, err.Error())
		return
	}
	httpx.GinResponseSuccess(c, httpx.SuccessCode, &event)
}


