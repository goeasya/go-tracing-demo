package ginx

import (
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
)

func Opentracing(tr opentracing.Tracer) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, _ := tr.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(c.Request.Header))
		sp := tr.StartSpan(c.Request.URL.Path, opentracing.ChildOf(ctx))
		c.Request = c.Request.WithContext(opentracing.ContextWithSpan(c.Request.Context(), sp))
		c.Next()
		sp.Finish()
	}
}
