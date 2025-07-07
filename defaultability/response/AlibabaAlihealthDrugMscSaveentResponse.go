package response

import (
	"github.com/LGYS07/topsdk/defaultability/domain"
)

type AlibabaAlihealthDrugMscSaveentResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   往来单位新增接口返回
	*/
	Result domain.AlibabaAlihealthDrugMscSaveentResultModel `json:"result,omitempty" `
}
