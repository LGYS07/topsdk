package domain


type AlibabaAlihealthDrugLsydGetentinfonewTopEntInfoReqDto struct {
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

func (s *AlibabaAlihealthDrugLsydGetentinfonewTopEntInfoReqDto) SetEntName(v string) *AlibabaAlihealthDrugLsydGetentinfonewTopEntInfoReqDto {
    s.EntName = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydGetentinfonewTopEntInfoReqDto) SetOrgCode(v string) *AlibabaAlihealthDrugLsydGetentinfonewTopEntInfoReqDto {
    s.OrgCode = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydGetentinfonewTopEntInfoReqDto) SetMedicalCode(v string) *AlibabaAlihealthDrugLsydGetentinfonewTopEntInfoReqDto {
    s.MedicalCode = &v
    return s
}
