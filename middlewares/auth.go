package middlewares

import (
	"encoding/base64"
	"errors"
	"net/http"
	"strings"

	"github.com/jaycel19/capstone-api/helpers"
	"github.com/jaycel19/capstone-api/services"
	"github.com/jaycel19/capstone-api/util"
)


func RequireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			helpers.WriteJSON(w, http.StatusUnauthorized, helpers.Envelope{"Error": "Missing Authorization Header."})
			return
		}

		splitAuth := strings.Fields(authHeader)
		if len(splitAuth) != 2 || splitAuth[0] != "Basic" {
			helpers.WriteJSON(w, http.StatusUnauthorized, helpers.Envelope{"Error": "Invalid Authorization Header Format. Use Basic Authentication."})
			return
		}

		credentials, err := base64.StdEncoding.DecodeString(splitAuth[1])
		if err != nil {
			helpers.WriteJSON(w, http.StatusUnauthorized, helpers.Envelope{"Error": "Error decoding credentials."})
			return
		}


		user, pass, err := extractCredentials(string(credentials))
		if err != nil || !authenticateUser(user, pass) {
		   helpers.WriteJSON(w, http.StatusUnauthorized, helpers.Envelope{"Error": "Invalid username or password."})
		   return
		}
		next.ServeHTTP(w, r)
	})
}

// Example function for extracting credentials (username:password)
func extractCredentials(credentials string) (string, string, error) {
	split := strings.Split(credentials, ":")
	if len(split) != 2 {
		return "", "", errors.New("Invalid credentials format")
	}
	return split[0], split[1], nil
}

// Example function for basic authentication
func authenticateUser(username, password string) bool {
	var models services.Models
	// Add your authentication logic here
	// Return true if the username and password are valid, otherwise false
	
	adminPayload := services.Admin{
		Username: username,
		Password: password,
	}
	personnelPayload := services.Personnel{
		Username: username,
		Password: password,
	}
	
	personnel, pErr := models.Personnel.PersonnelLogin(personnelPayload)
	
	admin, aErr := models.Admin.AdminLogin(adminPayload)
	if pErr != nil && aErr != nil{
		return false
	}
	
	if pErr == nil {
		err := util.CheckPassword(password, personnel.Password)
		if err != nil {
			return false
		}
	}

	if aErr == nil {
		err := util.CheckPassword(password, admin.Password)
		if err != nil {
			return false
		}
	}
	return true
}


