package main

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

func VerifyCaptcha(token string, secretKey string) (bool, error) {
	type turnstileResponse struct {
    	Success bool `json:"success"`
	}	

	formData := url.Values{}
	formData.Set("secret", secretKey)
	formData.Set("response", token)

	response, err := http.Post("https://challenges.cloudflare.com/turnstile/v0/siteverify", "application/x-www-form-urlencoded", strings.NewReader(formData.Encode()))
	if err != nil {
		
		return false, err
	}

	decoder := json.NewDecoder(response.Body)
	params := turnstileResponse{}
	err = decoder.Decode(&params)
	if err != nil {
		return false, err
	}

	return params.Success, nil
}