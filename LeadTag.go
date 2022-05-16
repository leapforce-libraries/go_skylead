package skylead

import (
	"fmt"
	"net/http"
	"strings"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
)

type LeadTag struct {
	Id     int64 `json:"id"`
	TagId  int64 `json:"tagId"`
	LeadId int64 `json:"leadId"`
}

type GetLeadTagsConfig struct {
	UserId    int64
	AccountId int64
	LeadIds   []string
}

func (service *Service) GetLeadTags(config *GetLeadTagsConfig) (*[]LeadTag, *errortools.Error) {
	if config == nil {
		return nil, errortools.ErrorMessage("config is nil")
	}

	result := struct {
		Result struct {
			Items []LeadTag `json:"items"`
		} `json:"result"`
	}{}

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodGet,
		Url:           service.url(fmt.Sprintf("users/%v/accounts/%v/leads/tags?leadIds=[%s]", config.UserId, config.AccountId, strings.Join(config.LeadIds, ","))),
		ResponseModel: &result,
	}

	_, _, e := service.httpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &result.Result.Items, nil
}
