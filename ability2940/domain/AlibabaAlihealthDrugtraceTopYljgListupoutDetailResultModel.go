package domain


type AlibabaAlihealthDrugtraceTopYljgListupoutDetailResultModel struct {
    /*
        最外层对象     */
    Model  *AlibabaAlihealthDrugtraceTopYljgListupoutDetailBillUpOutDetailDto `json:"model,omitempty" `

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

func (s *AlibabaAlihealthDrugtraceTopYljgListupoutDetailResultModel) SetModel(v AlibabaAlihealthDrugtraceTopYljgListupoutDetailBillUpOutDetailDto) *AlibabaAlihealthDrugtraceTopYljgListupoutDetailResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgListupoutDetailResultModel) SetMsgCode(v string) *AlibabaAlihealthDrugtraceTopYljgListupoutDetailResultModel {
    s.MsgCode = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgListupoutDetailResultModel) SetMsgInfo(v string) *AlibabaAlihealthDrugtraceTopYljgListupoutDetailResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgListupoutDetailResultModel) SetSuccess(v bool) *AlibabaAlihealthDrugtraceTopYljgListupoutDetailResultModel {
    s.Success = &v
    return s
}
