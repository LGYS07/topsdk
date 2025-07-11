package domain


type AlibabaAlihealthDrugtraceTopYljgDrugtableResultModel struct {
    /*
        返回模型     */
    Model  *AlibabaAlihealthDrugtraceTopYljgDrugtablePage `json:"model,omitempty" `

    /*
        状态码     */
    MsgCode  *string `json:"msg_code,omitempty" `

    /*
        状态值     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        是否响应成功     */
    ResponseSuccess  *bool `json:"response_success,omitempty" `

}

func (s *AlibabaAlihealthDrugtraceTopYljgDrugtableResultModel) SetModel(v AlibabaAlihealthDrugtraceTopYljgDrugtablePage) *AlibabaAlihealthDrugtraceTopYljgDrugtableResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgDrugtableResultModel) SetMsgCode(v string) *AlibabaAlihealthDrugtraceTopYljgDrugtableResultModel {
    s.MsgCode = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgDrugtableResultModel) SetMsgInfo(v string) *AlibabaAlihealthDrugtraceTopYljgDrugtableResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgDrugtableResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthDrugtraceTopYljgDrugtableResultModel {
    s.ResponseSuccess = &v
    return s
}
