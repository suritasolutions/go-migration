package util

import "fmt"

var colors = map[string]string{
	"red":    "\033[31m",
	"green":  "\033[32m",
	"yellow": "\033[33m",
	"gray":   "\033[90m",
	"reset":  "\033[0m",
}

func Print(color string, message string) {
	fmt.Println(colors[color], message, colors["reset"])
}
