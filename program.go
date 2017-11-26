package main

// Package is called aw
import "github.com/deanishe/awgo"

func run() {
	// Your workflow starts here
	aw.NewItem("First result!")
	aw.SendFeedback()
}

func main() {
	aw.Run(run)
}
