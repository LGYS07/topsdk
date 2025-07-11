package domain


type AlibabaAlihealthDrugtraceTopYljgQueryGetentinfoPUserEntInfoDto struct {
    /*
        企业唯一标识【ref_ent_id】     */
    RefEntId  *string `json:"ref_ent_id,omitempty" `

    /*
        企业ID【ent_id】     */
    EntId  *string `json:"ent_id,omitempty" `

    /*
        入网标识【0非入网;1开头代表入网企业;200代表入驻马上放心的企业】     */
    NetworkType  *string `json:"network_type,omitempty" `

    /*
        1-审核通过，0-审核中，2-审核不通过     */
    AuditStatus  *string `json:"audit_status,omitempty" `

}

func (s *AlibabaAlihealthDrugtraceTopYljgQueryGetentinfoPUserEntInfoDto) SetRefEntId(v string) *AlibabaAlihealthDrugtraceTopYljgQueryGetentinfoPUserEntInfoDto {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryGetentinfoPUserEntInfoDto) SetEntId(v string) *AlibabaAlihealthDrugtraceTopYljgQueryGetentinfoPUserEntInfoDto {
    s.EntId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryGetentinfoPUserEntInfoDto) SetNetworkType(v string) *AlibabaAlihealthDrugtraceTopYljgQueryGetentinfoPUserEntInfoDto {
    s.NetworkType = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryGetentinfoPUserEntInfoDto) SetAuditStatus(v string) *AlibabaAlihealthDrugtraceTopYljgQueryGetentinfoPUserEntInfoDto {
    s.AuditStatus = &v
    return s
}
