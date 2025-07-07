package response

import (
	"github.com/LGYS07/topsdk/ability2939/domain"
)

type AlibabaAlihealthDrugLsydGetentinfonewResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   resultModel
	*/
	Result domain.AlibabaAlihealthDrugLsydGetentinfonewResultModel `json:"result,omitempty" `
}
