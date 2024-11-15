package main

import "fmt"

var (
	commitHash string
	version    string
)

func main() {
	fmt.Printf("Version %s, on commit %s\n", version, commitHash)
}
