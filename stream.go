package main

import (
	"fmt"
	"net/http"
	"os"
)

// StreamVideo trata de encontrar o arquivo e enviá-lo por meio de buffer
func StreamVideo(w http.ResponseWriter, r *http.Request, videoName string) {
	filePath := "./Videos_OrbisFlow/" + videoName

	// Localizar o arquivo
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("File not found: %s\n", filePath)
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Error 404: Video file not found.")
		return
	}
	defer file.Close() // Garante que o arquivo será fechado ao fim da função

	//Coleta as informações do arquivo (como tamanho e data de modificação)
	fileInfo, err := file.Stat()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//Entrega o conteúdo com suporte a Seek (Barra de tempo)
	http.ServeContent(w, r, fileInfo.Name(), fileInfo.ModTime(), file)
}
