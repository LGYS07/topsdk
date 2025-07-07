package response

import (
	"github.com/LGYS07/topsdk/defaultability/domain"
)

type AlibabaAlihealthDrugMscGetbyentidResponse struct {

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
	Result domain.AlibabaAlihealthDrugMscGetbyentidResultModel `json:"result,omitempty" `
}
