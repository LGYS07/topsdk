package domain


type AlibabaAlihealthDrugtraceTopYljgGetdruglistResultModel struct {
    /*
        返回模型     */
    Model  *AlibabaAlihealthDrugtraceTopYljgGetdruglistPage `json:"model,omitempty" `

    /*
        状态码     */
    MsgCode  *string `json:"msg_code,omitempty" `

    /*
        状态值     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        是否响应成功     */
    ResponseSuccess  *string `json:"response_success,omitempty" `

}

func (s *AlibabaAlihealthDrugtraceTopYljgGetdruglistResultModel) SetModel(v AlibabaAlihealthDrugtraceTopYljgGetdruglistPage) *AlibabaAlihealthDrugtraceTopYljgGetdruglistResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgGetdruglistResultModel) SetMsgCode(v string) *AlibabaAlihealthDrugtraceTopYljgGetdruglistResultModel {
    s.MsgCode = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgGetdruglistResultModel) SetMsgInfo(v string) *AlibabaAlihealthDrugtraceTopYljgGetdruglistResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgGetdruglistResultModel) SetResponseSuccess(v string) *AlibabaAlihealthDrugtraceTopYljgGetdruglistResultModel {
    s.ResponseSuccess = &v
    return s
}
