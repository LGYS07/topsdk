package domain


type AlibabaAlihealthDrugMscBilloutDetailwithcodesCodeandparentlist struct {
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

func (s *AlibabaAlihealthDrugMscBilloutDetailwithcodesCodeandparentlist) SetCode(v string) *AlibabaAlihealthDrugMscBilloutDetailwithcodesCodeandparentlist {
    s.Code = &v
    return s
}
func (s *AlibabaAlihealthDrugMscBilloutDetailwithcodesCodeandparentlist) SetCodeLevel(v string) *AlibabaAlihealthDrugMscBilloutDetailwithcodesCodeandparentlist {
    s.CodeLevel = &v
    return s
}
func (s *AlibabaAlihealthDrugMscBilloutDetailwithcodesCodeandparentlist) SetParentCode(v string) *AlibabaAlihealthDrugMscBilloutDetailwithcodesCodeandparentlist {
    s.ParentCode = &v
    return s
}
