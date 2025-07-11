package domain


type AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailCodeandparentlist struct {
    /*
        追溯码     */
    Code  *string `json:"code,omitempty" `

    /*
        码级别     */
    CodeLevel  *string `json:"code_level,omitempty" `

    /*
        父码     */
    ParentCode  *string `json:"parent_code,omitempty" `

}

func (s *AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailCodeandparentlist) SetCode(v string) *AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailCodeandparentlist {
    s.Code = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailCodeandparentlist) SetCodeLevel(v string) *AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailCodeandparentlist {
    s.CodeLevel = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailCodeandparentlist) SetParentCode(v string) *AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailCodeandparentlist {
    s.ParentCode = &v
    return s
}
