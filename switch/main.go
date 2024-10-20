package main

import (
  "fmt"
  "time"
)

func main(){
  i := 2

  fmt.Print("Write ", i, " as ")

  switch i {
    case 1:
      fmt.Println("one was printed")
    case 2:
      fmt.Println("two was printed")

  }
  switch time.Now().Weekday() {
    case time.Saturday, time.Sunday:
      fmt.Println("It's the weekend")
    default:
       fmt.Println("It's the weekday")
  }

  t := time.Now()

  switch{
    case t.Hour() == 12:
     fmt.Println("noon amma")
    default:
     fmt.Println("Not Noon")
  }

  function :=  func(i interface{}){
    switch t:= i.(type){
    case bool:
      fmt.Println("boolean")
    case int:
      fmt.Println("integer")
    default:
      fmt.Println("type unknown", t)
    }
  }

    function(true)
    function(9)
    function("hi")
}
