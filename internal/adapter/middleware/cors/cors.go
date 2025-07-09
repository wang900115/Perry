package cors

import (
	"github.com/gin-gonic/gin"
	responser "github.com/wang900115/Perry/internal/adapter/response"
)

type CORS struct {
	response     responser.Response
	allowOrigins []string
}

func NewCORS(response responser.Response, allowOrigins []string) *CORS {
	return &CORS{response: response, allowOrigins: allowOrigins}
}

func (cors *CORS) Middleware(c *gin.Context) {
	origin := c.Request.Header.Get("Origin")

	for _, o := range cors.allowOrigins {
		if origin == o {
			// 根據 Origin 決定回傳內容
			c.Header("Vary", "Origin")
			c.Header("Access-Control-Allow-Origin", origin)
			// 支援cookie/jwt
			c.Header("Access-Control-Allow-Credentials", "true")
			c.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, PATCH, DELETE")
			// 哪些自訂 header 是允許被前端送出來的
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization, token")
			// 哪些 response header 是允許在前端程式碼中讀取的
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			break
		}
	}

	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(204)
		return
	}

	c.Next()
}
