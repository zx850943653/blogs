package untils

//错误处理
import (
	"github.com/gin-gonic/gin"
	"net/http"
	"zx/untils_model"
)

const (
	SERVER_ERROR    = 1000 // 系统错误
	NOT_FOUND       = 1001 // 401错误
	UNKNOWN_ERROR   = 1002 // 未知错误
	PARAMETER_ERROR = 1003 // 参数错误
	AUTH_ERROR      = 1004 // 错误
)

// 500 错误处理
func ServerError() *untils_model.APIException {
	return untils_model.NewAPIException(http.StatusInternalServerError, SERVER_ERROR, http.StatusText(http.StatusInternalServerError))
}

// 404 错误
func NotFound() *untils_model.APIException {
	return untils_model.NewAPIException(http.StatusNotFound, NOT_FOUND, http.StatusText(http.StatusNotFound))
}

// 未知错误
func UnknownError(message string) *untils_model.APIException {
	return untils_model.NewAPIException(http.StatusForbidden, UNKNOWN_ERROR, message)
}

// 参数错误
func ParameterError(message string) *untils_model.APIException {
	return untils_model.NewAPIException(http.StatusBadRequest, PARAMETER_ERROR, message)
}

func HandleNotFound(c *gin.Context) {
	handleErr := NotFound()
	handleErr.Request = c.Request.Method + " " + c.Request.URL.String()
	c.JSON(handleErr.Code, handleErr)
	return
}

func Wrapper(handler untils_model.HandlerFunc) func(c *gin.Context) {
	return func(c *gin.Context) {
		var (
			err error
		)
		err = handler(c)
		if err != nil {
			var apiException *untils_model.APIException
			if h, ok := err.(*untils_model.APIException); ok {
				apiException = h
			} else if e, ok := err.(error); ok {
				if gin.Mode() == "debug" {
					// 错误
					apiException = UnknownError(e.Error())
				} else {
					// 未知错误
					apiException = UnknownError(e.Error())
				}
			} else {
				apiException = ServerError()
			}
			apiException.Request = c.Request.Method + " " + c.Request.URL.String()
			c.JSON(apiException.Code, apiException)
			return
		}
	}
}
