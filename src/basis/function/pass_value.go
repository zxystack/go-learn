package main

import "fmt"

type Student struct {
    name, age int
}

func change(s *Student) {
    s.name = 1
}

func main(){
    var s Student
    fmt.Println(s)
    change(&s)
    fmt.Println(s)
}
