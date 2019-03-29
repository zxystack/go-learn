package main

import (
    "fmt"
    "sort"
    "math/rand"
)

type Hero struct {
    Name string
    Age int
}

type Heros []Hero

func (h Heros) Len() int {
    return len(h)
}

func (h Heros) Less(i, j int) bool{
    return h[i].Age < h[j].Age
}

func (h Heros) Swap(i, j int){
    temp := h[i]
    h[i] = h[j]
    h[j] = temp
}

func main(){
    var hs Heros
    for i:=0; i<10; i++ {
        h := Hero{
	    Name: fmt.Sprintf("name%d", i),
	    Age: rand.Intn(100),
	}
	hs = append(hs, h)
    }
    fmt.Println("before sort", hs)
    sort.Sort(hs)
    fmt.Println("after sort", hs)
}
