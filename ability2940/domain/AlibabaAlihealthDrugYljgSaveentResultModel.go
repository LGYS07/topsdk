package domain


type AlibabaAlihealthDrugYljgSaveentResultModel struct {
    /*
        true：接口调用成功     */
    Success  *bool `json:"success,omitempty" `

    /*
        具体返回值     */
    Model  *AlibabaAlihealthDrugYljgSaveentModel `json:"model,omitempty" `

    /*
        接口调用失败具体信息     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        接口调用失败具体code     */
    MsgCode  *string `json:"msg_code,omitempty" `

}

func (s *AlibabaAlihealthDrugYljgSaveentResultModel) SetSuccess(v bool) *AlibabaAlihealthDrugYljgSaveentResultModel {
    s.Success = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgSaveentResultModel) SetModel(v AlibabaAlihealthDrugYljgSaveentModel) *AlibabaAlihealthDrugYljgSaveentResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgSaveentResultModel) SetMsgInfo(v string) *AlibabaAlihealthDrugYljgSaveentResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgSaveentResultModel) SetMsgCode(v string) *AlibabaAlihealthDrugYljgSaveentResultModel {
    s.MsgCode = &v
    return s
}
