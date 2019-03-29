package main

import (
    "fmt"
)

type Wg struct {
    Name, age int
}

func (wg *Wg) print (){
    fmt.Println(wg.Name)
}
func main(){
    var wg Wg
    wg.print()
}
