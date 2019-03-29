package main

import (
    "fmt"
    "reflect"
)

func reflectTest(d interface{}){
    //rType 的真实类型为reflect.Type
    rType := reflect.TypeOf(d)
    fmt.Println(rType)
    //rVal 的真实类型为reflect.Value
    rVal := reflect.ValueOf(d)
    fmt.Printf("rval=%T\n", rVal)
    //realValue 为一个接口
    realValue := rVal.Interface()
    fmt.Printf("realValue %T\n",realValue)
    // 这里将报错  因为realValue 并不是一个int类型
    //n := realValue + 2
    // 通过断言 将realValue转换为一个Int类型
    value := realValue.(int)
    n := value + 2
    fmt.Println(value)
    fmt.Println(n)
}

func main(){
    var a int = 100
    reflectTest(a)
}
