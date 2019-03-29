package main

import (
    "fmt"
    "reflect"
)

func main(){
    var s string = "test"
    rs := reflect.ValueOf(s)
    fmt.Println(rs)
    rs.SetString("test2")
    fmt.Println(rs)
}
