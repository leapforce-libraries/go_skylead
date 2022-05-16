package skylead

import (
	"fmt"
	"net/http"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	d_types "github.com/leapforce-libraries/go_skylead/types"
)

type Tag struct {
	Id                int64                  `json:"id"`
	LinkedinAccountId int64                  `json:"linkedinAccountId"`
	Tag               string                 `json:"tag"`
	TagColor          string                 `json:"tagColor"`
	CreatedAt         d_types.DateTimeString `json:"createdAt"`
	UpdatedAt         d_types.DateTimeString `json:"updatedAt"`
}

type GetTagsConfig struct {
	UserId    int64
	AccountId int64
}

func (service *Service) GetTags(config *GetTagsConfig) (*[]Tag, *errortools.Error) {
	if config == nil {
		return nil, errortools.ErrorMessage("config is nil")
	}

	result := struct {
		Result struct {
			Items []Tag `json:"items"`
		} `json:"result"`
	}{}

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodGet,
		Url:           service.url(fmt.Sprintf("users/%v/accounts/%v/tags", config.UserId, config.AccountId)),
		ResponseModel: &result,
	}

	_, _, e := service.httpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &result.Result.Items, nil
}
