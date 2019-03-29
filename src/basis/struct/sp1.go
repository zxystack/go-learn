package main

import "basis/struct/sp1"
import "fmt"

func main(){
   // _ := sp1.Person{"1", "2"}  compile  error  the variable that unuse need define with var 
   var v sp1.Person = sp1.Person{"1", "2"}
   fmt.Println(v)
}
