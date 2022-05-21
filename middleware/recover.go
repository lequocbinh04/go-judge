package middleware

import (
	"github.com/gin-gonic/gin"
	"go-judge/appCommon"
	"go-judge/component"
)

func Recover(ac component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.Header("Content-Type", "application/json")

				if appErr, ok := err.(*appCommon.AppError); ok {
					c.AbortWithStatusJSON(appErr.StatusCode, appErr)
					panic(err)
					return
				}

				appErr := appCommon.ErrInternal(err.(error))
				c.AbortWithStatusJSON(appErr.StatusCode, appErr)
				panic(err)
				return
			}
		}()

		c.Next()
	}
}
