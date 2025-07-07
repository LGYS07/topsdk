package domain


type AlibabaAlihealthDrugtraceTopLsydQueryListpartsResultModel struct {
    /*
        返回对象     */
    Model  *AlibabaAlihealthDrugtraceTopLsydQueryListpartsPage `json:"model,omitempty" `

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

func (s *AlibabaAlihealthDrugtraceTopLsydQueryListpartsResultModel) SetModel(v AlibabaAlihealthDrugtraceTopLsydQueryListpartsPage) *AlibabaAlihealthDrugtraceTopLsydQueryListpartsResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryListpartsResultModel) SetMsgCode(v string) *AlibabaAlihealthDrugtraceTopLsydQueryListpartsResultModel {
    s.MsgCode = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryListpartsResultModel) SetMsgInfo(v string) *AlibabaAlihealthDrugtraceTopLsydQueryListpartsResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryListpartsResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthDrugtraceTopLsydQueryListpartsResultModel {
    s.ResponseSuccess = &v
    return s
}
