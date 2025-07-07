package response

import (
	"github.com/LGYS07/topsdk/ability2939/domain"
)

type AlibabaAlihealthDrugtraceTopLsydQueryUpbillcodeResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   接口返回model
	*/
	Result domain.AlibabaAlihealthDrugtraceTopLsydQueryUpbillcodeResult `json:"result,omitempty" `
}
