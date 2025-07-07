package domain


type AlibabaAlihealthDrugMscSaveentResultModel struct {
    /*
        true：接口调用成功     */
    Success  *bool `json:"success,omitempty" `

    /*
        具体返回值     */
    Model  *AlibabaAlihealthDrugMscSaveentModel `json:"model,omitempty" `

    /*
        接口调用失败具体信息     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        接口调用失败具体code     */
    MsgCode  *string `json:"msg_code,omitempty" `

}

func (s *AlibabaAlihealthDrugMscSaveentResultModel) SetSuccess(v bool) *AlibabaAlihealthDrugMscSaveentResultModel {
    s.Success = &v
    return s
}
func (s *AlibabaAlihealthDrugMscSaveentResultModel) SetModel(v AlibabaAlihealthDrugMscSaveentModel) *AlibabaAlihealthDrugMscSaveentResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthDrugMscSaveentResultModel) SetMsgInfo(v string) *AlibabaAlihealthDrugMscSaveentResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthDrugMscSaveentResultModel) SetMsgCode(v string) *AlibabaAlihealthDrugMscSaveentResultModel {
    s.MsgCode = &v
    return s
}
