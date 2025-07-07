package domain


type AlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidResultModel struct {
    /*
        返回对象     */
    Model  *AlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidPUserEntInfoDto `json:"model,omitempty" `

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

func (s *AlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidResultModel) SetModel(v AlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidPUserEntInfoDto) *AlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidResultModel) SetMsgCode(v string) *AlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidResultModel {
    s.MsgCode = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidResultModel) SetMsgInfo(v string) *AlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidResultModel {
    s.ResponseSuccess = &v
    return s
}
