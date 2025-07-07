package domain


type AlibabaAlihealthDrugMscSynonymauthsResultModel struct {
    /*
        是否响应成功     */
    ResponseSuccess  *bool `json:"response_success,omitempty" `

    /*
        返回对象     */
    Model  *AlibabaAlihealthDrugMscSynonymauthsPage `json:"model,omitempty" `

    /*
        状态值     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        状态码     */
    MsgCode  *string `json:"msg_code,omitempty" `

}

func (s *AlibabaAlihealthDrugMscSynonymauthsResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthDrugMscSynonymauthsResultModel {
    s.ResponseSuccess = &v
    return s
}
func (s *AlibabaAlihealthDrugMscSynonymauthsResultModel) SetModel(v AlibabaAlihealthDrugMscSynonymauthsPage) *AlibabaAlihealthDrugMscSynonymauthsResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthDrugMscSynonymauthsResultModel) SetMsgInfo(v string) *AlibabaAlihealthDrugMscSynonymauthsResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthDrugMscSynonymauthsResultModel) SetMsgCode(v string) *AlibabaAlihealthDrugMscSynonymauthsResultModel {
    s.MsgCode = &v
    return s
}
