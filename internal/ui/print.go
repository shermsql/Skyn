package ui

import (
	"fmt"
	"time"
	"skyn/internal/scanner"
)

func Result(result []scanner.Status, target string, latency time.Duration, duration time.Duration) {
	open := []scanner.Status{}
	closed := 0

	for _, r := range result {
		if r.Open {
			open = append(open, r)
		} else {
			closed++
		}
	}

	fmt.Printf("Starting Skyn ( https://github.com/shermsql/skyn ) At %s\n", time.Now().Format("2006-01-02 15:04"))
	fmt.Printf("Skyn Scan Report For %s\n", target)
	fmt.Printf("Host Is Up (%.3fs Latency).\n", latency.Seconds())

	if closed > 0 {
		fmt.Printf("Not Shown, %d Closed TCP Ports (Connection Refused)\n", closed)
	}

	fmt.Println("PORT     STATE SERVICE")
	for _, p := range open {
		fmt.Printf("%-8d Open  %s\n", p.Port, p.Service)
	}

	fmt.Printf("Skyn, Done! 1 IP Address (1 Host Up) Scanned In %.2f Seconds\n", duration.Seconds())
}