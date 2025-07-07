package domain


type AlibabaAlihealthDrugMscGetentinfonewResultModel struct {
    /*
        是否成功     */
    ResponseSuccess  *bool `json:"response_success,omitempty" `

    /*
        错误编码     */
    MsgCode  *string `json:"msg_code,omitempty" `

    /*
        错误信息     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        响应结果     */
    Model  *AlibabaAlihealthDrugMscGetentinfonewTopEntInfoRespDto `json:"model,omitempty" `

}

func (s *AlibabaAlihealthDrugMscGetentinfonewResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthDrugMscGetentinfonewResultModel {
    s.ResponseSuccess = &v
    return s
}
func (s *AlibabaAlihealthDrugMscGetentinfonewResultModel) SetMsgCode(v string) *AlibabaAlihealthDrugMscGetentinfonewResultModel {
    s.MsgCode = &v
    return s
}
func (s *AlibabaAlihealthDrugMscGetentinfonewResultModel) SetMsgInfo(v string) *AlibabaAlihealthDrugMscGetentinfonewResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthDrugMscGetentinfonewResultModel) SetModel(v AlibabaAlihealthDrugMscGetentinfonewTopEntInfoRespDto) *AlibabaAlihealthDrugMscGetentinfonewResultModel {
    s.Model = &v
    return s
}
