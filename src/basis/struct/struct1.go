package main 

import "fmt"

func main(){
    type singleton struct{name int}
    var s *singleton    //s 为singleton类型的指针  即s指向的地址的值的类型为singleton
    ss := singleton{}
    s = &ss
    fmt.Println((*s).name)
}
