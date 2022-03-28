package skylead

import (
	"net/http"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	d_types "github.com/leapforce-libraries/go_skylead/types"
)

type Account struct {
	Id                     int64                  `json:"id"`
	Email                  string                 `json:"email"`
	LinkedinUserId         int64                  `json:"linkedinUserId"`
	UserId                 int64                  `json:"userId"`
	AccountPaymentStatusId int64                  `json:"accountPaymentStatusId"`
	SubscriptionId         int64                  `json:"subscriptionId"`
	LinkedinSubscriptionId int64                  `json:"linkedinSubscriptionId"`
	FullName               string                 `json:"fullName"`
	CreatedAt              d_types.DateTimeString `json:"createdAt"`
	UpdatedAt              d_types.DateTimeString `json:"updatedAt"`
	JailPoint              int64                  `json:"jailPoint"`
	IsInJail               bool                   `json:"isInJail"`
	PlanId                 int64                  `json:"planId"`
}

// GetAccounts returns all accounts
//
func (service *Service) GetAccounts() (*[]Account, *errortools.Error) {
	result := struct {
		Result struct {
			Items []Account `json:"items"`
		} `json:"result"`
	}{}

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodGet,
		Url:           service.url("accounts"),
		ResponseModel: &result,
	}

	_, _, e := service.httpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &result.Result.Items, nil
}
