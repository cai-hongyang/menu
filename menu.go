package main

import "fmt"

func main() {
	var str string
	fmt.Println("输入hello、help、quit查看更多")
	for {
		fmt.Scanf("%s",&str)
		switch str {
		case "hello":	
			fmt.Println("Hello, 世界")
		case "help":	
			fmt.Println("This is a help about this project")
		case "quit":	
			break
		default:
			fmt.Println("error")
		}
	}
}