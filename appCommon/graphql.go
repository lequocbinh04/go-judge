package appCommon

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-judge/component"
)

func GinContextFromContext(ctx context.Context) (*gin.Context, component.AppContext, error) {
	v_ := ctx.Value("ctxVal")
	v, ok := v_.(map[string]interface{})
	ginContext := v["GinContextKey"]
	appContext := v["AppContext"]
	if ginContext == nil {
		err := fmt.Errorf("could not retrieve gin.Context")
		return nil, nil, err
	}
	appCtx, ok := appContext.(component.AppContext)
	gc, ok := ginContext.(*gin.Context)
	if !ok {
		err := fmt.Errorf("gin.Context has wrong type")
		return nil, nil, err
	}
	return gc, appCtx, nil
}
