package middleware

//
//import (
//	"go-judge/appCommon"
//	"go-judge/component"
//	"go-judge/component/tokenprovider/jwt"
//	"go-judge/modules/user/usermodel"
//	"context"
//	"errors"
//	"fmt"
//	"github.com/gin-gonic/gin"
//	"strings"
//)
//
//type AuthStore interface {
//	FindDataByCondition(ctx context.Context, conditions map[string]interface{}) (*usermodel.User, error)
//}
//
//func ErrWrongAuthHeader(err error) *appCommon.AppError {
//	return appCommon.NewCustomError(
//		err,
//		fmt.Sprintf("wrong authen header"),
//		fmt.Sprintf("ErrWrongAuthHeader"),
//	)
//}
//
//func extractTokenFromHeaderString(s string) (string, error) {
//	parts := strings.Split(s, " ")
//	//"Authorization" : "Bearer {token}"
//
//	if parts[0] != "Bearer" || len(parts) < 2 || strings.TrimSpace(parts[1]) == "" {
//		return "", ErrWrongAuthHeader(nil)
//	}
//
//	return parts[1], nil
//}
//
//func RequiredAuth(appCtx component.AppContext, authStore AuthStore) func(c *gin.Context) {
//	tokenProvider := jwt.NewTokenJWTProvider(appCtx.SecretKey())
//
//	return func(c *gin.Context) {
//		token, err := extractTokenFromHeaderString(c.GetHeader("Authorization"))
//
//		if err != nil {
//			panic(err)
//		}
//
//		payload, err := tokenProvider.Validate(token)
//		if err != nil {
//			panic(err)
//		}
//
//		user, err := authStore.FindDataByCondition(c.Request.Context(), map[string]interface{}{"id": payload.UserId})
//
//		if err != nil {
//			panic(err)
//		}
//
//		if user.Status != usermodel.UserStatusActive {
//			panic(appCommon.ErrNoPermission(errors.New("user has been deleted or banned")))
//		}
//
//		c.Set("user", user)
//		c.Next()
//	}
//}
