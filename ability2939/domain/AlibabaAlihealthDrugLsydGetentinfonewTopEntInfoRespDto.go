package domain


type AlibabaAlihealthDrugLsydGetentinfonewTopEntInfoRespDto struct {
    /*
        企业ID【ent_id】（单据上传时的收发货企业id就是填这个字段）     */
    EntId  *string `json:"ent_id,omitempty" `

    /*
        企业唯一标识【ref_ent_id】（单据上传时的货主企业ref_user_id就是填这个字段）     */
    RefEntId  *string `json:"ref_ent_id,omitempty" `

    /*
        1-审核通过，0-审核中，2-审核不通过     */
    AuditStatus  *int64 `json:"audit_status,omitempty" `

    /*
        企业名称     */
    EntName  *string `json:"ent_name,omitempty" `

    /*
        唯一代码     */
    UniqueCode  *string `json:"unique_code,omitempty" `

    /*
        唯一代码来源的资质名称（非精准）     */
    LicTypeName  *string `json:"lic_type_name,omitempty" `

    /*
        唯一代码来源的资质代码（非精准）     */
    LicTypeCode  *int64 `json:"lic_type_code,omitempty" `

    /*
        企业所在省份名称     */
    ProvName  *string `json:"prov_name,omitempty" `

    /*
        企业所在省份代码     */
    ProvCode  *string `json:"prov_code,omitempty" `

    /*
        企业所在城市名称     */
    CityName  *string `json:"city_name,omitempty" `

    /*
        企业所在城市代码     */
    CityCode  *string `json:"city_code,omitempty" `

    /*
        企业所在区县名称     */
    AreaName  *string `json:"area_name,omitempty" `

    /*
        企业所在区县代码     */
    AreaCode  *string `json:"area_code,omitempty" `

}

func (s *AlibabaAlihealthDrugLsydGetentinfonewTopEntInfoRespDto) SetEntId(v string) *AlibabaAlihealthDrugLsydGetentinfonewTopEntInfoRespDto {
    s.EntId = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydGetentinfonewTopEntInfoRespDto) SetRefEntId(v string) *AlibabaAlihealthDrugLsydGetentinfonewTopEntInfoRespDto {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydGetentinfonewTopEntInfoRespDto) SetAuditStatus(v int64) *AlibabaAlihealthDrugLsydGetentinfonewTopEntInfoRespDto {
    s.AuditStatus = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydGetentinfonewTopEntInfoRespDto) SetEntName(v string) *AlibabaAlihealthDrugLsydGetentinfonewTopEntInfoRespDto {
    s.EntName = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydGetentinfonewTopEntInfoRespDto) SetUniqueCode(v string) *AlibabaAlihealthDrugLsydGetentinfonewTopEntInfoRespDto {
    s.UniqueCode = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydGetentinfonewTopEntInfoRespDto) SetLicTypeName(v string) *AlibabaAlihealthDrugLsydGetentinfonewTopEntInfoRespDto {
    s.LicTypeName = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydGetentinfonewTopEntInfoRespDto) SetLicTypeCode(v int64) *AlibabaAlihealthDrugLsydGetentinfonewTopEntInfoRespDto {
    s.LicTypeCode = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydGetentinfonewTopEntInfoRespDto) SetProvName(v string) *AlibabaAlihealthDrugLsydGetentinfonewTopEntInfoRespDto {
    s.ProvName = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydGetentinfonewTopEntInfoRespDto) SetProvCode(v string) *AlibabaAlihealthDrugLsydGetentinfonewTopEntInfoRespDto {
    s.ProvCode = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydGetentinfonewTopEntInfoRespDto) SetCityName(v string) *AlibabaAlihealthDrugLsydGetentinfonewTopEntInfoRespDto {
    s.CityName = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydGetentinfonewTopEntInfoRespDto) SetCityCode(v string) *AlibabaAlihealthDrugLsydGetentinfonewTopEntInfoRespDto {
    s.CityCode = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydGetentinfonewTopEntInfoRespDto) SetAreaName(v string) *AlibabaAlihealthDrugLsydGetentinfonewTopEntInfoRespDto {
    s.AreaName = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydGetentinfonewTopEntInfoRespDto) SetAreaCode(v string) *AlibabaAlihealthDrugLsydGetentinfonewTopEntInfoRespDto {
    s.AreaCode = &v
    return s
}
