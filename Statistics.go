package skylead

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	d_types "github.com/leapforce-libraries/go_skylead/types"
)

type Statistic struct {
	Date  d_types.DateString `json:"date"`
	Value int64              `json:"value"`
}

type GetStatisticsConfig struct {
	UserId     int64
	AccountId  int64
	From       time.Time
	To         time.Time
	TimeZone   *string
	Curves     []Curve
	CampaignId *int64
}

func (service *Service) GetStatistics(config *GetStatisticsConfig) (map[string][]Statistic, *errortools.Error) {
	if config == nil {
		return nil, errortools.ErrorMessage("config is nil")
	}

	values := url.Values{}
	values.Set("from", fmt.Sprintf("%v", config.From.Unix()))
	values.Set("to", fmt.Sprintf("%v", config.To.Unix()))
	if config.TimeZone != nil {
		values.Set("timeZone", *config.TimeZone)
	} else {
		values.Set("timeZone", service.timeZone)
	}
	if config.CampaignId != nil {
		values.Set("campaignId", fmt.Sprintf("%v", *config.CampaignId))
	}

	curves := strings.ReplaceAll(fmt.Sprintf("%v", config.Curves), " ", ",")

	result := struct {
		Result struct {
			DailyStatistic map[string][]Statistic `json:"dailyStatistic"`
			Total          map[string]int64       `json:"total"`
		} `json:"result"`
	}{}

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodGet,
		Url:           service.url(fmt.Sprintf("users/%v/accounts/%v/statistics?%s&curves=%s", config.UserId, config.AccountId, values.Encode(), curves)),
		ResponseModel: &result,
	}

	_, _, e := service.httpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return result.Result.DailyStatistic, nil
}
