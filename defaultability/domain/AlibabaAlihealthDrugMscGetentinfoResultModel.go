package domain


type AlibabaAlihealthDrugMscGetentinfoResultModel struct {
    /*
        是否响应成功     */
    ResponseSuccess  *bool `json:"response_success,omitempty" `

    /*
        返回对象     */
    Model  *AlibabaAlihealthDrugMscGetentinfoPUserEntInfoDTO `json:"model,omitempty" `

    /*
        返回信息     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        返回码     */
    MsgCode  *string `json:"msg_code,omitempty" `

}

func (s *AlibabaAlihealthDrugMscGetentinfoResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthDrugMscGetentinfoResultModel {
    s.ResponseSuccess = &v
    return s
}
func (s *AlibabaAlihealthDrugMscGetentinfoResultModel) SetModel(v AlibabaAlihealthDrugMscGetentinfoPUserEntInfoDTO) *AlibabaAlihealthDrugMscGetentinfoResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthDrugMscGetentinfoResultModel) SetMsgInfo(v string) *AlibabaAlihealthDrugMscGetentinfoResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthDrugMscGetentinfoResultModel) SetMsgCode(v string) *AlibabaAlihealthDrugMscGetentinfoResultModel {
    s.MsgCode = &v
    return s
}
