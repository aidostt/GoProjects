package main

import "groupie-tracker/internal/app"

const (
	port = 3000
)

func main() {
	s := app.NewServer("localhost", port)
	s.Start()
}
