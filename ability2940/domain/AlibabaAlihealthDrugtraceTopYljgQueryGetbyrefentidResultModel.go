package domain


type AlibabaAlihealthDrugtraceTopYljgQueryGetbyrefentidResultModel struct {
    /*
        返回对象     */
    Model  *AlibabaAlihealthDrugtraceTopYljgQueryGetbyrefentidPUserEntInfoDto `json:"model,omitempty" `

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

func (s *AlibabaAlihealthDrugtraceTopYljgQueryGetbyrefentidResultModel) SetModel(v AlibabaAlihealthDrugtraceTopYljgQueryGetbyrefentidPUserEntInfoDto) *AlibabaAlihealthDrugtraceTopYljgQueryGetbyrefentidResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryGetbyrefentidResultModel) SetMsgCode(v string) *AlibabaAlihealthDrugtraceTopYljgQueryGetbyrefentidResultModel {
    s.MsgCode = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryGetbyrefentidResultModel) SetMsgInfo(v string) *AlibabaAlihealthDrugtraceTopYljgQueryGetbyrefentidResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryGetbyrefentidResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthDrugtraceTopYljgQueryGetbyrefentidResultModel {
    s.ResponseSuccess = &v
    return s
}
