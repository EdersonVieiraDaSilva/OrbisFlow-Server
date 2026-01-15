package main

import (
	"net/http"
	"path/filepath"
)

// Garante que o usuário só acesse arquivos dentro da pasta designada
func ValidatePath(videoName string) string {
	// filepath.Base remove qualquer tentativa de usar "../" ou caminhos absolutos
	return filepath.Base(videoName)
}

// Cria uma barreira de usuário e senha
func BasicAuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, pass, ok := r.BasicAuth()

		// Defina credenciais para teste público
		const adminUser = "generico_admin"
		const adminPass = "generico_2026"

		if !ok || user != adminUser || pass != adminPass {
			w.Header().Set("WWW-Authenticate", `Basic realm="OrbisFlow Restricted"`)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// Se estiver ok, prossegue para a função original (StreamVideo)
		next(w, r)
	}
}
