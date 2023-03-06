package main

import "fmt"

func Run() error {
	fmt.Println("starting up application")
	return nil
}

func main() {
	fmt.Println("rest api")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
