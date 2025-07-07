package domain


type AlibabaAlihealthDrugtraceTopLsydServiceGetenddateResultModel struct {
    /*
        true：接口调用成功     */
    Success  *bool `json:"success,omitempty" `

    /*
        服务截止时间     */
    EndDate  *string `json:"end_date,omitempty" `

    /*
        接口调用失败具体信息     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        接口调用失败具体code     */
    MsgCode  *string `json:"msg_code,omitempty" `

}

func (s *AlibabaAlihealthDrugtraceTopLsydServiceGetenddateResultModel) SetSuccess(v bool) *AlibabaAlihealthDrugtraceTopLsydServiceGetenddateResultModel {
    s.Success = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydServiceGetenddateResultModel) SetEndDate(v string) *AlibabaAlihealthDrugtraceTopLsydServiceGetenddateResultModel {
    s.EndDate = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydServiceGetenddateResultModel) SetMsgInfo(v string) *AlibabaAlihealthDrugtraceTopLsydServiceGetenddateResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydServiceGetenddateResultModel) SetMsgCode(v string) *AlibabaAlihealthDrugtraceTopLsydServiceGetenddateResultModel {
    s.MsgCode = &v
    return s
}
