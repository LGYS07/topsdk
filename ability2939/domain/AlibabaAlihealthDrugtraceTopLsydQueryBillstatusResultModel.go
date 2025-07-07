package domain


type AlibabaAlihealthDrugtraceTopLsydQueryBillstatusResultModel struct {
    /*
        返回对象     */
    Model  *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusPageInfoDto `json:"model,omitempty" `

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

func (s *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusResultModel) SetModel(v AlibabaAlihealthDrugtraceTopLsydQueryBillstatusPageInfoDto) *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusResultModel) SetMsgCode(v string) *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusResultModel {
    s.MsgCode = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusResultModel) SetMsgInfo(v string) *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusResultModel {
    s.ResponseSuccess = &v
    return s
}
