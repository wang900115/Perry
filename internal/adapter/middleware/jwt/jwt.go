package jwt

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
	responser "github.com/wang900115/Perry/internal/adapter/response"
	redisinterface "github.com/wang900115/Perry/internal/domain/interface/redis"
)

type JWT struct {
	response responser.Response
	token    redisinterface.Token
}

func NewJWT(response responser.Response, token *redisinterface.Token) *JWT {
	return &JWT{response: response, token: *token}
}

func (j *JWT) Middleware(c *gin.Context) {
	token, err := j.extractToken(c)
	if err != nil {
		j.response.ClientFail400(c, err)
		c.Abort()
		return
	}

	tokenClaims, err := j.token.Validate(c, token)
	if err != nil {
		j.response.ClientFail401(c, err)
		c.Abort()
		return
	}
	c.Set("user_id", tokenClaims.Subject)
	c.Set("session_id", tokenClaims.Id)
	c.Next()
}

func (j *JWT) extractToken(c *gin.Context) (string, error) {
	authorization := c.GetHeader("Authorization")

	if authorization == "" {
		return "", errors.New("no token")
	}

	if !strings.HasPrefix(authorization, "Bearer ") {
		return "", errors.New("invalid authorization format")
	}

	token := strings.TrimPrefix(authorization, "Bearer ")
	return token, nil
}
