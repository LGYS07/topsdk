package domain


type AlibabaAlihealthDrugMscBilloutDetailwithcodesResultModel struct {
    /*
        对象模型信息     */
    Model  *AlibabaAlihealthDrugMscBilloutDetailwithcodesBillInOutDetailDto `json:"model,omitempty" `

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

func (s *AlibabaAlihealthDrugMscBilloutDetailwithcodesResultModel) SetModel(v AlibabaAlihealthDrugMscBilloutDetailwithcodesBillInOutDetailDto) *AlibabaAlihealthDrugMscBilloutDetailwithcodesResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthDrugMscBilloutDetailwithcodesResultModel) SetMsgCode(v string) *AlibabaAlihealthDrugMscBilloutDetailwithcodesResultModel {
    s.MsgCode = &v
    return s
}
func (s *AlibabaAlihealthDrugMscBilloutDetailwithcodesResultModel) SetMsgInfo(v string) *AlibabaAlihealthDrugMscBilloutDetailwithcodesResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthDrugMscBilloutDetailwithcodesResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthDrugMscBilloutDetailwithcodesResultModel {
    s.ResponseSuccess = &v
    return s
}
