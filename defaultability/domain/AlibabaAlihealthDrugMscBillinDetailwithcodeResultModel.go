package domain


type AlibabaAlihealthDrugMscBillinDetailwithcodeResultModel struct {
    /*
        对象模型信息     */
    Model  *AlibabaAlihealthDrugMscBillinDetailwithcodeBillInOutDetailDto `json:"model,omitempty" `

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

func (s *AlibabaAlihealthDrugMscBillinDetailwithcodeResultModel) SetModel(v AlibabaAlihealthDrugMscBillinDetailwithcodeBillInOutDetailDto) *AlibabaAlihealthDrugMscBillinDetailwithcodeResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthDrugMscBillinDetailwithcodeResultModel) SetMsgCode(v string) *AlibabaAlihealthDrugMscBillinDetailwithcodeResultModel {
    s.MsgCode = &v
    return s
}
func (s *AlibabaAlihealthDrugMscBillinDetailwithcodeResultModel) SetMsgInfo(v string) *AlibabaAlihealthDrugMscBillinDetailwithcodeResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthDrugMscBillinDetailwithcodeResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthDrugMscBillinDetailwithcodeResultModel {
    s.ResponseSuccess = &v
    return s
}
