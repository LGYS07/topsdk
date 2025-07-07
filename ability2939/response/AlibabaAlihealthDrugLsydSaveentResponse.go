package response

import (
    "topsdk/ability2939/domain"
)

type AlibabaAlihealthDrugLsydSaveentResponse struct {

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
    Result  domain.AlibabaAlihealthDrugLsydSaveentResultModel `json:"result,omitempty" `
}
