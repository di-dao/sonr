package middleware

import (
	"encoding/json"

	"github.com/di-dao/sonr/internal/models"
	"github.com/go-webauthn/webauthn/protocol"
	"github.com/labstack/echo/v4"
)

func ParseCredentialFromForm(c echo.Context) (*models.Credential, error) {
	// Get the serialized credential data from the form
	credentialDataJSON := c.FormValue("credentialData")

	// Deserialize the JSON into a temporary struct
	var tempCredential protocol.CredentialCreationResponse
	err := json.Unmarshal([]byte(credentialDataJSON), &tempCredential)
	if err != nil {
		return nil, err
	}

	// Parse the CredentialCreationResponse
	parsedData, err := tempCredential.Parse()
	if err != nil {
		return nil, err
	}

	// Create the Credential
	credential := models.MakeNewCredential(parsedData)

	// Set additional fields
	credential.DisplayName = tempCredential.ID // You might want to set this to a more meaningful value
	credential.Origin = c.Request().Host       // Set the origin to the current host
	credential.Controller = ""                 // Set this to the appropriate controller value

	return credential, nil
}
