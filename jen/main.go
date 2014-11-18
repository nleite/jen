package main

import (
    "fmt"
    g "jen/generators"
)

func main(){
    fmt.Println("Generate JSON")
    var gen = g.NewGenerator()
    fmt.Println( "something: ", gen.GenerateString("hey", 100 ))
}



