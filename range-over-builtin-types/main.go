package main


import "fmt"

func main(){

  nums := []int{2, 3, 4}

  nu := []string{"1", "2"};
  fmt.Println("hi there");

  for idx, value := range nums{
    if value == 3{
      fmt.Println(3, idx)
    }
  }

  kvs := map[string]string{"a":"apple", "b":"ban"}

  for k, v := range kvs{
    fmt.Printf("%s -> %s\n", k ,v)
  }

}
