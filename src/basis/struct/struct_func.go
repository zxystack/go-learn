package main

import "fmt"

type test struct {
    fn []func() int
}
func f1() int {
    return 1
}
func f2() int {
    return 2
}

func main(){
    //fff = []func(){f1, f2}
    var t test = test{[]func() int {f1, f2}}
    fmt.Println(t.fn)
}
