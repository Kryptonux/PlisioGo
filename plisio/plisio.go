package plisio

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"time"
)

type PlisioClient struct {
	APIKey      string
	Email       string
	CallbackURL string
}

type InvoiceResponse struct {
	TxnID                 string `json:"txn_id"`
	InvoiceURL            string `json:"invoice_url"`
	Amount                string `json:"amount"`
	PendingAmount         string `json:"pending_amount"`
	WalletHash            string `json:"wallet_hash"`
	PsysCID               string `json:"psys_cid"`
	Currency              string `json:"currency"`
	Status                string `json:"status"`
	SourceCurrency        string `json:"source_currency"`
	SourceRate            string `json:"source_rate"`
	ExpireUTC             int    `json:"expire_utc"`
	ExpectedConfirmations string `json:"expected_confirmations"`
	QRCode                string `json:"qr_code"`
	VerifyHash            string `json:"verify_hash"`
	InvoiceCommission     string `json:"invoice_commission"`
	InvoiceSum            string `json:"invoice_sum"`
	InvoiceTotalSum       string `json:"invoice_total_sum"`
}

type Params struct {
	OrderNumber    string `json:"order_number"`
	OrderName      string `json:"order_name"`
	SourceAmount   string `json:"source_amount"`
	SourceCurrency string `json:"source_currency"`
	Amount         string `json:"amount"`
	SourceRate     string `json:"source_rate"`
	Email          string `json:"email"`
}

func NewPlisioClient(apiKey, email, callbackURL string) *PlisioClient {
	return &PlisioClient{
		APIKey:      apiKey,
		Email:       email,
		CallbackURL: callbackURL,
	}
}

func (c *PlisioClient) CreateInvoice(sourceCurrency string, sourceAmount float64, orderNumber string, currency, orderName string) (string, error) {
	baseURL := "https://plisio.net/api/v1/invoices/new"
	params := url.Values{}
	params.Set("source_currency", sourceCurrency)
	params.Set("source_amount", fmt.Sprintf("%.4f", sourceAmount))
	params.Set("order_number", orderNumber)
	params.Set("currency", currency)
	params.Set("email", c.Email)
	params.Set("order_name", orderName)
	params.Set("callback_url", c.CallbackURL)
	params.Set("api_key", c.APIKey)

	reqURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

	resp, err := http.Get(reqURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("API request failed with status code: %d", resp.StatusCode)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(bodyBytes), nil
}

func GenerateTransactionID(length int) string {
	source := rand.NewSource(time.Now().UnixNano())

	random := rand.New(source)

	characters := "abcdefghijklmnopqrstuvwxyz1234567890"

	var result string

	for i := 0; i < length; i++ {
		randomIndex := random.Intn(len(characters))
		result += string(characters[randomIndex])
	}

	return result
}
