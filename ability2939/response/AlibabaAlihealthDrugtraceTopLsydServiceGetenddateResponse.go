package response

import (
	"github.com/LGYS07/topsdk/ability2939/domain"
)

type AlibabaAlihealthDrugtraceTopLsydServiceGetenddateResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   接口返回
	*/
	Result domain.AlibabaAlihealthDrugtraceTopLsydServiceGetenddateResultModel `json:"result,omitempty" `
}
