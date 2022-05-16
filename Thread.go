package skylead

import (
	"fmt"
	"net/http"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
)

type Lead struct {
	Id                      int64 `json:"id"`
	IsDeclined              int64 `json:"isDeclined"`
	BackToCampaignTimestamp int64 `json:"backToCampaignTimestamp"`
	NextStepId              int64 `json:"nextStepId"`
}

type Thread struct {
	Account      int64        `json:"account"`
	Thread       string       `json:"thread"`
	Dashboard    int64        `json:"dashboard"`
	CampaignId   int64        `json:"threadId"`
	Lead         Lead         `json:"lead"`
	LinkedInUser LinkedInUser `json:"linkedinUser"`
}

type GetThreadsConfig struct {
	UserId     int64
	AccountId  int64
	CampaignId int64
}

func (service *Service) GetThreads(config *GetThreadsConfig) (*[]Thread, *errortools.Error) {
	if config == nil {
		return nil, errortools.ErrorMessage("config is nil")
	}

	result := struct {
		Result struct {
			Items []Thread `json:"items"`
		} `json:"result"`
	}{}

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodGet,
		Url:           service.url(fmt.Sprintf("users/%v/accounts/%v/campaigns/%v/messages", config.UserId, config.AccountId, config.CampaignId)),
		ResponseModel: &result,
	}

	_, _, e := service.httpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &result.Result.Items, nil
}
