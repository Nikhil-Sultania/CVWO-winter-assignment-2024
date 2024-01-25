package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/CVWO/sample-go-app/internal/models"
	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
}

var Db *sql.DB

func Init() {
	var err error
	Db, err = sql.Open("mysql",
		"newuser:1@tcp(127.0.0.1:3306)/datas")
	if err != nil {
		log.Fatal(err)
	}

}

func CreatePost(post models.Thread) {
	fmt.Print("HI")
	result, err := Db.Exec("Insert INTO threads VALUES(?,?,?)", post.Owner, post.Title, post.Description)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(result)
}
func GetDB() (*Database, error) {
	return &Database{}, nil
}
