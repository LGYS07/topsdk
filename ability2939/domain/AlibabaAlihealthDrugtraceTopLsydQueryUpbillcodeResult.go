package domain


type AlibabaAlihealthDrugtraceTopLsydQueryUpbillcodeResult struct {
    /*
        是否成功     */
    Success  *bool `json:"success,omitempty" `

    /*
        model     */
    ModelList  *[]AlibabaAlihealthDrugtraceTopLsydQueryUpbillcodeBillUpstreamDTO `json:"model_list,omitempty" `

    /*
        msgInfo     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        msgCode     */
    MsgCode  *string `json:"msg_code,omitempty" `

}

func (s *AlibabaAlihealthDrugtraceTopLsydQueryUpbillcodeResult) SetSuccess(v bool) *AlibabaAlihealthDrugtraceTopLsydQueryUpbillcodeResult {
    s.Success = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryUpbillcodeResult) SetModelList(v []AlibabaAlihealthDrugtraceTopLsydQueryUpbillcodeBillUpstreamDTO) *AlibabaAlihealthDrugtraceTopLsydQueryUpbillcodeResult {
    s.ModelList = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryUpbillcodeResult) SetMsgInfo(v string) *AlibabaAlihealthDrugtraceTopLsydQueryUpbillcodeResult {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryUpbillcodeResult) SetMsgCode(v string) *AlibabaAlihealthDrugtraceTopLsydQueryUpbillcodeResult {
    s.MsgCode = &v
    return s
}
