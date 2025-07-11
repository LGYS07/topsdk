package domain


type AlibabaAlihealthDrugtraceTopYljgQueryBillstatusResultModel struct {
    /*
        返回对象     */
    Model  *AlibabaAlihealthDrugtraceTopYljgQueryBillstatusPageInfoDto `json:"model,omitempty" `

    /*
        状态码     */
    MsgCode  *string `json:"msg_code,omitempty" `

    /*
        状态值     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        响应结果     */
    ResponseSuccess  *bool `json:"response_success,omitempty" `

}

func (s *AlibabaAlihealthDrugtraceTopYljgQueryBillstatusResultModel) SetModel(v AlibabaAlihealthDrugtraceTopYljgQueryBillstatusPageInfoDto) *AlibabaAlihealthDrugtraceTopYljgQueryBillstatusResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryBillstatusResultModel) SetMsgCode(v string) *AlibabaAlihealthDrugtraceTopYljgQueryBillstatusResultModel {
    s.MsgCode = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryBillstatusResultModel) SetMsgInfo(v string) *AlibabaAlihealthDrugtraceTopYljgQueryBillstatusResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryBillstatusResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthDrugtraceTopYljgQueryBillstatusResultModel {
    s.ResponseSuccess = &v
    return s
}
