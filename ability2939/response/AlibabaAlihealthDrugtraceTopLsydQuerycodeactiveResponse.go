package response

import (
)

type AlibabaAlihealthDrugtraceTopLsydQuerycodeactiveResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        未激活或者不存在的码
    */
    Models  []string `json:"models,omitempty" `
    /*
        消息码
    */
    MsgCode  string `json:"msg_code,omitempty" `
    /*
        消息提示内容
    */
    MsgInfo  string `json:"msg_info,omitempty" `
    /*
        是否成功(true 成功 ,false失败)
    */
    ResponseSuccess  bool `json:"response_success,omitempty" `
}
