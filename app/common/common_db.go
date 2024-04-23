package common

import "github.com/injoyai/goutil/database/sqlite"

var (
	DB = sqlite.NewXormWithPath("./data/database/db.db")
)
