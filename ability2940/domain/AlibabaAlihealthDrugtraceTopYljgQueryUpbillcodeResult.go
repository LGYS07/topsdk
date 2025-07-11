package domain


type AlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeResult struct {
    /*
        是否成功     */
    Success  *bool `json:"success,omitempty" `

    /*
        model     */
    ModelList  *[]AlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeBillUpstreamDTO `json:"model_list,omitempty" `

    /*
        msgInfo     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        msgCode     */
    MsgCode  *string `json:"msg_code,omitempty" `

}

func (s *AlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeResult) SetSuccess(v bool) *AlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeResult {
    s.Success = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeResult) SetModelList(v []AlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeBillUpstreamDTO) *AlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeResult {
    s.ModelList = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeResult) SetMsgInfo(v string) *AlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeResult {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeResult) SetMsgCode(v string) *AlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeResult {
    s.MsgCode = &v
    return s
}
