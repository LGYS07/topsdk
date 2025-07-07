package domain


type AlibabaAlihealthDrugtraceTopLsydQueryRelationPkgInfoDto struct {
    /*
        码信息     */
    CodeList  *[]string `json:"code_list,omitempty" `

}

func (s *AlibabaAlihealthDrugtraceTopLsydQueryRelationPkgInfoDto) SetCodeList(v []string) *AlibabaAlihealthDrugtraceTopLsydQueryRelationPkgInfoDto {
    s.CodeList = &v
    return s
}
