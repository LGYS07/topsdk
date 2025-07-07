package domain


type AlibabaAlihealthDrugMscSearchbillResultModel struct {
    /*
        返回模型     */
    Model  *AlibabaAlihealthDrugMscSearchbillPageInfoDto `json:"model,omitempty" `

    /*
        返回码     */
    MsgCode  *string `json:"msg_code,omitempty" `

    /*
        返回信息     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        是否响应成功     */
    ResponseSuccess  *bool `json:"response_success,omitempty" `

}

func (s *AlibabaAlihealthDrugMscSearchbillResultModel) SetModel(v AlibabaAlihealthDrugMscSearchbillPageInfoDto) *AlibabaAlihealthDrugMscSearchbillResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthDrugMscSearchbillResultModel) SetMsgCode(v string) *AlibabaAlihealthDrugMscSearchbillResultModel {
    s.MsgCode = &v
    return s
}
func (s *AlibabaAlihealthDrugMscSearchbillResultModel) SetMsgInfo(v string) *AlibabaAlihealthDrugMscSearchbillResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthDrugMscSearchbillResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthDrugMscSearchbillResultModel {
    s.ResponseSuccess = &v
    return s
}
