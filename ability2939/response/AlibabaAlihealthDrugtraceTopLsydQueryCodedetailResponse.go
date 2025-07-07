package response

import (
	"github.com/LGYS07/topsdk/ability2939/domain"
)

type AlibabaAlihealthDrugtraceTopLsydQueryCodedetailResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   最外层结果
	*/
	Result domain.AlibabaAlihealthDrugtraceTopLsydQueryCodedetailResultModel `json:"result,omitempty" `
}
