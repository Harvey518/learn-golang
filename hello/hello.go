package main

import "fmt"

const englistHelloPrefix = "Hello, "

func Hello(name string) string {
	return englistHelloPrefix+name
}

func main(){
	fmt.Println(Hello("Jerry"))
}