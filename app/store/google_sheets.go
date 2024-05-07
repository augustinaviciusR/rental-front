package store

import (
	"context"
	"fmt"
	"sync"

	"github.com/Pervadinti/rental-front/app/model"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

const MinimumRecordLength = 2 // declared constant

type GoogleSheetsStore struct {
	sheet         *sheets.Service
	mu            sync.RWMutex
	spreadsheetID string
}

func NewGoogleSheetsStore(spreadsheetID string, jsonKey []byte) (*GoogleSheetsStore, error) {
	if spreadsheetID == "" {
		return nil, fmt.Errorf("GoogleSheetsStore: provided SpreadsheetId is empty")
	}
	srv, err := createSheetsService(jsonKey)
	if err != nil {
		return nil, fmt.Errorf("failed to create Google Sheets service: %w", err)
	}
	return &GoogleSheetsStore{
		sheet:         srv,
		mu:            sync.RWMutex{},
		spreadsheetID: spreadsheetID,
	}, nil
}

func createSheetsService(jsonKey []byte) (*sheets.Service, error) {
	config, err := google.JWTConfigFromJSON(jsonKey, "https://www.googleapis.com/auth/spreadsheets")
	if err != nil {
		return nil, fmt.Errorf("failed to obtain JWT: %w", err)
	}
	client := config.Client(context.Background())
	srv, err := sheets.NewService(context.Background(), option.WithHTTPClient(client))
	if err != nil {
		return nil, fmt.Errorf("failed to create new sheets client: %w", err)
	}

	return srv, nil
}

func (g *GoogleSheetsStore) StoreNewInquiry(ctx context.Context, order *model.InquiryForm) error {
	record := orderToSheetRecord(order)
	g.mu.Lock()
	_, err := g.sheet.Spreadsheets.Values.Append(g.spreadsheetID, "INQUIRY", &sheets.ValueRange{
		Values: [][]interface{}{record},
	}).Context(ctx).ValueInputOption("USER_ENTERED").Do()
	g.mu.Unlock()
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println(err.Error())
		fmt.Println(err.Error())
	}
	return err
}

func orderToSheetRecord(inquiry *model.InquiryForm) []interface{} {
	return []interface{}{
		inquiry.Name,
		inquiry.Email,
		inquiry.Description,
	}
}
