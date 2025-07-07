package domain


type AlibabaAlihealthDrugtraceTopLsydQueryCodedetailPUserEntDto struct {
    /*
        企业id     */
    RefEntId  *string `json:"ref_ent_id,omitempty" `

    /*
        企业名称     */
    EntName  *string `json:"ent_name,omitempty" `

}

func (s *AlibabaAlihealthDrugtraceTopLsydQueryCodedetailPUserEntDto) SetRefEntId(v string) *AlibabaAlihealthDrugtraceTopLsydQueryCodedetailPUserEntDto {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryCodedetailPUserEntDto) SetEntName(v string) *AlibabaAlihealthDrugtraceTopLsydQueryCodedetailPUserEntDto {
    s.EntName = &v
    return s
}
