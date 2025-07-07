package domain


type AlibabaAlihealthDrugtraceTopLsydListupoutResultModel struct {
    /*
        返回对象     */
    Model  *AlibabaAlihealthDrugtraceTopLsydListupoutPageInfoDto `json:"model,omitempty" `

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

func (s *AlibabaAlihealthDrugtraceTopLsydListupoutResultModel) SetModel(v AlibabaAlihealthDrugtraceTopLsydListupoutPageInfoDto) *AlibabaAlihealthDrugtraceTopLsydListupoutResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydListupoutResultModel) SetMsgCode(v string) *AlibabaAlihealthDrugtraceTopLsydListupoutResultModel {
    s.MsgCode = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydListupoutResultModel) SetMsgInfo(v string) *AlibabaAlihealthDrugtraceTopLsydListupoutResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydListupoutResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthDrugtraceTopLsydListupoutResultModel {
    s.ResponseSuccess = &v
    return s
}
