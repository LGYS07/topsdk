package domain


type AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailResultModel struct {
    /*
        对象模型信息     */
    Model  *AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailBillInOutDetailDto `json:"model,omitempty" `

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

func (s *AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailResultModel) SetModel(v AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailBillInOutDetailDto) *AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailResultModel) SetMsgCode(v string) *AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailResultModel {
    s.MsgCode = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailResultModel) SetMsgInfo(v string) *AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailResultModel {
    s.ResponseSuccess = &v
    return s
}
