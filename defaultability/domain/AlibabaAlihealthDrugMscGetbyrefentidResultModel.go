package domain


type AlibabaAlihealthDrugMscGetbyrefentidResultModel struct {
    /*
        是否响应成功     */
    ResponseSuccess  *bool `json:"response_success,omitempty" `

    /*
        返回对象     */
    Model  *AlibabaAlihealthDrugMscGetbyrefentidPUserEntInfoDTO `json:"model,omitempty" `

    /*
        返回值     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        返回码     */
    MsgCode  *string `json:"msg_code,omitempty" `

}

func (s *AlibabaAlihealthDrugMscGetbyrefentidResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthDrugMscGetbyrefentidResultModel {
    s.ResponseSuccess = &v
    return s
}
func (s *AlibabaAlihealthDrugMscGetbyrefentidResultModel) SetModel(v AlibabaAlihealthDrugMscGetbyrefentidPUserEntInfoDTO) *AlibabaAlihealthDrugMscGetbyrefentidResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthDrugMscGetbyrefentidResultModel) SetMsgInfo(v string) *AlibabaAlihealthDrugMscGetbyrefentidResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthDrugMscGetbyrefentidResultModel) SetMsgCode(v string) *AlibabaAlihealthDrugMscGetbyrefentidResultModel {
    s.MsgCode = &v
    return s
}
