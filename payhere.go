package payhere

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/vannleonheart/goutil"
)

func New(config *Config) *Client {
	return &Client{Config: config}
}

func (c *Client) GetBalance() (*Balance, error) {
	var resp Response

	url := fmt.Sprintf("%s/%s", c.Config.BaseUrl, UrlGetBalance)
	headers := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": fmt.Sprintf("Bearer %s", c.Config.Token),
	}

	if _, err := goutil.SendHttpGet(url, nil, &headers, &resp); err != nil {
		return nil, err
	}

	if !resp.Success {
		if resp.Error != nil {
			return nil, &RequestError{Code: resp.Error.Code, Message: resp.Error.Message}
		}

		return nil, errors.New("error when get balance")
	}

	if resp.Data == nil {
		return nil, errors.New("invalid response")
	}

	var result Balance

	byResult, err := json.Marshal(resp.Data)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(byResult, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) CreateDisbursement(data CreateDisbursementRequest) (*Disbursement, error) {
	var resp Response

	url := fmt.Sprintf("%s/%s", c.Config.BaseUrl, UrlDisbursement)
	headers := map[string]string{
		"Content-Type":  "application/x-www-form-urlencoded",
		"Authorization": fmt.Sprintf("Bearer %s", c.Config.Token),
	}

	if _, err := goutil.SendHttpPost(url, data, &headers, &resp); err != nil {
		return nil, err
	}

	if !resp.Success {
		if resp.Error != nil {
			requestError := RequestError{Code: resp.Error.Code, Message: resp.Error.Message}

			if resp.Error.Fields != nil {
				requestError.Fields = resp.Error.Fields
			}

			return nil, &requestError
		}

		return nil, errors.New("error when create disbursement")
	}

	if resp.Data == nil {
		return nil, errors.New("invalid response")
	}

	var result Disbursement

	byResult, err := json.Marshal(resp.Data)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(byResult, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) InquiryDisbursementByUuid(uuid string) (*Disbursement, error) {
	var resp Response

	url := fmt.Sprintf("%s/%s/%s", c.Config.BaseUrl, UrlInquiryDisbursement, uuid)
	headers := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": fmt.Sprintf("Bearer %s", c.Config.Token),
	}

	if _, err := goutil.SendHttpGet(url, nil, &headers, &resp); err != nil {
		return nil, err
	}

	if !resp.Success {
		if resp.Error != nil {
			return nil, &RequestError{Code: resp.Error.Code, Message: resp.Error.Message}
		}

		return nil, errors.New("error when inquiry disbursement by uuid")
	}

	if resp.Data == nil {
		return nil, errors.New("invalid response")
	}

	var result Disbursement

	byResult, err := json.Marshal(resp.Data)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(byResult, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) InquiryDisbursementByMerchantTrx(merchantTrx string) (*Disbursement, error) {
	var resp Response

	url := fmt.Sprintf("%s/%s/%s", c.Config.BaseUrl, UrlInquiryDisbursementByMerchantTrx, merchantTrx)
	headers := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": fmt.Sprintf("Bearer %s", c.Config.Token),
	}

	if _, err := goutil.SendHttpGet(url, nil, &headers, &resp); err != nil {
		return nil, err
	}

	if !resp.Success {
		if resp.Error != nil {
			return nil, &RequestError{Code: resp.Error.Code, Message: resp.Error.Message}
		}

		return nil, errors.New("error when inquiry disbursement by uuid")
	}

	if resp.Data == nil {
		return nil, errors.New("invalid response")
	}

	var result Disbursement

	byResult, err := json.Marshal(resp.Data)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(byResult, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
