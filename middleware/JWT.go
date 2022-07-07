/**
  @author:BOEN
  @data:2022/6/24
  @note:
**/
package middleware

import (
	"github.com/gin-gonic/gin"
	"memorandumProject/pkg/e"
	"memorandumProject/pkg/utils"
	"time"
)

/**
 * @Author Cofeesy
 * @Description //JWT鉴权
 * @Date 12:40 2022/6/24
 * @Param
 * @return gin.HandlerFunc "HandlerFunc defines the handler used by gin middleware as return value."
 **/
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		code := 200
		var data interface{}
		token := c.GetHeader("Authorization")
		if token == "" {
			code = 404
		} else {
			claim, err := utils.ParseToken(token)
			if err != nil {
				code = e.ErrorAuthCheckTokenFail //无权限
			} else if time.Now().Unix() > claim.ExpiresAt {
				code = e.ErrorAuthCheckTokenTimeout //过期无效token
			}
		}
		if code != e.SUCCESS {
			c.JSON(400, gin.H{
				"Status": code,
				"msg":    e.GetMsg(code),
				"data":   data,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
