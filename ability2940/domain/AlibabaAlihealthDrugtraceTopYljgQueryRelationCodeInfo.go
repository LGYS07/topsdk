package domain


type AlibabaAlihealthDrugtraceTopYljgQueryRelationCodeInfo struct {
    /*
        码状态（A:已激活;I:已核注;O:已核销;C:已注销;E:码不存在）     */
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

func (s *AlibabaAlihealthDrugtraceTopYljgQueryRelationCodeInfo) SetStatus(v string) *AlibabaAlihealthDrugtraceTopYljgQueryRelationCodeInfo {
    s.Status = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryRelationCodeInfo) SetCode(v string) *AlibabaAlihealthDrugtraceTopYljgQueryRelationCodeInfo {
    s.Code = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryRelationCodeInfo) SetCodeLevel(v string) *AlibabaAlihealthDrugtraceTopYljgQueryRelationCodeInfo {
    s.CodeLevel = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryRelationCodeInfo) SetParentCode(v string) *AlibabaAlihealthDrugtraceTopYljgQueryRelationCodeInfo {
    s.ParentCode = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryRelationCodeInfo) SetCodePackLevel(v string) *AlibabaAlihealthDrugtraceTopYljgQueryRelationCodeInfo {
    s.CodePackLevel = &v
    return s
}
