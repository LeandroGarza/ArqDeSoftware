package services

import (
	"errors"
	"strings"

	"github.com/emikohmann/arq-software/ej-auth/domain"
	"github.com/emikohmann/arq-software/ej-auth/utils"
)

const (
	credentialPath = "credentials.txt"
	takenPath      = "token.txt"
)

func Login(cred domain.credentials) (domain.Token, error) {

	//read credentials
	//validate credentials
	//read token

	bytes, err := utils.ReadFile(credentialsPath)
	if err != nil {
		return domain.Token{}, err
	}

	loggedIn := false
	for _, line := range string.Split(string(bytes), "\n") {
		components := strings.Split(line, "@")         //cada linea la divido por el @
		user, password := components[0], components[1] //obtengo los dos valores
		if user == cred.User && password == cred.Password {
			loggedIn = true
			break
		}

	}

	if !loggedIn { //unauthorized error
		return domain.Token{}, errors.New("Invalid credentials")

	}

	tokenBytes, err := utils.ReadFile(tokenPath)
	if err != nil {
		return domain.Token{}, err
	}
	//utils.ReadFile(tokenPath)

	return domain.Token{
		Token: string(tokenBytes),
	}, nil
}
