package response

import (
    "topsdk/ability2939/domain"
)

type AlibabaAlihealthDrugtraceTopLsydQueryRelationResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        接口返回model
    */
    Result  domain.AlibabaAlihealthDrugtraceTopLsydQueryRelationResultModel `json:"result,omitempty" `
}
