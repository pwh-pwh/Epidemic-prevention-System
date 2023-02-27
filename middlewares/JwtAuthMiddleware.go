package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/pwh-pwh/Epidemic-prevention-System/common"
	"github.com/pwh-pwh/Epidemic-prevention-System/response"
	"github.com/pwh-pwh/Epidemic-prevention-System/service/user_service"
	"net/http"
)

func JwtAuth(sau string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenStr := ctx.GetHeader(common.GetJwtHeader())
		if tokenStr == "" {
			response.Response(ctx, http.StatusUnauthorized, "请先登录", nil)
			ctx.Abort()
			return
		}
		_, claims, err := common.ParseToken(tokenStr)
		if err != nil {
			response.Response(ctx, http.StatusForbidden, "token 不合法", nil)
			ctx.Abort()
			return
		}
		isExpire := common.IsExpire(claims)
		if isExpire {
			response.Response(ctx, http.StatusForbidden, "token过期", nil)
			ctx.Abort()
			return
		}
		//如果所需权限不为空
		if sau != "" {
			authorityList := user_service.GetUserAuthorityList(claims.UserName)
			ctx.Set("auList", authorityList)
			flag := false
			for _, au := range authorityList {
				if au == sau {
					flag = true
					break
				}
			}
			//权限不通过
			if !flag {
				response.Response(ctx, http.StatusForbidden, "权限不足", nil)
				ctx.Abort()
				return
			}
		}
		ctx.Set("username", claims.UserName)
		ctx.Next()
	}
}
