package main

import (
	"log"
	"os"

	_ "github.com/rocketvision/beatriz/checks/a"
	_ "github.com/rocketvision/beatriz/checks/attrs"
	_ "github.com/rocketvision/beatriz/checks/basic"
	_ "github.com/rocketvision/beatriz/checks/img"
	_ "github.com/rocketvision/beatriz/checks/inline"
	_ "github.com/rocketvision/beatriz/checks/stack"
)

func main() {
	log.SetFlags(0)

	args := ParseFlags()
	if len(args) < 1 {
		log.Printf("Uso: %v <arquivo ou diretório>", os.Args[0])
		os.Exit(0)
	}

	manager := NewManager()
	defer manager.DeleteTemporaries()

	for _, arg := range args {
		if err := manager.Process(arg); err != nil {
			log.Fatal(err)
		}
	}
	manager.PrintSummary()
}
