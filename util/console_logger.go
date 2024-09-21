package util

import "fmt"

func NewConsoleLogger() Logger {
	return &ConsoleLogger{
		colors: map[string]string{
			"red":    "\033[91m",
			"green":  "\033[92m",
			"orange": "\033[33m",
			"yellow": "\033[93m",
			"gray":   "\033[90m",
			"cyan":   "\033[36m",
			"reset":  "\033[0m",
		},
	}
}

type ConsoleLogger struct {
	colors map[string]string
}

func (l ConsoleLogger) Trace(message string) {
	fmt.Println(message)
}

func (l ConsoleLogger) Debug(message string) {
	fmt.Println(l.colors["gray"], message, l.colors["reset"])
}

func (l ConsoleLogger) Info(message string) {
	fmt.Println(l.colors["cyan"], message, l.colors["reset"])
}

func (l ConsoleLogger) Warn(message string) {
	fmt.Println(l.colors["orange"], message, l.colors["reset"])
}

func (l ConsoleLogger) Error(message string) {
	fmt.Println(l.colors["red"], message, l.colors["reset"])
}

func (l ConsoleLogger) Fatal(message string) {
	fmt.Println(l.colors["red"], message, l.colors["reset"])
}

func (l ConsoleLogger) Success(message string) {
	fmt.Println(l.colors["green"], message, l.colors["reset"])
}
