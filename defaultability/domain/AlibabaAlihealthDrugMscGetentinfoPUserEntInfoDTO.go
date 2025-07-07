package domain


type AlibabaAlihealthDrugMscGetentinfoPUserEntInfoDTO struct {
    /*
        企业唯一标识【ref_ent_id】（单据上传时的货主企业ref_user_id就是填这个字段）     */
    RefEntId  *string `json:"ref_ent_id,omitempty" `

    /*
        企业ID【ent_id】（单据上传时的收发货企业id就是填这个字段）     */
    EntId  *string `json:"ent_id,omitempty" `

    /*
        入网标识【0非入网;1开头代表入网企业;200代表入驻码上放心的企业】     */
    NetworkType  *string `json:"network_type,omitempty" `

    /*
        1-审核通过，0-审核中，2-审核不通过     */
    AuditStatus  *string `json:"audit_status,omitempty" `

}

func (s *AlibabaAlihealthDrugMscGetentinfoPUserEntInfoDTO) SetRefEntId(v string) *AlibabaAlihealthDrugMscGetentinfoPUserEntInfoDTO {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugMscGetentinfoPUserEntInfoDTO) SetEntId(v string) *AlibabaAlihealthDrugMscGetentinfoPUserEntInfoDTO {
    s.EntId = &v
    return s
}
func (s *AlibabaAlihealthDrugMscGetentinfoPUserEntInfoDTO) SetNetworkType(v string) *AlibabaAlihealthDrugMscGetentinfoPUserEntInfoDTO {
    s.NetworkType = &v
    return s
}
func (s *AlibabaAlihealthDrugMscGetentinfoPUserEntInfoDTO) SetAuditStatus(v string) *AlibabaAlihealthDrugMscGetentinfoPUserEntInfoDTO {
    s.AuditStatus = &v
    return s
}
