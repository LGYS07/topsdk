package response

import (
)

type AlibabaAlihealthDrugtraceTopYljgUploadinoutbillResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        返回值
    */
    Model  string `json:"model,omitempty" `
    /*
        返回编码(BILL_DECODE_ERROR 单据转码失败  BILL_FILE_NAME_DUPLICATE_UPLOAD 文件名重复)
    */
    MsgCode  string `json:"msg_code,omitempty" `
    /*
        返回信息
    */
    MsgInfo  string `json:"msg_info,omitempty" `
    /*
        是否成功(true 成功 false 失败)
    */
    ResponseSuccess  bool `json:"response_success,omitempty" `
}
