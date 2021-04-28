package main

import "fmt"

const spanish = "Spanish"
const frech = "Frech"
const englistHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frechHelloPrefix = "Bonjour, "


func Hello(name string, language string) string {
	if name == ""{
		name = "world"
	}

	// if language == spanish {
	// 	return spanishHelloPrefix + name
	// }
	// if language == frech {
	// 	return frechHelloPrefix + name 
	// }
	// prefix := englistHelloPrefix + name

	// switch language {
	// case spanish:
	// 	prefix = spanishHelloPrefix + name
	// case frech:
	// 	prefix = frechHelloPrefix + name
	// }
	return greetingPrefix(language)+name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case frech:
		prefix = frechHelloPrefix
	default:
		prefix = englistHelloPrefix
		
	}
	return 
}

func main(){
	fmt.Println(Hello("Jerry", "English"))
}