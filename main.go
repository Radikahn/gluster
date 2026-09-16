package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {

	tasks := os.Args[1:]

	fmt.Printf("Arguments with length %d\n", len(tasks))
	fmt.Println(tasks)

	var head CommandDirective = startList(0, 1, tasks[0])

	for i := 1; i < len(tasks); i++ {
		if i == 1 {
			var _ = createNode(i, &head, 1, tasks[i])
		}

	}

	cmd := exec.Command("echo", "hello world")

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", out.String())
}
