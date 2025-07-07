package domain


type AlibabaAlihealthDrugtraceTopLsydListupoutDetailResultModel struct {
    /*
        最外层对象     */
    Model  *AlibabaAlihealthDrugtraceTopLsydListupoutDetailBillUpOutDetailDto `json:"model,omitempty" `

    /*
        提示信息编码     */
    MsgCode  *string `json:"msg_code,omitempty" `

    /*
        提示信息内容     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        成功失败标记     */
    Success  *bool `json:"success,omitempty" `

}

func (s *AlibabaAlihealthDrugtraceTopLsydListupoutDetailResultModel) SetModel(v AlibabaAlihealthDrugtraceTopLsydListupoutDetailBillUpOutDetailDto) *AlibabaAlihealthDrugtraceTopLsydListupoutDetailResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydListupoutDetailResultModel) SetMsgCode(v string) *AlibabaAlihealthDrugtraceTopLsydListupoutDetailResultModel {
    s.MsgCode = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydListupoutDetailResultModel) SetMsgInfo(v string) *AlibabaAlihealthDrugtraceTopLsydListupoutDetailResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydListupoutDetailResultModel) SetSuccess(v bool) *AlibabaAlihealthDrugtraceTopLsydListupoutDetailResultModel {
    s.Success = &v
    return s
}
