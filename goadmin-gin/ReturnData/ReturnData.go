package ReturnData

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
)

const (
	ERROR   = 401
	SUCCESS = 200
)

// 定义返回数据结构体
type ReturnData struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

// 构造函数,配置默认值
func newReturnData(status int, message string, data map[string]interface{}) ReturnData {
	returnData := ReturnData{}
	returnData.Status = 200
	returnData.Message = message
	returnData.Data = data
	return returnData
}

// 获取请求入口的数据
func GetRequestInfo(r *gin.Context, info interface{}) (err error) {
	// 首先获取请求的类型
	requestMethod := r.Request.Method
	if requestMethod == "GET" {
		err = r.ShouldBindQuery(&info)
	} else if requestMethod == "POST" {
		err = r.ShouldBindJSON(&info)
	} else {
		err = errors.New("不支持请求类型")
	}
	return err
}

func result(status int, message string, data map[string]interface{}, c *gin.Context) {
	returnData := newReturnData(status, message, data)
	c.JSON(http.StatusOK, returnData)
}

// 操作成功,并返回数据和成功信息
func OK(message string, data map[string]interface{}, c *gin.Context) {
	result(SUCCESS, message, data, c)
}

// 操作成功,只返回成功信息
func OKWithMessage(message string, c *gin.Context) {
	result(SUCCESS, message, nil, c)
}

// 操作成功,只返回成功数据
func OKWithData(data map[string]interface{}, c *gin.Context) {
	result(SUCCESS, "操作成功", data, c)
}

//func OKWithList[T any](List []T, count any, c *gin.Context) {
//	if len(List) == 0 {
//		List = []T{}
//	}
//	Result(SUCCESS, ListResponse[T]{
//		List:  List,
//		Count: count,
//	}, "成功", c)
//}

// 操作失败,并返回数据和失败信息
func Fail(data map[string]interface{}, message string, c *gin.Context) {
	result(ERROR, message, data, c)
}

// 操作失败,并返回失败信息
func FailWithMessage(message string, c *gin.Context) {
	result(ERROR, message, nil, c)
}

//func FailWithCode(code ErrorCode, c *gin.Context) {
//	msg, ok := ErrorMap[code]
//	if ok {
//		Result(int(code), map[string]interface{}{}, msg, c)
//		return
//	}
//	Result(ERROR, map[string]interface{}{}, "未知错误", c)
//}
