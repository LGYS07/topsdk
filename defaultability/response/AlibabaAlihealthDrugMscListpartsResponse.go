package response

import (
    "topsdk/defaultability/domain"
)

type AlibabaAlihealthDrugMscListpartsResponse struct {

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
    Result  domain.AlibabaAlihealthDrugMscListpartsResultModel `json:"result,omitempty" `
}
