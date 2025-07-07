package domain


type AlibabaAlihealthDrugMscBillSearchstatusResultModel struct {
    /*
        返回对象     */
    Model  *AlibabaAlihealthDrugMscBillSearchstatusPageInfoDto `json:"model,omitempty" `

    /*
        状态码     */
    MsgCode  *string `json:"msg_code,omitempty" `

    /*
        状态值     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        响应结果     */
    ResponseSuccess  *bool `json:"response_success,omitempty" `

}

func (s *AlibabaAlihealthDrugMscBillSearchstatusResultModel) SetModel(v AlibabaAlihealthDrugMscBillSearchstatusPageInfoDto) *AlibabaAlihealthDrugMscBillSearchstatusResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthDrugMscBillSearchstatusResultModel) SetMsgCode(v string) *AlibabaAlihealthDrugMscBillSearchstatusResultModel {
    s.MsgCode = &v
    return s
}
func (s *AlibabaAlihealthDrugMscBillSearchstatusResultModel) SetMsgInfo(v string) *AlibabaAlihealthDrugMscBillSearchstatusResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthDrugMscBillSearchstatusResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthDrugMscBillSearchstatusResultModel {
    s.ResponseSuccess = &v
    return s
}
