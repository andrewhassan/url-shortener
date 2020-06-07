package models

import (
	"fmt"
	"os"

	"github.com/astaxie/beego/orm"

	// Import the MySQL driver
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterModel(new(ShortenedURL))
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:root@(mysql:3306)/url_shortener_dev")
}

// SyncDatabase syncs the code-defined schema with the real database
// TODO: We should move this setup to explicit tasks
func SyncDatabase() {
	err := orm.RunSyncdb("default", false, true)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
