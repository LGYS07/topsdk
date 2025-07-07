package response

import (
)

type AlibabaAlihealthDrugtraceTopLsydUploadretailResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        上传单据文件队列表标识
    */
    Model  string `json:"model,omitempty" `
    /*
        错误码(BILL_DECODE_ERROR 单据转码失败 2.BILL_FILE_NAME_DUPLICATE_UPLOAD 文件名重复)
    */
    MsgCode  string `json:"msg_code,omitempty" `
    /*
        错误信息
    */
    MsgInfo  string `json:"msg_info,omitempty" `
    /*
        操作是否成功(true 成功 ,false失败)
    */
    ResponseSuccess  bool `json:"response_success,omitempty" `
}
