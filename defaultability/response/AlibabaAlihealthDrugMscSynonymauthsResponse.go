package response

import (
	"github.com/LGYS07/topsdk/defaultability/domain"
)

type AlibabaAlihealthDrugMscSynonymauthsResponse struct {

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
	Result domain.AlibabaAlihealthDrugMscSynonymauthsResultModel `json:"result,omitempty" `
}
