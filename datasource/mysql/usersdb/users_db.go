package usersdb

import (
	"database/sql"
	"fmt"
	"log"

	// imported but not used in this package
	_ "github.com/go-sql-driver/mysql"
	"github.com/uwaifo/bookstore_users_api/utils"
)

const ()

//Client . . .
var (
	Client   *sql.DB
	username = "db_username"
	password = "db_userpassword"
	host     = "db_host"
	schema   = "db_schema"
)

func init() {
	//golangappdb.ce5cqvqcduyi.us-east-1.rds.amazonaws.com
	datasourceName :=
		fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
			utils.GoDotEnvVariable(username),
			utils.GoDotEnvVariable(password),
			utils.GoDotEnvVariable(host),
			utils.GoDotEnvVariable(schema),

			//"uwaifo",
			//"247Secured",
			//"golangappdb.ce5cqvqcduyi.us-east-1.rds.amazonaws.com",
			//"gbs_users_db",
		)
	var err error
	Client, err = sql.Open("mysql", datasourceName)
	if err != nil {
		panic((err))
	}

	if err = Client.Ping(); err != nil {
		panic(err)
	}
	//utils.GetEnv()
	log.Println("database successfully configured")
	//dbvar := utils.GetEnv()
	//fmt.Println(utils.GetEnv())

}
