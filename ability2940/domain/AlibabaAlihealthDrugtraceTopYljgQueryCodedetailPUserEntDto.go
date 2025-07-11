package domain


type AlibabaAlihealthDrugtraceTopYljgQueryCodedetailPUserEntDto struct {
    /*
        生产企业entid，不是refentid     */
    RefEntId  *string `json:"ref_ent_id,omitempty" `

    /*
        生产企业名称     */
    EntName  *string `json:"ent_name,omitempty" `

}

func (s *AlibabaAlihealthDrugtraceTopYljgQueryCodedetailPUserEntDto) SetRefEntId(v string) *AlibabaAlihealthDrugtraceTopYljgQueryCodedetailPUserEntDto {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryCodedetailPUserEntDto) SetEntName(v string) *AlibabaAlihealthDrugtraceTopYljgQueryCodedetailPUserEntDto {
    s.EntName = &v
    return s
}
