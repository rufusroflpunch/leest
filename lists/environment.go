package lists

import "github.com/jinzhu/gorm"

type HandlerEnvironment struct {
	Db *gorm.DB
}
