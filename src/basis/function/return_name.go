package main

import (
    "fmt"
)
// 如果返回值 被命名 则参数不用在函数体当中定义，函数必须以return 结束
func sub (x, y int) (z int){
    z = x + y
    return
}

func main(){
    result := sub(3, 4)
    fmt.Println("result is", result)
}
