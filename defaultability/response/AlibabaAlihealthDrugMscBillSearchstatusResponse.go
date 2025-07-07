package response

import (
    "topsdk/defaultability/domain"
)

type AlibabaAlihealthDrugMscBillSearchstatusResponse struct {

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
    Result  domain.AlibabaAlihealthDrugMscBillSearchstatusResultModel `json:"result,omitempty" `
}
