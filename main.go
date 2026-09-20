package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	tasks := os.Args[1:]

	fmt.Printf("Arguments with length %d\n", len(tasks))
	fmt.Println(tasks)

	var firstTask *CommandDirective = initFromArray(tasks)

	if firstTask == nil {
		log.Fatal("no tasks provided")
	}

	var taskCount int = getLength(firstTask)
	fmt.Printf("Number of tasks: %d\n", taskCount)

	fmt.Println("Let's run the first command")
	commandExec(firstTask.arg)

}
