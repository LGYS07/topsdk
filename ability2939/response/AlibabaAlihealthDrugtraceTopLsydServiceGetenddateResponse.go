package response

import (
    "topsdk/ability2939/domain"
)

type AlibabaAlihealthDrugtraceTopLsydServiceGetenddateResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        接口返回
    */
    Result  domain.AlibabaAlihealthDrugtraceTopLsydServiceGetenddateResultModel `json:"result,omitempty" `
}
