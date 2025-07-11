package domain


type AlibabaAlihealthDrugtraceTopYljgListupoutDetailCodeInfoListDto struct {
    /*
        制剂规格     */
    PrepnSpec  *string `json:"prepn_spec,omitempty" `

    /*
        最小制剂数量     */
    PrepnAmount  *string `json:"prepn_amount,omitempty" `

    /*
        最小包装数量     */
    PkgAmount  *string `json:"pkg_amount,omitempty" `

    /*
        监管码级别     */
    CodeLevel  *string `json:"code_level,omitempty" `

    /*
        监管码     */
    Code  *string `json:"code,omitempty" `

}

func (s *AlibabaAlihealthDrugtraceTopYljgListupoutDetailCodeInfoListDto) SetPrepnSpec(v string) *AlibabaAlihealthDrugtraceTopYljgListupoutDetailCodeInfoListDto {
    s.PrepnSpec = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgListupoutDetailCodeInfoListDto) SetPrepnAmount(v string) *AlibabaAlihealthDrugtraceTopYljgListupoutDetailCodeInfoListDto {
    s.PrepnAmount = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgListupoutDetailCodeInfoListDto) SetPkgAmount(v string) *AlibabaAlihealthDrugtraceTopYljgListupoutDetailCodeInfoListDto {
    s.PkgAmount = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgListupoutDetailCodeInfoListDto) SetCodeLevel(v string) *AlibabaAlihealthDrugtraceTopYljgListupoutDetailCodeInfoListDto {
    s.CodeLevel = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgListupoutDetailCodeInfoListDto) SetCode(v string) *AlibabaAlihealthDrugtraceTopYljgListupoutDetailCodeInfoListDto {
    s.Code = &v
    return s
}
