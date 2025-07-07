package domain


type AlibabaAlihealthDrugtraceTopLsydQueryGetentinfoPUserEntInfoDto struct {
    /*
        企业唯一标识【ref_ent_id】     */
    RefEntId  *string `json:"ref_ent_id,omitempty" `

    /*
        企业ID【ent_id】     */
    EntId  *string `json:"ent_id,omitempty" `

    /*
        入网标识【0非入网;1开头代表入网企业;200代表入驻码上放心的企业】     */
    NetworkType  *string `json:"network_type,omitempty" `

    /*
        1-审核通过，0-审核中，2-审核不通过     */
    AuditStatus  *string `json:"audit_status,omitempty" `

}

func (s *AlibabaAlihealthDrugtraceTopLsydQueryGetentinfoPUserEntInfoDto) SetRefEntId(v string) *AlibabaAlihealthDrugtraceTopLsydQueryGetentinfoPUserEntInfoDto {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryGetentinfoPUserEntInfoDto) SetEntId(v string) *AlibabaAlihealthDrugtraceTopLsydQueryGetentinfoPUserEntInfoDto {
    s.EntId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryGetentinfoPUserEntInfoDto) SetNetworkType(v string) *AlibabaAlihealthDrugtraceTopLsydQueryGetentinfoPUserEntInfoDto {
    s.NetworkType = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryGetentinfoPUserEntInfoDto) SetAuditStatus(v string) *AlibabaAlihealthDrugtraceTopLsydQueryGetentinfoPUserEntInfoDto {
    s.AuditStatus = &v
    return s
}
