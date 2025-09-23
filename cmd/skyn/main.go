package main

import (
	"time"
	"skyn/internal/args"
	"skyn/internal/scanner"
	"skyn/internal/ui"
)

func main() {
	Config := args.Parse()

	if Config.Target == "" || Config.Help {
		ui.Help()
		return
	}

	start := time.Now()
	delay := time.Now()
	result := scanner.Scan(Config.Target)
	latency := time.Since(delay)
	duration := time.Since(start)

	ui.Result(result, Config.Target, latency, duration)
}