package domain


type AlibabaAlihealthDrugtraceTopLsydQueryCodedetailResultModel struct {
    /*
        内层大对象     */
    Models  *[]AlibabaAlihealthDrugtraceTopLsydQueryCodedetailCodeFullInfoDto `json:"models,omitempty" `

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

func (s *AlibabaAlihealthDrugtraceTopLsydQueryCodedetailResultModel) SetModels(v []AlibabaAlihealthDrugtraceTopLsydQueryCodedetailCodeFullInfoDto) *AlibabaAlihealthDrugtraceTopLsydQueryCodedetailResultModel {
    s.Models = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryCodedetailResultModel) SetMsgCode(v string) *AlibabaAlihealthDrugtraceTopLsydQueryCodedetailResultModel {
    s.MsgCode = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryCodedetailResultModel) SetMsgInfo(v string) *AlibabaAlihealthDrugtraceTopLsydQueryCodedetailResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryCodedetailResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthDrugtraceTopLsydQueryCodedetailResultModel {
    s.ResponseSuccess = &v
    return s
}
