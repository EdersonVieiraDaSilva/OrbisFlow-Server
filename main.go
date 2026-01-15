package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("OrbisFlow System: Loading modules...")

	// Função no server.go é chamada
	go StartServer()

	// Loop de comandos
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("OrbisFlow CLI Terminal - Type 'help' for commands")

	for {
		fmt.Print("OrbisFlow> ")
		if scanner.Scan() {
			input := scanner.Text()

			switch input {
			case "show data":
				ShowRAMUsage() // Esta função esta no monitor.go
			case "finish":
				fmt.Println("Shutting down...")
				os.Exit(0)
			case "help":
				fmt.Println("Available: show data, show clients, finish, exit")
			default:
				fmt.Println("Unknown command. Type 'help'.")
			case "show clients":
				ShowClients()
			}
		}
	}
}
