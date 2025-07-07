package response

import (
    "topsdk/defaultability/domain"
)

type AlibabaAlihealthDrugMscGetentinfonewResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        resultModel
    */
    Result  domain.AlibabaAlihealthDrugMscGetentinfonewResultModel `json:"result,omitempty" `
}
