package domain


type AlibabaAlihealthDrugtraceTopYljgQueryGetentinfoResultModel struct {
    /*
        返回对象     */
    Model  *AlibabaAlihealthDrugtraceTopYljgQueryGetentinfoPUserEntInfoDto `json:"model,omitempty" `

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

func (s *AlibabaAlihealthDrugtraceTopYljgQueryGetentinfoResultModel) SetModel(v AlibabaAlihealthDrugtraceTopYljgQueryGetentinfoPUserEntInfoDto) *AlibabaAlihealthDrugtraceTopYljgQueryGetentinfoResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryGetentinfoResultModel) SetMsgCode(v string) *AlibabaAlihealthDrugtraceTopYljgQueryGetentinfoResultModel {
    s.MsgCode = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryGetentinfoResultModel) SetMsgInfo(v string) *AlibabaAlihealthDrugtraceTopYljgQueryGetentinfoResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryGetentinfoResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthDrugtraceTopYljgQueryGetentinfoResultModel {
    s.ResponseSuccess = &v
    return s
}
