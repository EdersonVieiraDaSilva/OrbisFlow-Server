package main

import (
	"fmt"
	"net/http"
	"time"
)

// StartServer configura o roteamento e sobe o servidor na porta 8080
func StartServer() {
	port := ":8080"

	// Configurando o roteamento principal com Middleware de segurança
	http.HandleFunc("/", BasicAuthMiddleware(func(w http.ResponseWriter, r *http.Request) {
		videoName := r.URL.Path
		if videoName == "/" {
			fmt.Fprintf(w, "Welcome to OrbisFlow! Protected Acess")
			return
		}

		// Validação e limpeza do caminho do arquivo
		safeVideoName := ValidatePath(videoName)

		//Sessão para o monitoramento interno (show clients)
		newSession := ClientSession{
			IP:        r.RemoteAddr, // Pega o IP e Porta do cliente
			VideoName: safeVideoName,
			Timestamp: time.Now().Format("15:04:05"), // Hora atual formatada
		}
		ActiveClients = append(ActiveClients, newSession)

		// Chamada do Log. Registros salvos em arquivo de texto externo
		logMsg := fmt.Sprintf("CLIENT_AUTH_ACESS: %s | FILE: %s", r.RemoteAddr, safeVideoName)
		LogToFile(logMsg)

		// Módulo para cuidar do arquivo e do streaming é chamado
		StreamVideo(w, r, safeVideoName)

	}))

	fmt.Println("Server started at http://localhost" + port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println("Fatal error in server:", err)
	}
}
