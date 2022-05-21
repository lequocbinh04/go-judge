package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"go-judge/component"
)

func GinContextToContextMiddleware(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		v := map[string]interface{}{
			"AppContext":    appCtx,
			"GinContextKey": c,
		}
		ctx := context.WithValue(c.Request.Context(), "ctxVal", v)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
