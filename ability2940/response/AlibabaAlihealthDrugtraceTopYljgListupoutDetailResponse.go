package response

import (
	"github.com/LGYS07/topsdk/ability2940/domain"
)

type AlibabaAlihealthDrugtraceTopYljgListupoutDetailResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   监控宝推送网站监控信息，返回结果
	*/
	Result domain.AlibabaAlihealthDrugtraceTopYljgListupoutDetailResultModel `json:"result,omitempty" `
}
