package client

import (
	"crypto/rand"
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

func (h *Handler) GetClients(w http.ResponseWriter, r *http.Request) {

	clients, err := h.service.db.Conn.Query("SELECT client_name, client_id, grant_types, scopes FROM oauth_clients")
	if err != nil {
		http.Error(w, "Failed to fetch clients", http.StatusInternalServerError)
		return
	}
	defer clients.Close()

	var clientList []map[string]interface{}
	for clients.Next() {
		var clientName, clientID, grantTypes, scopes string
		err := clients.Scan(&clientName, &clientID, &grantTypes, &scopes)
		if err != nil {
			http.Error(w, "Failed to scan clients", http.StatusInternalServerError)
			return
		}
		clientList = append(clientList, map[string]interface{}{
			"client_name": clientName,
			"client_id":   clientID,
			"grant_types": strings.Split(grantTypes, ","),
			"scopes":      strings.Split(scopes, ","),
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(clientList)
}

func (h *Handler) StoreClient(w http.ResponseWriter, r *http.Request) {

	reqBody := json.NewDecoder(r.Body)
	var clientBody RequestClientBody
	err := reqBody.Decode(&clientBody)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	// Convert slice fields to TEXT for SQLite
	grantTypesStr := strings.Join(clientBody.Grant_Type, ",")
	var scopeStr string
	switch v := any(clientBody.Scope).(type) {
	case []string:
		scopeStr = strings.Join(v, ",")
	case string:
		scopeStr = v
	default:
		scopeStr = ""
	}
	// Insert client into database
	query := `INSERT INTO oauth_clients (client_name, client_id, client_secret, name, grant_types, scopes) 
			  VALUES (?, ?, ?, ?, ?, ?)`

	clientID := NewRandomClientID()     // Generate a unique client ID
	clientSecret := NewRandomClientID() // Generate a secure client secret

	_, err = h.service.db.Conn.Exec(query,
		clientBody.Client_Name,
		clientID,
		clientSecret,
		clientBody.Client_Name, // or use a separate Name field if available
		grantTypesStr,
		scopeStr,
	)

	if err != nil {
		log.Print(err)
		http.Error(w, "Failed to store client", http.StatusInternalServerError)
		return
	}

	response := ResponseClient{
		Client_Name:  clientBody.Client_Name,
		Client_ID:    clientID,
		ClientSecret: clientSecret,
		Grant_Type:   clientBody.Grant_Type,
		Scope:        clientBody.Scope,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func randomAlphaNumeric(n int) string {
	const alnum = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}
	for i := range b {
		b[i] = alnum[int(b[i])%len(alnum)]
	}
	return string(b)
}

func NewRandomClientID() string {
	return randomAlphaNumeric(12)
}
