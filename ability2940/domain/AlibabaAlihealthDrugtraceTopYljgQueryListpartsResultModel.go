package domain


type AlibabaAlihealthDrugtraceTopYljgQueryListpartsResultModel struct {
    /*
        返回对象     */
    Model  *AlibabaAlihealthDrugtraceTopYljgQueryListpartsPage `json:"model,omitempty" `

    /*
        返回码     */
    MsgCode  *string `json:"msg_code,omitempty" `

    /*
        返回值     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        是否响应成功     */
    ResponseSuccess  *bool `json:"response_success,omitempty" `

}

func (s *AlibabaAlihealthDrugtraceTopYljgQueryListpartsResultModel) SetModel(v AlibabaAlihealthDrugtraceTopYljgQueryListpartsPage) *AlibabaAlihealthDrugtraceTopYljgQueryListpartsResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryListpartsResultModel) SetMsgCode(v string) *AlibabaAlihealthDrugtraceTopYljgQueryListpartsResultModel {
    s.MsgCode = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryListpartsResultModel) SetMsgInfo(v string) *AlibabaAlihealthDrugtraceTopYljgQueryListpartsResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryListpartsResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthDrugtraceTopYljgQueryListpartsResultModel {
    s.ResponseSuccess = &v
    return s
}
