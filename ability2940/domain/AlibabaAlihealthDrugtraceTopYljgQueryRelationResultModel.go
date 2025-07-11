package domain


type AlibabaAlihealthDrugtraceTopYljgQueryRelationResultModel struct {
    /*
        model     */
    ModelList  *[]AlibabaAlihealthDrugtraceTopYljgQueryRelationCodeRelationDto `json:"model_list,omitempty" `

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

func (s *AlibabaAlihealthDrugtraceTopYljgQueryRelationResultModel) SetModelList(v []AlibabaAlihealthDrugtraceTopYljgQueryRelationCodeRelationDto) *AlibabaAlihealthDrugtraceTopYljgQueryRelationResultModel {
    s.ModelList = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryRelationResultModel) SetMsgCode(v string) *AlibabaAlihealthDrugtraceTopYljgQueryRelationResultModel {
    s.MsgCode = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryRelationResultModel) SetMsgInfo(v string) *AlibabaAlihealthDrugtraceTopYljgQueryRelationResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryRelationResultModel) SetResponseSuccess(v bool) *AlibabaAlihealthDrugtraceTopYljgQueryRelationResultModel {
    s.ResponseSuccess = &v
    return s
}
