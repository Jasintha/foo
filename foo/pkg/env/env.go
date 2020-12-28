package env

import (
	model "foo/pkg/model"
	gorm "github.com/jinzhu/gorm"
	"os"
)

var RDB *gorm.DB
var CMDCtrl *model.StorageCtrl
var QueryCtrl *model.StorageCtrl
var ESCtrl *model.MessageCtrl

const E2_DIALET = "E2_DIALET"

var E2DIALET = "sqlite3"

const E2_DB = "E2_DB"

var E2Db = "db"

const E2_HOST = "E2_HOST"

var E2Host = ""

const E2_PORT = "E2_PORT"

var E2Port = ""

const E2_USER = "E2_USER"

var E2User = ""

const E2_PWD = "E2_PWD"

var E2Pwd = ""

const E2_URL = "E2_URL"

var E2Url = ""

func LoadEnvs() {

	if val := os.Getenv(E2_DIALET); val != "" {
		E2DIALET = val
	}
	if val := os.Getenv(E2_DB); val != "" {
		E2Db = val
	}
	if val := os.Getenv(E2_HOST); val != "" {
		E2Host = val
	}
	if val := os.Getenv(E2_PORT); val != "" {
		E2Port = val
	}
	if val := os.Getenv(E2_USER); val != "" {
		E2User = val
	}
	if val := os.Getenv(E2_PWD); val != "" {
		E2Pwd = val
	}
	if val := os.Getenv(E2_URL); val != "" {
		E2Url = val
	}
}
