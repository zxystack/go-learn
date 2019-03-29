package database

import (
    "fmt"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

func init(){
    Db, err := sql.Open("mysql", "dbname=test")
    if err != nil{
        fmt.Println(err)
    }
    fmt.Println("connected", Db)
    ping_err := Db.Ping()
    if err != nil{
        fmt.Println(ping_err)
    }else{
        rows, selectErr := Db.Query("select * from user")
    }
}

func main(){
    fmt.Println("connect to mysql")
}
