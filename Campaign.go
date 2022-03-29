package skylead

import (
	"fmt"
	"net/http"
	"net/url"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	d_types "github.com/leapforce-libraries/go_skylead/types"
	go_types "github.com/leapforce-libraries/go_types"
)

const limitDefault int64 = 30

type CampaignState int64

const (
	CampaignStateActive   CampaignState = 1
	CampaignStateDraft    CampaignState = 2
	CampaignStateArchived CampaignState = 3
)

type SortOrder string

const (
	SortOrderAsc  SortOrder = "ASC"
	SortOrderDesc SortOrder = "DESC"
)

type Campaign struct {
	Id                int64                  `json:"id"`
	LinkedinAccountId int64                  `json:"linkedinAccountId"`
	Name              string                 `json:"name"`
	CreatedAt         d_types.DateTimeString `json:"createdAt"`
	UpdatedAt         d_types.DateTimeString `json:"updatedAt"`
	StateId           int64                  `json:"stateId"`
	CampaignStats     CampaignStats          `json:"campaignStats"`
	LeadSources       []LeadSource           `json:"leadSources"`
}

type CampaignStats struct {
	CampaignId                 int64                   `json:"campaignId"`
	ProfileViewsMade           int64                   `json:"profileViewsMade"`
	InmailsSent                int64                   `json:"inmailsSent"`
	EmailsSent                 int64                   `json:"emailsSent"`
	ConnectionsRequested       int64                   `json:"connectionsRequested"`
	MessagesSent               int64                   `json:"messagesSent"`
	ConnectionRequestsAccepted int64                   `json:"connectionRequestsAccepted"`
	ConnectionReplies          int64                   `json:"connectionReplies"`
	ResponseRate               *go_types.Float64String `json:"responseRate"`
	AcceptanceRate             *go_types.Float64String `json:"acceptanceRate"`
	OpenRate                   *go_types.Float64String `json:"openRate"`
	ClickRate                  *go_types.Float64String `json:"clickRate"`
	TotalLeads                 int64                   `json:"totalLeads"`
	IsActive                   bool                    `json:"isActive"`
	TotalFollowCount           int64                   `json:"totalFollowCount"`
	EmailsVerified             int64                   `json:"emailsVerified"`
	BounceRate                 *go_types.Float64String `json:"bounceRate"`
	EmailsBounced              int64                   `json:"emailsBounced"`
}

type LeadSource struct {
	Id             int64  `json:"id"`
	Dashboard      int64  `json:"dashboard"`
	GeneralStatus  string `json:"generalStatus"`
	LeadSourceType string `json:"leadSourceType"`
	LeadSourceUrl  string `json:"leadSourceUrl"`
	PageNumber     int64  `json:"pageNumber"`
}

type GetCampaignsConfig struct {
	UserId        int64
	AccountId     int64
	CampaignState *CampaignState
	SortOrder     *SortOrder
	SortColumn    *string
	Limit         *int64
}

func (service *Service) GetCampaigns(config *GetCampaignsConfig) (*[]Campaign, *errortools.Error) {
	if config == nil {
		return nil, errortools.ErrorMessage("config is nil")
	}

	values := url.Values{}
	if config.CampaignState != nil {
		values.Set("campaignState", fmt.Sprintf("%v", int64(*config.CampaignState)))

	}
	if config.SortOrder != nil {
		values.Set("sortOrder", string(*config.SortOrder))

	}
	if config.SortOrder != nil {
		values.Set("sortColumn", *config.SortColumn)
	}
	limit := limitDefault
	if config.Limit != nil {
		limit = *config.Limit
	}
	values.Set("limit", fmt.Sprintf("%v", limit))
	offset := int64(0)

	campaigns := []Campaign{}

	for {
		values.Set("offset", fmt.Sprintf("%v", offset))

		result := struct {
			Result struct {
				Count     int64      `json:"count"`
				Campaigns []Campaign `json:"campaigns"`
			} `json:"result"`
		}{}

		requestConfig := go_http.RequestConfig{
			Method:        http.MethodGet,
			Url:           service.url(fmt.Sprintf("users/%v/accounts/%v/campaigns?%s", config.UserId, config.AccountId, values.Encode())),
			ResponseModel: &result,
		}

		_, _, e := service.httpRequest(&requestConfig)
		if e != nil {
			return nil, e
		}

		campaigns = append(campaigns, result.Result.Campaigns...)

		if result.Result.Count < limit {
			break
		}

		offset += limit
	}

	return &campaigns, nil
}
