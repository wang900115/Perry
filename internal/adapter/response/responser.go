package responser

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// 成功 (200, 201, 204)
// client 失敗 (400, 401, 403, 404, 409, 422)
// server 失敗 (500, 502, 503, 504)

// 服務成功
func (r Response) Success204(ctx *gin.Context) {
	ctx.JSON(http.StatusNoContent, Response{
		Code:    200,
		Message: "success",
	})
}

// 服務成功(帶資源)
func (r Response) Success200(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "success",
		Data:    data,
	})
}

// 服務成功(建立資源)
func (r Response) Success201(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusCreated, Response{
		Code:    201,
		Message: "success",
		Data:    data,
	})
}

// 服務失敗
func (r Response) ServerFail500(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusInternalServerError, Response{
		Code:    500,
		Message: "fail",
		Data:    err,
	})
}

// 服務失敗(上游錯誤)
func (r Response) ServerFail502(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusBadGateway, Response{
		Code:    502,
		Message: "fail",
		Data:    err,
	})
}

// 服務失敗(無法附載)
func (r Response) SereverFail503(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusServiceUnavailable, Response{
		Code:    503,
		Message: "fail",
		Data:    err,
	})
}

// 服務失敗(連線逾時)
func (r Response) ServerFail504(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusGatewayTimeout, Response{
		Code:    504,
		Message: "fail",
		Data:    err,
	})
}

// 客戶格式錯誤
func (r Response) ClientFail400(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusBadRequest, Response{
		Code:    400,
		Message: "fail",
		Data:    err,
	})
}

// 客戶權限不足
func (r Response) ClientFail401(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusUnauthorized, Response{
		Code:    401,
		Message: "fail",
		Data:    err,
	})
}

// 客戶權限禁止訪問
func (r Response) ClientFail403(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusForbidden, Response{
		Code:    403,
		Message: "fail",
		Data:    err,
	})
}

// 客戶找不到資源
func (r Response) ClientFail404(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusNotFound, Response{
		Code:    404,
		Message: "fail",
		Data:    err,
	})
}

// 客戶資源衝突
func (r Response) ClientFail409(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusConflict, Response{
		Code:    409,
		Message: "fail",
		Data:    err,
	})
}

// 客戶格式正確但沒有該資料
func (r Response) ClientFail422(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusUnprocessableEntity, Response{
		Code:    422,
		Message: "fail",
		Data:    err,
	})
}

// RateLimiter
func (r Response) ClientFail429(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusTooManyRequests, Response{
		Code:    429,
		Message: "fail",
		Data:    err,
	})
}
