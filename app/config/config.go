package config

import (
	"encoding/base64"
	"os"
)

type Configuration struct {
	GoogleServiceKey []byte
	GoogleAuthURL    string
	SpreadsheetId    string
}

func LoadConfiguration() (*Configuration, error) {
	cfg := Configuration{}
	keyBytes, err := base64.StdEncoding.DecodeString(os.Getenv("GOOGLE_SERVICE_KEY"))
	if err != nil {
		return nil, err
	}

	cfg.GoogleServiceKey = keyBytes

	cfg.SpreadsheetId = os.Getenv("SPREAD_SHEET_ID")

	return &cfg, nil
}
