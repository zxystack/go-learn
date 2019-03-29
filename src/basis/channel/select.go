package main

import "fmt"
import "time"

func add(ch chan int){
   ch <- 1
   fmt.Println("in add", ch)
}

func main(){
    ch := make(chan int)
    go func(){
        add(ch)
    }()
loop:
    for {
        select {
            case <- ch:
	        fmt.Println("in case 1")
//		return
                break loop
	    default:
	        fmt.Println("in default")
		time.Sleep(1 * time.Second)
        }
    }
    close(ch)
    fmt.Println(<-ch)
}
