package main

import (
    "fmt"
    "time"
    "sync"
)

func main(){
    var wg sync.WaitGroup
    for i:=0; i < 2; i++{
        wg.Add(1)
	go func(){
	    defer wg.Done()
	    time.Sleep(5 * time.Second)
	    ii := i
	    fmt.Println(ii)
	}()
    }
    wg.Wait()
}
