// handlers/user_handler.go
package handlers

import (
    "encoding/json"
    "net/http"
    "firstprojectgo/models"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
    var user models.User

    // Decodifica o corpo da requisição para o struct User
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        http.Error(w, "Dados inválidos", http.StatusBadRequest)
        return
    }

    // Para simplificar, estamos apenas retornando os dados recebidos como JSON
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(user)
}
