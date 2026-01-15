package main

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

func LogToFile(message string) {
	// os.O_APPEND: Adiciona ao final
	// os.O_CREATE: Cria se não existir
	// os.O_WRONLY: Abre apenas para escrita
	f, err := os.OpenFile("orbisflow_access.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	defer f.Close()

	// O formato "2006-01-02 15:04:05" é o padrão do Go para data/hora
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	f.WriteString(fmt.Sprintf("[%s] %s\n", timestamp, message))
}

// Guarda os dados de quem acessou o vídeo
type ClientSession struct {
	IP        string
	VideoName string
	Timestamp string
}

// Vetor(Slice) de acessos (Global)
var ActiveClients []ClientSession

// ShowClients imprime a lista de quem acessou o servidor
func ShowClients() {
	fmt.Println("\n--- Access History ---")
	if len(ActiveClients) == 0 {
		fmt.Println("No clients have requested videos yet.")
	} else {
		for _, client := range ActiveClients {
			fmt.Printf("[%s] IP: %s requested: %s\n", client.Timestamp, client.IP, client.VideoName)
		}
	}
	fmt.Println("-----------------------")
}

// Mostra os dados que o programa está manipulando no momento
func ShowRAMUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// Converte bytes para Megabytes para ficar legível
	fmt.Printf("RAM Alloc = %v MiB\n", m.Alloc/1024/1024)
	fmt.Printf("Total Alloc = %v MiB\n", m.TotalAlloc/1024/1024)
	fmt.Printf("Sys = %v MiB\n", m.Sys/1024/1024)
}
