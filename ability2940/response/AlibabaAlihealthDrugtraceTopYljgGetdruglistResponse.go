package response

import (
	"github.com/LGYS07/topsdk/ability2940/domain"
)

type AlibabaAlihealthDrugtraceTopYljgGetdruglistResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   返回结果
	*/
	Result domain.AlibabaAlihealthDrugtraceTopYljgGetdruglistResultModel `json:"result,omitempty" `
}
