package response

import (
)

type AlibabaAlihealthDrugMscRemnantbillUploadResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        是否成功 
    */
    ResponseStatus  bool `json:"response_status,omitempty" `
    /*
        model
    */
    Model  string `json:"model,omitempty" `
    /*
        msgInfo
    */
    MsgInfo  string `json:"msg_info,omitempty" `
    /*
        msgCode
    */
    MsgCode  string `json:"msg_code,omitempty" `
}
