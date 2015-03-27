package main

import "fmt"

func main(){

m := map[string]string{
	"hi": "hi1",
	"yo": "yo1",
}

if name, ok := m["hi"]; ok {
	fmt.Println(name, ok)
}



}
