package main

import (
    "fmt"
    "sync"
    "runtime"
)

var (
    counter int
    wg sync.WaitGroup
)

func increase(id int) {
    defer wg.Done()
    for i:=0; i<2;i++ {
	runtime.Gosched()
        counter ++
	fmt.Println(counter)
    }
}

func main(){
    wg.Add(2)
   go increase(1)
   go increase(2)
   wg.Wait()
   fmt.Println(counter)
}
