// main.go
package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "firstprojectgo/handlers"  // Importa o handler de usuário
)

func main() {
    r := mux.NewRouter()

    // Rota para lidar com o cadastro de usuários
    r.HandleFunc("/users", handlers.CreateUser).Methods("POST")
    
    // Porta do servidor
    log.Println("Servidor iniciado na porta 8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}
