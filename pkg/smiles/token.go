package smiles

import "github.com/go-resty/resty/v2"

type TokenRequest struct {
	Audience     string `json:"audience"`
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	GrantType    string `json:"grant_type"`
}

type TokenResult struct {
	AccessToken string `json:"access_token"`
}

func GetToken() (string, error) {
	client := resty.New()

	tokenRequest := TokenRequest{
		Audience:     "https://smiles.api",
		ClientId:     "2gpRUWTOBFgi2uypotR3gBUhCtVuYs2G",
		ClientSecret: "U2_qer1iuZBhdIS7IlUXlvkdesl98Yjv38OmW5eu__XlXz-3aWLhAFPVNcig3V3e",
		GrantType:    "client_credentials",
	}

	tokenResult := TokenResult{}

	_, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetResult(&tokenResult).
		SetBody(tokenRequest).
		Post("https://apigw.smiles.com.br/b2b/partner/oauth/token")

	if err != nil {
		return "", err
	}

	return tokenResult.AccessToken, nil
}