package ReturnData

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

// 定义返回数据结构体
type ReturnData struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

// 构造函数,配置默认值
func NewReturnData() ReturnData {
	returnData := ReturnData{}
	returnData.Status = 200
	data := make(map[string]interface{})
	returnData.Data = data
	return returnData
}

type Info struct {
	ReturnData
}

func GetRequestInfo(r *gin.Context, info interface{}) (bindInfo interface{}, err error) {
	// 首先获取请求的类型
	requestMethod := r.Request.Method
	if requestMethod == "GET" {
		err = r.ShouldBindQuery(&info)
	} else if requestMethod == "POST" {
		err = r.ShouldBindJSON(&info)
	} else {
		err = errors.New("不支持请求类型")
	}
	return info, err
}
