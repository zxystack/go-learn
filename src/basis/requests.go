package main

import "net/http"
import "fmt"
import _ "basis/database"

func main(){
    url := "http://www.baidu.com"
    resp, err := http.Get(url)
    if err != nil {
       fmt.Println("err is ", err) 
       return
    }
    for header, content := range(resp.Header){
        fmt.Println(header, "=", content)
    }
}
