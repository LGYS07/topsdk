package domain


type AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailCodeandparentlist struct {
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

func (s *AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailCodeandparentlist) SetCode(v string) *AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailCodeandparentlist {
    s.Code = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailCodeandparentlist) SetCodeLevel(v string) *AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailCodeandparentlist {
    s.CodeLevel = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailCodeandparentlist) SetParentCode(v string) *AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailCodeandparentlist {
    s.ParentCode = &v
    return s
}
