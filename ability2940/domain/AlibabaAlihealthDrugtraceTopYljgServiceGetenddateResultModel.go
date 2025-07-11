package domain


type AlibabaAlihealthDrugtraceTopYljgServiceGetenddateResultModel struct {
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

func (s *AlibabaAlihealthDrugtraceTopYljgServiceGetenddateResultModel) SetSuccess(v bool) *AlibabaAlihealthDrugtraceTopYljgServiceGetenddateResultModel {
    s.Success = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgServiceGetenddateResultModel) SetEndDate(v string) *AlibabaAlihealthDrugtraceTopYljgServiceGetenddateResultModel {
    s.EndDate = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgServiceGetenddateResultModel) SetMsgInfo(v string) *AlibabaAlihealthDrugtraceTopYljgServiceGetenddateResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgServiceGetenddateResultModel) SetMsgCode(v string) *AlibabaAlihealthDrugtraceTopYljgServiceGetenddateResultModel {
    s.MsgCode = &v
    return s
}
