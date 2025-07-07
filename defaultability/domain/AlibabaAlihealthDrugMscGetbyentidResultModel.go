package domain


type AlibabaAlihealthDrugMscGetbyentidResultModel struct {
    /*
        是否响应成功     */
    ResponseSuccess  *bool `json:"response_success,omitempty" `

    /*
        返回对象     */
    Model  *AlibabaAlihealthDrugMscGetbyentidPUserEntInfoDTO `json:"model,omitempty" `

    /*
        返回值     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        返回码     */
    MsgCode  *string `json:"msg_code,omitempty" `

}

func (s *AlibabaAlihealthDrugMscGetbyentidResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthDrugMscGetbyentidResultModel {
    s.ResponseSuccess = &v
    return s
}
func (s *AlibabaAlihealthDrugMscGetbyentidResultModel) SetModel(v AlibabaAlihealthDrugMscGetbyentidPUserEntInfoDTO) *AlibabaAlihealthDrugMscGetbyentidResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthDrugMscGetbyentidResultModel) SetMsgInfo(v string) *AlibabaAlihealthDrugMscGetbyentidResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthDrugMscGetbyentidResultModel) SetMsgCode(v string) *AlibabaAlihealthDrugMscGetbyentidResultModel {
    s.MsgCode = &v
    return s
}
