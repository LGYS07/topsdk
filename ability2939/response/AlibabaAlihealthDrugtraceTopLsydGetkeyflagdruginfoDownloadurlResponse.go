package response

import (
    "topsdk/ability2939/domain"
)

type AlibabaAlihealthDrugtraceTopLsydGetkeyflagdruginfoDownloadurlResponse struct {

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
    Result  domain.AlibabaAlihealthDrugtraceTopLsydGetkeyflagdruginfoDownloadurlResultModel `json:"result,omitempty" `
}
