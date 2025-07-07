package domain


type AlibabaAlihealthDrugMscSearchbillDetailCodeandparentlist struct {
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

func (s *AlibabaAlihealthDrugMscSearchbillDetailCodeandparentlist) SetCode(v string) *AlibabaAlihealthDrugMscSearchbillDetailCodeandparentlist {
    s.Code = &v
    return s
}
func (s *AlibabaAlihealthDrugMscSearchbillDetailCodeandparentlist) SetCodeLevel(v string) *AlibabaAlihealthDrugMscSearchbillDetailCodeandparentlist {
    s.CodeLevel = &v
    return s
}
func (s *AlibabaAlihealthDrugMscSearchbillDetailCodeandparentlist) SetParentCode(v string) *AlibabaAlihealthDrugMscSearchbillDetailCodeandparentlist {
    s.ParentCode = &v
    return s
}
