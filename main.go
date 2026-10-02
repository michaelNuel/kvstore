package main 

import (
    "fmt"
)

func main() {
    t := &Testing{}
    t.Set("name", "Michael")
    fmt.Println(t.Get("name"))

    r := t.Get("name")
    fmt.Println(r)
} 