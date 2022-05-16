package skylead

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
)

type Message struct {
	Id             int64           `json:"id"`
	Attachments    json.RawMessage `json:"attachments"`
	LinkedinUserId int64           `json:"linkedinUserId"`
	Account        int64           `json:"account"`
	Thread         string          `json:"thread"`
	Message        string          `json:"message"`
	Receiver       int64           `json:"receiver"`
	Dashboard      int64           `json:"dashboard"`
	CampaignId     int64           `json:"campaignId"`
	SeenAt         int64           `json:"seenAt"`
	MessageStatus  string          `json:"messageStatus"`
	MessageType    string          `json:"messageType"`
}

type GetMessagesConfig struct {
	UserId    int64
	AccountId int64
	ThreadIds []string
}

func (service *Service) GetMessages(config *GetMessagesConfig) (map[string][]Message, *errortools.Error) {
	if config == nil {
		return nil, errortools.ErrorMessage("config is nil")
	}

	result := struct {
		Result struct {
			Items map[string][]Message `json:"items"`
		} `json:"result"`
	}{}

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodGet,
		Url:           service.url(fmt.Sprintf("users/%v/accounts/%v/conversations/threads?threads=[\"%s\"]", config.UserId, config.AccountId, strings.Join(config.ThreadIds, "\",\""))),
		ResponseModel: &result,
	}

	_, _, e := service.httpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return result.Result.Items, nil
}
