package controllers

import (
    "encoding/json"
    "net/http"

    "github.com/jeffrysusilo/repotes/golang-api/config"
    "github.com/google/uuid"
)

func CreateTerminal(w http.ResponseWriter, r *http.Request) {
    type TerminalRequest struct {
        Name     string `json:"name"`
        Location string `json:"location"`
    }

    var req TerminalRequest
    json.NewDecoder(r.Body).Decode(&req)

    id := uuid.New()
    _, err := config.DB.Exec(`INSERT INTO terminals(terminal_id, name, location, created_at) VALUES($1,$2,$3,NOW())`,
        id, req.Name, req.Location)

    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(map[string]string{
        "terminal_id": id.String(),
    })
}
