package domain


type AlibabaAlihealthDrugMscGetentinfonewTopEntInfoReqDto struct {
    /*
        查询参数：企业名称，无其他查询条件时不能为空     */
    EntName  *string `json:"ent_name,omitempty" `

    /*
        查询参数：统一社会信用代码，无其他查询条件时不能为空     */
    OrgCode  *string `json:"org_code,omitempty" `

    /*
        查询参数：诊所备案号或医疗单位登记号，无其他查询条件时不能为空     */
    MedicalCode  *string `json:"medical_code,omitempty" `

}

func (s *AlibabaAlihealthDrugMscGetentinfonewTopEntInfoReqDto) SetEntName(v string) *AlibabaAlihealthDrugMscGetentinfonewTopEntInfoReqDto {
    s.EntName = &v
    return s
}
func (s *AlibabaAlihealthDrugMscGetentinfonewTopEntInfoReqDto) SetOrgCode(v string) *AlibabaAlihealthDrugMscGetentinfonewTopEntInfoReqDto {
    s.OrgCode = &v
    return s
}
func (s *AlibabaAlihealthDrugMscGetentinfonewTopEntInfoReqDto) SetMedicalCode(v string) *AlibabaAlihealthDrugMscGetentinfonewTopEntInfoReqDto {
    s.MedicalCode = &v
    return s
}
