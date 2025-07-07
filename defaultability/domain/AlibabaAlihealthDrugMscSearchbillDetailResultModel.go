package domain


type AlibabaAlihealthDrugMscSearchbillDetailResultModel struct {
    /*
        对象模型信息     */
    Model  *AlibabaAlihealthDrugMscSearchbillDetailBillInOutDetailDto `json:"model,omitempty" `

    /*
        消息码     */
    MsgCode  *string `json:"msg_code,omitempty" `

    /*
        消息     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        成功失败     */
    ResponseSuccess  *bool `json:"response_success,omitempty" `

}

func (s *AlibabaAlihealthDrugMscSearchbillDetailResultModel) SetModel(v AlibabaAlihealthDrugMscSearchbillDetailBillInOutDetailDto) *AlibabaAlihealthDrugMscSearchbillDetailResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthDrugMscSearchbillDetailResultModel) SetMsgCode(v string) *AlibabaAlihealthDrugMscSearchbillDetailResultModel {
    s.MsgCode = &v
    return s
}
func (s *AlibabaAlihealthDrugMscSearchbillDetailResultModel) SetMsgInfo(v string) *AlibabaAlihealthDrugMscSearchbillDetailResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthDrugMscSearchbillDetailResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthDrugMscSearchbillDetailResultModel {
    s.ResponseSuccess = &v
    return s
}
