package domain


type AlibabaAlihealthDrugtraceTopYljgListupoutResultModel struct {
    /*
        返回对象     */
    Model  *AlibabaAlihealthDrugtraceTopYljgListupoutPageInfoDto `json:"model,omitempty" `

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

func (s *AlibabaAlihealthDrugtraceTopYljgListupoutResultModel) SetModel(v AlibabaAlihealthDrugtraceTopYljgListupoutPageInfoDto) *AlibabaAlihealthDrugtraceTopYljgListupoutResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgListupoutResultModel) SetMsgCode(v string) *AlibabaAlihealthDrugtraceTopYljgListupoutResultModel {
    s.MsgCode = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgListupoutResultModel) SetMsgInfo(v string) *AlibabaAlihealthDrugtraceTopYljgListupoutResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgListupoutResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthDrugtraceTopYljgListupoutResultModel {
    s.ResponseSuccess = &v
    return s
}
