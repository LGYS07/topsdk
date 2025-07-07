package domain


type AlibabaAlihealthDrugLsydGetentinfonewResultModel struct {
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
    Model  *AlibabaAlihealthDrugLsydGetentinfonewTopEntInfoRespDto `json:"model,omitempty" `

}

func (s *AlibabaAlihealthDrugLsydGetentinfonewResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthDrugLsydGetentinfonewResultModel {
    s.ResponseSuccess = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydGetentinfonewResultModel) SetMsgCode(v string) *AlibabaAlihealthDrugLsydGetentinfonewResultModel {
    s.MsgCode = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydGetentinfonewResultModel) SetMsgInfo(v string) *AlibabaAlihealthDrugLsydGetentinfonewResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydGetentinfonewResultModel) SetModel(v AlibabaAlihealthDrugLsydGetentinfonewTopEntInfoRespDto) *AlibabaAlihealthDrugLsydGetentinfonewResultModel {
    s.Model = &v
    return s
}
