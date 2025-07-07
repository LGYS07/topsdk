package domain


type AlibabaAlihealthDrugMscGetentinfonewTopEntInfoRespDto struct {
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

func (s *AlibabaAlihealthDrugMscGetentinfonewTopEntInfoRespDto) SetEntId(v string) *AlibabaAlihealthDrugMscGetentinfonewTopEntInfoRespDto {
    s.EntId = &v
    return s
}
func (s *AlibabaAlihealthDrugMscGetentinfonewTopEntInfoRespDto) SetRefEntId(v string) *AlibabaAlihealthDrugMscGetentinfonewTopEntInfoRespDto {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugMscGetentinfonewTopEntInfoRespDto) SetAuditStatus(v int64) *AlibabaAlihealthDrugMscGetentinfonewTopEntInfoRespDto {
    s.AuditStatus = &v
    return s
}
func (s *AlibabaAlihealthDrugMscGetentinfonewTopEntInfoRespDto) SetEntName(v string) *AlibabaAlihealthDrugMscGetentinfonewTopEntInfoRespDto {
    s.EntName = &v
    return s
}
func (s *AlibabaAlihealthDrugMscGetentinfonewTopEntInfoRespDto) SetUniqueCode(v string) *AlibabaAlihealthDrugMscGetentinfonewTopEntInfoRespDto {
    s.UniqueCode = &v
    return s
}
func (s *AlibabaAlihealthDrugMscGetentinfonewTopEntInfoRespDto) SetLicTypeName(v string) *AlibabaAlihealthDrugMscGetentinfonewTopEntInfoRespDto {
    s.LicTypeName = &v
    return s
}
func (s *AlibabaAlihealthDrugMscGetentinfonewTopEntInfoRespDto) SetLicTypeCode(v int64) *AlibabaAlihealthDrugMscGetentinfonewTopEntInfoRespDto {
    s.LicTypeCode = &v
    return s
}
func (s *AlibabaAlihealthDrugMscGetentinfonewTopEntInfoRespDto) SetProvName(v string) *AlibabaAlihealthDrugMscGetentinfonewTopEntInfoRespDto {
    s.ProvName = &v
    return s
}
func (s *AlibabaAlihealthDrugMscGetentinfonewTopEntInfoRespDto) SetProvCode(v string) *AlibabaAlihealthDrugMscGetentinfonewTopEntInfoRespDto {
    s.ProvCode = &v
    return s
}
func (s *AlibabaAlihealthDrugMscGetentinfonewTopEntInfoRespDto) SetCityName(v string) *AlibabaAlihealthDrugMscGetentinfonewTopEntInfoRespDto {
    s.CityName = &v
    return s
}
func (s *AlibabaAlihealthDrugMscGetentinfonewTopEntInfoRespDto) SetCityCode(v string) *AlibabaAlihealthDrugMscGetentinfonewTopEntInfoRespDto {
    s.CityCode = &v
    return s
}
func (s *AlibabaAlihealthDrugMscGetentinfonewTopEntInfoRespDto) SetAreaName(v string) *AlibabaAlihealthDrugMscGetentinfonewTopEntInfoRespDto {
    s.AreaName = &v
    return s
}
func (s *AlibabaAlihealthDrugMscGetentinfonewTopEntInfoRespDto) SetAreaCode(v string) *AlibabaAlihealthDrugMscGetentinfonewTopEntInfoRespDto {
    s.AreaCode = &v
    return s
}
