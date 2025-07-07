package response

import (
    "topsdk/ability2939/domain"
)

type AlibabaAlihealthDrugtraceTopLsydListupoutResponse struct {

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
    Result  domain.AlibabaAlihealthDrugtraceTopLsydListupoutResultModel `json:"result,omitempty" `
}
