package skylead

import (
	"fmt"
	"net/http"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
)

const (
	apiName         string = "Skylead"
	apiPath         string = "https://api.multilead.io/api/open-api/v1"
	defaultTimeZone string = "Europe/Belgrade"
)

type Curve int64

const (
	CurveViews               Curve = 1
	CurveFollows             Curve = 2
	CurveConnectionsSent     Curve = 3
	CurveMessagesSent        Curve = 4
	CurveEmailsSent          Curve = 5
	CurveConnectionsAccepted Curve = 6
	CurveRepliesReceived     Curve = 7
)

type Service struct {
	apiKey      string
	timeZone    string
	httpService *go_http.Service
}

type ServiceConfig struct {
	ApiKey   string
	TimeZone *string
}

func NewService(serviceConfig *ServiceConfig) (*Service, *errortools.Error) {
	if serviceConfig == nil {
		return nil, errortools.ErrorMessage("ServiceConfig must not be a nil pointer")
	}

	if serviceConfig.ApiKey == "" {
		return nil, errortools.ErrorMessage("Service ApiKey not provided")
	}

	httpService, e := go_http.NewService(&go_http.ServiceConfig{})
	if e != nil {
		return nil, e
	}

	timeZone := defaultTimeZone
	if serviceConfig.TimeZone != nil {
		timeZone = *serviceConfig.TimeZone
	}

	return &Service{
		apiKey:      serviceConfig.ApiKey,
		timeZone:    timeZone,
		httpService: httpService,
	}, nil
}

func (service *Service) httpRequest(requestConfig *go_http.RequestConfig) (*http.Request, *http.Response, *errortools.Error) {
	// add authentication header
	header := http.Header{}
	header.Set("Authorization", service.apiKey)
	(*requestConfig).NonDefaultHeaders = &header

	// add error model
	errorResponse := ErrorResponse{}
	(*requestConfig).ErrorModel = &errorResponse

	request, response, e := service.httpService.HttpRequest(requestConfig)
	if errorResponse.Error.Message != "" {
		e.SetMessage(errorResponse.Error.Message)
	}

	return request, response, e
}

func (service *Service) url(path string) string {
	return fmt.Sprintf("%s/%s", apiPath, path)
}

func (service *Service) ApiName() string {
	return apiName
}

func (service *Service) ApiKey() string {
	return service.apiKey
}

func (service *Service) ApiCallCount() int64 {
	return service.httpService.RequestCount()
}

func (service *Service) ApiReset() {
	service.httpService.ResetRequestCount()
}
