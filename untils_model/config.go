package untils_model

import "github.com/gin-gonic/gin"

//配置文件实体类
type Input struct {
	//view:用于注释
	//json:json形式
	//from:解释该字段来自哪里，比如那个表
	//binding: required:必须字段 email:email形式
	//grom:数据库中列名
	Id int `view:"id号" json:"id" from:"id" binding:"required" gorm:"column:id"`
}

type Output struct {
	Database DatabaseConfig `json:"database"`
	Self     SelfConfig     `json:"self"`
}

type DatabaseConfig struct {
	Types  string `json:"types"`
	Local  string `json:"local"`
	Online string `json:"online"`
}

type SelfConfig struct {
	Port string `json:"port"`
	Flag int    `json:"flag"`
	Tag  int    `json:"tag"`
}

// api错误的结构体
type APIException struct {
	Code      int    `json:"-"`
	ErrorCode int    `json:"error_code"`
	Msg       string `json:"msg"`
	Request   string `json:"request"`
}

// 实现接口
func (e *APIException) Error() string {
	return e.Msg
}

func NewAPIException(code int, errorCode int, msg string) *APIException {
	return &APIException{
		Code:      code,
		ErrorCode: errorCode,
		Msg:       msg,
	}
}

type HandlerFunc func(c *gin.Context) error
