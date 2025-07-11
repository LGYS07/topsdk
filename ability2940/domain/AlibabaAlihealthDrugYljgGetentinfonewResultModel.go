package domain


type AlibabaAlihealthDrugYljgGetentinfonewResultModel struct {
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
    Model  *AlibabaAlihealthDrugYljgGetentinfonewTopEntInfoRespDto `json:"model,omitempty" `

}

func (s *AlibabaAlihealthDrugYljgGetentinfonewResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthDrugYljgGetentinfonewResultModel {
    s.ResponseSuccess = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgGetentinfonewResultModel) SetMsgCode(v string) *AlibabaAlihealthDrugYljgGetentinfonewResultModel {
    s.MsgCode = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgGetentinfonewResultModel) SetMsgInfo(v string) *AlibabaAlihealthDrugYljgGetentinfonewResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgGetentinfonewResultModel) SetModel(v AlibabaAlihealthDrugYljgGetentinfonewTopEntInfoRespDto) *AlibabaAlihealthDrugYljgGetentinfonewResultModel {
    s.Model = &v
    return s
}
