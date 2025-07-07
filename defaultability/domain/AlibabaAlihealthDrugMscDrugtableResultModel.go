package domain


type AlibabaAlihealthDrugMscDrugtableResultModel struct {
    /*
        是否响应成功     */
    ResponseSuccess  *bool `json:"response_success,omitempty" `

    /*
        返回对象     */
    Model  *AlibabaAlihealthDrugMscDrugtablePage `json:"model,omitempty" `

    /*
        状态值     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        状态码     */
    MsgCode  *string `json:"msg_code,omitempty" `

}

func (s *AlibabaAlihealthDrugMscDrugtableResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthDrugMscDrugtableResultModel {
    s.ResponseSuccess = &v
    return s
}
func (s *AlibabaAlihealthDrugMscDrugtableResultModel) SetModel(v AlibabaAlihealthDrugMscDrugtablePage) *AlibabaAlihealthDrugMscDrugtableResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthDrugMscDrugtableResultModel) SetMsgInfo(v string) *AlibabaAlihealthDrugMscDrugtableResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthDrugMscDrugtableResultModel) SetMsgCode(v string) *AlibabaAlihealthDrugMscDrugtableResultModel {
    s.MsgCode = &v
    return s
}
