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

	// TODO: This should probably be moved to some env handler util
	username := os.Getenv("MYSQL_USERNAME")
	password := os.Getenv("MYSQL_PASSWORD")
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	database := os.Getenv("MYSQL_DB")

	// XXX does golang support template strings?
	connectionString := username + ":" + password + "@(" + host + ":" + port + ")/" + database

	orm.RegisterDataBase("default", "mysql", connectionString)
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
