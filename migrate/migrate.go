package migrate

import (
	"zx/database"
	"zx/model"
)

var db = database.Init()

func Model() {
	db.AutoMigrate(&model.User{})
}
