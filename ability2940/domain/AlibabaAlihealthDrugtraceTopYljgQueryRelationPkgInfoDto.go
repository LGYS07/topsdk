package domain


type AlibabaAlihealthDrugtraceTopYljgQueryRelationPkgInfoDto struct {
    /*
        码信息     */
    CodeList  *[]string `json:"code_list,omitempty" `

}

func (s *AlibabaAlihealthDrugtraceTopYljgQueryRelationPkgInfoDto) SetCodeList(v []string) *AlibabaAlihealthDrugtraceTopYljgQueryRelationPkgInfoDto {
    s.CodeList = &v
    return s
}
