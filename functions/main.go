package main

import  "fmt"

func plus(a int, b int) int {
  return a + b
}

func multiply(a int, c  int) int{
  return a * c
}

func main(){
  res := plus(1, 2)
  fmt.Println("1+2: ", res)
  fmt.Println("2*3: ", multiply(2, 3))
}
