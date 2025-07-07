package response

import (
	"github.com/LGYS07/topsdk/defaultability/domain"
)

type AlibabaAlihealthDrugMscGetentinfoResponse struct {

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
	Result domain.AlibabaAlihealthDrugMscGetentinfoResultModel `json:"result,omitempty" `
}
