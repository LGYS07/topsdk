package domain


type AlibabaAlihealthDrugtraceTopLsydQueryGetentinfoResultModel struct {
    /*
        返回对象     */
    Model  *AlibabaAlihealthDrugtraceTopLsydQueryGetentinfoPUserEntInfoDto `json:"model,omitempty" `

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

func (s *AlibabaAlihealthDrugtraceTopLsydQueryGetentinfoResultModel) SetModel(v AlibabaAlihealthDrugtraceTopLsydQueryGetentinfoPUserEntInfoDto) *AlibabaAlihealthDrugtraceTopLsydQueryGetentinfoResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryGetentinfoResultModel) SetMsgCode(v string) *AlibabaAlihealthDrugtraceTopLsydQueryGetentinfoResultModel {
    s.MsgCode = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryGetentinfoResultModel) SetMsgInfo(v string) *AlibabaAlihealthDrugtraceTopLsydQueryGetentinfoResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryGetentinfoResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthDrugtraceTopLsydQueryGetentinfoResultModel {
    s.ResponseSuccess = &v
    return s
}
