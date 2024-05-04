package payhere

import (
	"encoding/json"
	"time"
)

type Client struct {
	Config *Config
}

type Config struct {
	BaseUrl string `json:"base_url"`
	Token   string `json:"token"`
}

type RequestError struct {
	Code, Message string
	Fields        *map[string]interface{} `json:"fields"`
}

type Response struct {
	Status  interface{}    `json:"status"`
	Success bool           `json:"success"`
	Data    interface{}    `json:"data,omitempty"`
	Error   *ResponseError `json:"error,omitempty"`
}

type ResponseError struct {
	Code    string                  `json:"code"`
	Message string                  `json:"message"`
	Fields  *map[string]interface{} `json:"fields"`
}

type Balance struct {
	Balance   json.Number `json:"balance"`
	UpdatedAt time.Time   `json:"updated_at"`
}

type CreateDisbursementRequest struct {
	MerchantTrxId          string `json:"merchant_trx_id"`
	RecipientName          string `json:"recipient_name"`
	RecipientBankCode      string `json:"recipient_bank_code"`
	RecipientAccountNumber string `json:"recipient_account_number"`
	Amount                 int64  `json:"amount"`
	Description            string `json:"description,omitempty"`
	CallbackUrl            string `json:"callback_url,omitempty"`
}

type Disbursement struct {
	Uuid                   string    `json:"uuid"`
	MerchantTrxId          string    `json:"merchant_trx_id"`
	RecipientName          string    `json:"recipient_name"`
	RecipientBankCode      string    `json:"recipient_bank_code"`
	RecipientAccountNumber string    `json:"recipient_account_number"`
	Amount                 int64     `json:"amount"`
	CreatedTime            time.Time `json:"created_time"`
	Status                 string    `json:"status"`
	StatusMessage          string    `json:"status_message"`
	Description            string    `json:"description"`
	CallbackUrl            string    `json:"callback_url"`
	DisbursedTime          time.Time `json:"disbursed_time"`
}
