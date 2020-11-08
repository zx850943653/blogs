package untils_model

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
