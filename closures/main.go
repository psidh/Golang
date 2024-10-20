package main

import  "fmt"

func fun() func() int{

  i :=0

  return func() int{
    i++
    return i

  }
}

func main(){

  nextInt := fun()

  fmt.Println(nextInt())
  fmt.Println(nextInt())
  fmt.Println(nextInt())
  fmt.Println(nextInt())
  fmt.Println(nextInt())
}
