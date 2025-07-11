package response

import (
	"github.com/LGYS07/topsdk/ability2940/domain"
)

type AlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeResponse struct {

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
	Result domain.AlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeResult `json:"result,omitempty" `
}
