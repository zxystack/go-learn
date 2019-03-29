package main

import (
    "database/sql"
    "log"
    _ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB
type User struct {
    Name string
    Age int
}

func init(){
    var err error
    Db, err = sql.Open("mysql", "root:123456@tcp(localhost:3306)/test")
    if err != nil{
        log.Fatal(err)
    }
}

func main(){
    statement := "insert into user (name, age) values (?, ?)"
    stmt, err:= Db.Prepare(statement)
    if err != nil{
        log.Fatal(err)
    }
    user := User{Name: "tom", Age: 28}
    stmt.Exec(user.Name, user.Age)
}
