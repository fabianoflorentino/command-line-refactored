package main

import (
	"fmt"

	"github.com/mateusprado/command-line-refactored/handlers"
)

func main() {
	fmt.Println("running rundeck job")
	fmt.Printf("%s with id %s\n", handlers.SliBuilder().Name, handlers.SliBuilder().Id)
	fmt.Println(handlers.ListAllFromHttpService())
	fmt.Println(handlers.RetrieveShippings())

}
