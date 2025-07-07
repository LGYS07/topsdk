package domain


type AlibabaAlihealthDrugtraceTopLsydQueryRelationCodeInfo struct {
    /*
        码状态(A.已激活,I.已核注,O.已核销,X.已销毁,N.已停用,C.已注销,E.不存在,B.申请未激活)     */
    Status  *string `json:"status,omitempty" `

    /*
        追溯码     */
    Code  *string `json:"code,omitempty" `

    /*
        包装等级     */
    CodeLevel  *string `json:"code_level,omitempty" `

    /*
        父码     */
    ParentCode  *string `json:"parent_code,omitempty" `

    /*
        码等级：1代表最小码     */
    CodePackLevel  *string `json:"code_pack_level,omitempty" `

}

func (s *AlibabaAlihealthDrugtraceTopLsydQueryRelationCodeInfo) SetStatus(v string) *AlibabaAlihealthDrugtraceTopLsydQueryRelationCodeInfo {
    s.Status = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryRelationCodeInfo) SetCode(v string) *AlibabaAlihealthDrugtraceTopLsydQueryRelationCodeInfo {
    s.Code = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryRelationCodeInfo) SetCodeLevel(v string) *AlibabaAlihealthDrugtraceTopLsydQueryRelationCodeInfo {
    s.CodeLevel = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryRelationCodeInfo) SetParentCode(v string) *AlibabaAlihealthDrugtraceTopLsydQueryRelationCodeInfo {
    s.ParentCode = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryRelationCodeInfo) SetCodePackLevel(v string) *AlibabaAlihealthDrugtraceTopLsydQueryRelationCodeInfo {
    s.CodePackLevel = &v
    return s
}
