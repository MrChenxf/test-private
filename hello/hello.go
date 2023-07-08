package main

import(

   "fmt"

    "example.com/greetings"
)

func main(){
    message :=greetings.Hello("gladys")
    fmt.Println(message)
}