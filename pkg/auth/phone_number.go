package auth

import (
	"fmt"
	"github.com/ad/domru/handlers"
	"net/http"
	"regexp"
)

const (
	BaseURL          = "https://myhome.novotelecom.ru"
	phoneNumberRegex = `^\+7\d{10}$`
)

type SmsCodeGetter func() (string, error)

type PhoneNumberAuthenticator struct {
	phoneNumber string
	getSmsCode  SmsCodeGetter
	operatorID  int

	userAccounts []handlers.Account
}

func NewPhoneNumberAuthenticator(phoneNumber string, getSmsCode SmsCodeGetter) *PhoneNumberAuthenticator {
	return &PhoneNumberAuthenticator{
		phoneNumber: phoneNumber,
		getSmsCode:  getSmsCode,
	}
}

func (a *PhoneNumberAuthenticator) isPhoneNumberValid() bool {
	r, err := regexp.Compile(phoneNumberRegex)
	if err != nil {
		panic(err)
	}

	return r.MatchString(a.phoneNumber)
}

func (a *PhoneNumberAuthenticator) Authenticate() (AuthenticationResponse, error) {
	if !a.isPhoneNumberValid() {
		return AuthenticationResponse{}, fmt.Errorf("phone number is invalid format. It should be +7XXXXXXXXXX")
	}

	accounts, err := a.getUserAccounts()
	if err != nil {
		return AuthenticationResponse{}, fmt.Errorf("failed to get user accounts: %w", err)
	}

	// send sms code request
	a.requestConfirmationCode(accounts[0])

	smsCode, err := a.getSmsCode()
	if err != nil {
		return AuthenticationResponse{}, fmt.Errorf("failed to get sms code: %w", err)
	}

	response, err := a.sendConfirmationCode(smsCode)
	if err != nil {
		return AuthenticationResponse{}, fmt.Errorf("failed to authenticate: %w", err)
	}

	return response, nil
}

func (a *PhoneNumberAuthenticator) getUserAccounts() ([]handlers.Account, error) {
	authUrl := fmt.Sprintf("%s/auth/v2/login/%s", BaseURL, a.phoneNumber)

	var accounts []handlers.Account
	err := SendRequest(authUrl, http.MethodGet, nil, &accounts)
	if err != nil {
		return []handlers.Account{}, fmt.Errorf("failed to check account: %w", err)
	}
	if len(accounts) == 0 {
		return []handlers.Account{}, fmt.Errorf("empty response for accounts found for phone number %s", a.phoneNumber)
	}

	return accounts, nil
}

func (a *PhoneNumberAuthenticator) requestConfirmationCode(account handlers.Account) error {
	confirmUrl := fmt.Sprintf("%s/auth/v2/confirmation/%s", BaseURL, a.phoneNumber)
	confirmRequest := &handlers.Account{
		Address:      account.Address,
		OperatorID:   account.OperatorID,
		ProfileID:    account.ProfileID,
		SubscriberID: account.SubscriberID,
	}

	err := SendRequest(confirmUrl, http.MethodPost, confirmRequest, nil)
	if err != nil {
		return fmt.Errorf("failed to request confirmation code: %w", err)
	}
	return nil
}

func (a *PhoneNumberAuthenticator) sendConfirmationCode(smsCode string) (AuthenticationResponse, error) {
	confirmUrl := fmt.Sprintf("%s/auth/v2/auth/%s/confirmation", BaseURL, a.phoneNumber)
	confirmRequest := &ConfirmationRequest{
		Confirm1:     smsCode,
		Confirm2:     smsCode,
		Login:        a.phoneNumber,
		OperatorID:   1,
		ProfileID:    "",
		SubscriberID: "0",
	}
	var confirmResponse AuthenticationResponse
	err := SendRequest(confirmUrl, http.MethodPost, confirmRequest, &confirmResponse)
	if err != nil {
		return AuthenticationResponse{}, fmt.Errorf("failed to send confirmation code: %w", err)
	}
	return confirmResponse, nil
}