package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/lunny/xorm"
)

func main() {
	xorm.NewEngine("mysql", dataSourceName)
}
