package domain


type AlibabaAlihealthDrugtraceTopYljgQueryCodedetailResultModel struct {
    /*
        内层大对象     */
    Models  *[]AlibabaAlihealthDrugtraceTopYljgQueryCodedetailCodeFullInfoDto `json:"models,omitempty" `

    /*
        消息码     */
    MsgCode  *string `json:"msg_code,omitempty" `

    /*
        消息提示内容     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        查询成功失败标记     */
    ResponseSuccess  *bool `json:"response_success,omitempty" `

}

func (s *AlibabaAlihealthDrugtraceTopYljgQueryCodedetailResultModel) SetModels(v []AlibabaAlihealthDrugtraceTopYljgQueryCodedetailCodeFullInfoDto) *AlibabaAlihealthDrugtraceTopYljgQueryCodedetailResultModel {
    s.Models = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryCodedetailResultModel) SetMsgCode(v string) *AlibabaAlihealthDrugtraceTopYljgQueryCodedetailResultModel {
    s.MsgCode = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryCodedetailResultModel) SetMsgInfo(v string) *AlibabaAlihealthDrugtraceTopYljgQueryCodedetailResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryCodedetailResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthDrugtraceTopYljgQueryCodedetailResultModel {
    s.ResponseSuccess = &v
    return s
}
