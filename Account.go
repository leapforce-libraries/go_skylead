package skylead

import (
	"net/http"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	d_types "github.com/leapforce-libraries/go_skylead/types"
)

type LinkedInUser struct {
	Education                 string                 `json:"education"`
	Experience                string                 `json:"experience"`
	Volunteer                 string                 `json:"volunteer"`
	Id                        int64                  `json:"id"`
	Email                     string                 `json:"email"`
	BusinessEmail             string                 `json:"businessEmail"`
	FullName                  string                 `json:"fullName"`
	Picture                   string                 `json:"picture"`
	Occupation                string                 `json:"occupation"`
	Country                   string                 `json:"country"`
	CurrentCompanyLogo        string                 `json:"currentCompanyLogo"`
	PrimaryIdentifier         string                 `json:"primaryIdentifier"`
	PrimaryIdentifierTypeId   int64                  `json:"primaryIdentifierTypeId"`
	CurrentCompany            string                 `json:"currentCompany"`
	YearsInCurrentCompany     int64                  `json:"yearsInCurrentCompany"`
	TotalCareerPositionsCount int64                  `json:"totalCareerPositionsCount"`
	TotalYearsInCareer        int64                  `json:"totalYearsInCareer"`
	CollegeName               string                 `json:"collegeName"`
	TemporaryField            string                 `json:"temporaryField"`
	Phone                     string                 `json:"phone"`
	Twitter                   string                 `json:"twitter"`
	Website                   string                 `json:"website"`
	Domain                    string                 `json:"domain"`
	ConnectionDegree          string                 `json:"connectionDegree"`
	Activities                string                 `json:"activities"`
	Recommendations           string                 `json:"recommendations"`
	CurrentLinkedinCompanyId  int64                  `json:"currentLinkedinCompanyId"`
	CreatedAt                 d_types.DateTimeString `json:"created_at"`
	UpdatedAt                 d_types.DateTimeString `json:"updated_at"`
}

type Account struct {
	Id                     int64                  `json:"id"`
	Email                  string                 `json:"email"`
	LinkedinUserId         int64                  `json:"linkedinUserId"`
	LinkedinUser           LinkedInUser           `json:"linkedinUser"`
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
