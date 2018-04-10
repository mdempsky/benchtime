// The benchtime command runs a program and summarizes its resource
// usage in Go benchmark format.
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
	"unicode"
)

var name = flag.String("name", "Exec", "benchmark name")

func main() {
	flag.Parse()

	name := *name
	if len(name) == 0 {
		log.Fatal("benchmark name cannot be empty")
	}
	for i, c := range name {
		if i == 0 && !unicode.IsUpper(c) {
			log.Fatal("benchmark name must begin with an upper case character")
		}
		if unicode.IsSpace(c) {
			log.Fatal("benchmark name must not contain space characters")
		}
	}

	args := flag.Args()
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	start := time.Now()
	err := cmd.Run()
	real := time.Since(start)

	user := cmd.ProcessState.UserTime()
	sys := cmd.ProcessState.SystemTime()

	fmt.Printf("Benchmark%s 1 %d real-ns/op %d user-ns/op %d sys-ns/op\n", name, real, user, sys)
	if err != nil {
		log.Fatal(err)
	}
}
