package domain


type AlibabaAlihealthDrugLsydSaveentResultModel struct {
    /*
        true：接口调用成功     */
    Success  *bool `json:"success,omitempty" `

    /*
        具体返回值     */
    Model  *AlibabaAlihealthDrugLsydSaveentModel `json:"model,omitempty" `

    /*
        接口调用失败具体信息     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        接口调用失败具体code     */
    MsgCode  *string `json:"msg_code,omitempty" `

}

func (s *AlibabaAlihealthDrugLsydSaveentResultModel) SetSuccess(v bool) *AlibabaAlihealthDrugLsydSaveentResultModel {
    s.Success = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydSaveentResultModel) SetModel(v AlibabaAlihealthDrugLsydSaveentModel) *AlibabaAlihealthDrugLsydSaveentResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydSaveentResultModel) SetMsgInfo(v string) *AlibabaAlihealthDrugLsydSaveentResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydSaveentResultModel) SetMsgCode(v string) *AlibabaAlihealthDrugLsydSaveentResultModel {
    s.MsgCode = &v
    return s
}
