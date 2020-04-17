package sdk

import (
	"errors"
	"fmt"
)

// RequestExtraInfoExtraFee :
type RequestExtraInfoExtraFee struct {
	Type        string `json:"type"`
	ReferenceID string `json:"referenceId"`
	Amount      uint64 `json:"amount"`
}

// RequestCreatePaymentCheckout :
type RequestCreatePaymentCheckout struct {
	Order struct {
		ID             string `json:"id"`
		Title          string `json:"title"`
		Detail         string `json:"detail"`
		AdditionalData string `json:"additionalData"`
		CurrencyType   string `json:"currencyType"`
		Amount         uint64 `json:"amount"`
	} `json:"order"`
	Method        []string `json:"method"`
	Type          string   `json:"type"`
	StoreID       string   `json:"storeId"`
	RedirectURL   string   `json:"redirectUrl"`
	NotifyURL     string   `json:"notifyUrl"`
	LayoutVersion string   `json:"layoutVersion"`
	ExtraInfo     struct {
		ExtraFee []RequestExtraInfoExtraFee `json:"extraFee"`
	} `json:"extraInfo"`
}

// ResponseCreatePaymentCheckout :
type ResponseCreatePaymentCheckout struct {
	Item struct {
		CheckoutID string `json:"checkoutId"`
		URL        string `json:"url"`
	} `json:"item"`
	Code string `json:"code"`
	Err  *Error `json:"error"`
}

// CreatePaymentCheckout :
func (c Client) CreatePaymentCheckout(request RequestCreatePaymentCheckout) (*ResponseCreatePaymentCheckout, error) {
	if c.err != nil {
		return nil, c.err
	}

	method := pathAPICreatePaymentCheckoutURL.method
	requestURL := c.prepareAPIURL(pathAPICreatePaymentCheckoutURL)

	response := new(ResponseCreatePaymentCheckout)
	if err := c.httpAPI(method, requestURL, request, response); err != nil {
		return nil, err
	}

	if response.Err != nil {
		return response, errors.New(response.Err.Message)
	}

	return response, nil
}

// ResponseGetQRCodeByCheckoutID :
type ResponseGetQRCodeByCheckoutID struct {
	Item struct {
		QRCodeImageBase64 string `json:"qrCodeImageBase64"`
		URL               string `json:"url"`
	} `json:"item"`
	Code string `json:"code"`
	Err  *Error `json:"error"`
}

// GetQRCodeByCheckoutID :
func (c Client) GetQRCodeByCheckoutID(checkoutID, walletMethod string) (*ResponseGetQRCodeByCheckoutID, error) {
	if c.err != nil {
		return nil, c.err
	}

	method := pathAPIGetQRCodeByCheckoutIDURL.method
	requestURL := c.prepareAPIURL(pathAPIGetQRCodeByCheckoutIDURL)

	response := new(ResponseGetQRCodeByCheckoutID)
	if err := c.httpAPI(method, fmt.Sprintf("%s?checkoutId=%s&method=%s", requestURL, checkoutID, walletMethod), nil, response); err != nil {
		return nil, err
	}

	if response.Err != nil {
		return response, errors.New(response.Err.Message)
	}

	return response, nil
}