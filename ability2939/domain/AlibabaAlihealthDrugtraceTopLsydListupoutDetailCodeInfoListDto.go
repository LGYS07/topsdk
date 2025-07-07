package domain


type AlibabaAlihealthDrugtraceTopLsydListupoutDetailCodeInfoListDto struct {
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

func (s *AlibabaAlihealthDrugtraceTopLsydListupoutDetailCodeInfoListDto) SetPrepnSpec(v string) *AlibabaAlihealthDrugtraceTopLsydListupoutDetailCodeInfoListDto {
    s.PrepnSpec = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydListupoutDetailCodeInfoListDto) SetPrepnAmount(v string) *AlibabaAlihealthDrugtraceTopLsydListupoutDetailCodeInfoListDto {
    s.PrepnAmount = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydListupoutDetailCodeInfoListDto) SetPkgAmount(v string) *AlibabaAlihealthDrugtraceTopLsydListupoutDetailCodeInfoListDto {
    s.PkgAmount = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydListupoutDetailCodeInfoListDto) SetCodeLevel(v string) *AlibabaAlihealthDrugtraceTopLsydListupoutDetailCodeInfoListDto {
    s.CodeLevel = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydListupoutDetailCodeInfoListDto) SetCode(v string) *AlibabaAlihealthDrugtraceTopLsydListupoutDetailCodeInfoListDto {
    s.Code = &v
    return s
}
