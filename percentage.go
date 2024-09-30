package main

import "fmt"

func Percentage(partial float64, total float64) string {
	return fmt.Sprintf("%.2f", (100*partial)/total) + "%"
}
