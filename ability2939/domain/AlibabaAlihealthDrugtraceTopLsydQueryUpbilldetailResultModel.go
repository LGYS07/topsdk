package domain


type AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailResultModel struct {
    /*
        对象模型信息     */
    Model  *AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailBillInOutDetailDto `json:"model,omitempty" `

    /*
        消息码     */
    MsgCode  *string `json:"msg_code,omitempty" `

    /*
        消息     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        成功失败     */
    ResponseSuccess  *bool `json:"response_success,omitempty" `

}

func (s *AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailResultModel) SetModel(v AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailBillInOutDetailDto) *AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailResultModel) SetMsgCode(v string) *AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailResultModel {
    s.MsgCode = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailResultModel) SetMsgInfo(v string) *AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailResultModel {
    s.ResponseSuccess = &v
    return s
}
