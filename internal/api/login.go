package api

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
	"os"
)

func Login(w http.ResponseWriter, r *http.Request) {

	sendToAuth(w, r)
}

func sendToAuth(w http.ResponseWriter, r *http.Request) {
	authDomain := os.Getenv("AUTH_DOMAIN")
	clientID := os.Getenv("CLIENT_ID")
	serverDomain := os.Getenv("SERVER_DOMAIN")

	if authDomain == "" || clientID == "" || serverDomain == "" {
		slog.Error("Error while logging user in", "err", errors.New("environment variables are not set"))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	params := url.Values{}
	params.Add("client_id", clientID)
	params.Add("redirect_uri", fmt.Sprintf("%s/api/auth", serverDomain))
	params.Add("response_type", "code")
	params.Add("scope", "name slack_id")
	authURL := fmt.Sprintf("%s/oauth/authorize?%s", authDomain, params.Encode())

	http.Redirect(w, r, authURL, http.StatusSeeOther)
}
