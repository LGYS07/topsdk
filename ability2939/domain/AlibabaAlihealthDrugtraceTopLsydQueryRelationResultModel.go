package domain


type AlibabaAlihealthDrugtraceTopLsydQueryRelationResultModel struct {
    /*
        model     */
    ModelList  *[]AlibabaAlihealthDrugtraceTopLsydQueryRelationCodeRelationDto `json:"model_list,omitempty" `

    /*
        msgCode     */
    MsgCode  *string `json:"msg_code,omitempty" `

    /*
        msgInfo     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        是否成功     */
    ResponseSuccess  *bool `json:"response_success,omitempty" `

}

func (s *AlibabaAlihealthDrugtraceTopLsydQueryRelationResultModel) SetModelList(v []AlibabaAlihealthDrugtraceTopLsydQueryRelationCodeRelationDto) *AlibabaAlihealthDrugtraceTopLsydQueryRelationResultModel {
    s.ModelList = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryRelationResultModel) SetMsgCode(v string) *AlibabaAlihealthDrugtraceTopLsydQueryRelationResultModel {
    s.MsgCode = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryRelationResultModel) SetMsgInfo(v string) *AlibabaAlihealthDrugtraceTopLsydQueryRelationResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryRelationResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthDrugtraceTopLsydQueryRelationResultModel {
    s.ResponseSuccess = &v
    return s
}
