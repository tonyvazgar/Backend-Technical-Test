package infrastructure

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

var secrets struct {
	Type         string
	ProjectID    string
	PrivateKeyID string
	PrivateKey   string
	ClientEmail  string
	ClientID     string
	AuthURI      string
	TokenURI     string
	AuthProvider string
	Client       string
}

type FireBaseAuth struct {
	Type         string `json:"type"`
	ProjectID    string `json:"project_id"`
	PrivateKeyID string `json:"private_key_id"`
	PrivateKey   string `json:"private_key"`
	ClientEmail  string `json:"client_email"`
	ClientID     string `json:"client_id"`
	AuthURI      string `json:"auth_uri"`
	TokenURI     string `json:"token_uri"`
	AuthProvider string `json:"auth_provider_x509_cert_url"`
	Client       string `json:"client_x509_cert_url"`
}

func InitFirebase() (*firestore.Client, error) {
	ctx := context.Background()
	formattedPrivateKey := fmt.Sprintf("%s", secrets.PrivateKey)
	formattedPrivateKey = strings.Replace(formattedPrivateKey, `\n`, "\n", -1)
	cred := FireBaseAuth{
		Type:         secrets.Type,
		ProjectID:    secrets.ProjectID,
		PrivateKeyID: secrets.PrivateKeyID,
		PrivateKey:   formattedPrivateKey,
		ClientEmail:  secrets.ClientEmail,
		ClientID:     secrets.ClientID,
		AuthURI:      secrets.AuthURI,
		TokenURI:     secrets.TokenURI,
		AuthProvider: secrets.AuthProvider,
		Client:       secrets.Client,
	}

	b, err := json.Marshal(cred)
	if err != nil {
		return nil, err
	}

	opt := option.WithCredentialsJSON(b)
	client, err := firestore.NewClient(ctx, secrets.ProjectID, opt)
	if err != nil {
		return nil, err
	}

	return client, nil
}
