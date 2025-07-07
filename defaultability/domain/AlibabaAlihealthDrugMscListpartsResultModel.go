package domain


type AlibabaAlihealthDrugMscListpartsResultModel struct {
    /*
        是否响应成功     */
    ResponseSuccess  *bool `json:"response_success,omitempty" `

    /*
        返回对象     */
    Model  *AlibabaAlihealthDrugMscListpartsPage `json:"model,omitempty" `

    /*
        返回值     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        返回码     */
    MsgCode  *string `json:"msg_code,omitempty" `

}

func (s *AlibabaAlihealthDrugMscListpartsResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthDrugMscListpartsResultModel {
    s.ResponseSuccess = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListpartsResultModel) SetModel(v AlibabaAlihealthDrugMscListpartsPage) *AlibabaAlihealthDrugMscListpartsResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListpartsResultModel) SetMsgInfo(v string) *AlibabaAlihealthDrugMscListpartsResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListpartsResultModel) SetMsgCode(v string) *AlibabaAlihealthDrugMscListpartsResultModel {
    s.MsgCode = &v
    return s
}
