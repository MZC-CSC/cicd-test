package main

import "fmt"

var version = "dev"

func main() {

 fmt.Printf("Version: %s\n", version)

 fmt.Println(hello())
 fmt.Println("hi")
}

func hello() string {
 return "Hello Glang"
}
