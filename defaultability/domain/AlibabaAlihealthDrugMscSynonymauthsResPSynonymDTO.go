package domain


type AlibabaAlihealthDrugMscSynonymauthsResPSynonymDTO struct {
    /*
        企业名称     */
    EntName  *string `json:"ent_name,omitempty" `

    /*
        市     */
    CityDesc  *string `json:"city_desc,omitempty" `

    /*
        省     */
    ProvDesc  *string `json:"prov_desc,omitempty" `

    /*
        区     */
    AreaDesc  *string `json:"area_desc,omitempty" `

    /*
        企业id     */
    EntId  *string `json:"ent_id,omitempty" `

    /*
        区域编码     */
    DictRegionCode  *string `json:"dict_region_code,omitempty" `

    /*
        企业唯一标识     */
    RefEntId  *string `json:"ref_ent_id,omitempty" `

    /*
        货主     */
    SynOwnEntId  *string `json:"syn_own_ent_id,omitempty" `

    /*
        货主标识     */
    UserEntId  *string `json:"user_ent_id,omitempty" `

    /*
        创建日期     */
    CrtDate  *string `json:"crt_date,omitempty" `

    /*
        角色     */
    UserRoleType  *string `json:"user_role_type,omitempty" `

}

func (s *AlibabaAlihealthDrugMscSynonymauthsResPSynonymDTO) SetEntName(v string) *AlibabaAlihealthDrugMscSynonymauthsResPSynonymDTO {
    s.EntName = &v
    return s
}
func (s *AlibabaAlihealthDrugMscSynonymauthsResPSynonymDTO) SetCityDesc(v string) *AlibabaAlihealthDrugMscSynonymauthsResPSynonymDTO {
    s.CityDesc = &v
    return s
}
func (s *AlibabaAlihealthDrugMscSynonymauthsResPSynonymDTO) SetProvDesc(v string) *AlibabaAlihealthDrugMscSynonymauthsResPSynonymDTO {
    s.ProvDesc = &v
    return s
}
func (s *AlibabaAlihealthDrugMscSynonymauthsResPSynonymDTO) SetAreaDesc(v string) *AlibabaAlihealthDrugMscSynonymauthsResPSynonymDTO {
    s.AreaDesc = &v
    return s
}
func (s *AlibabaAlihealthDrugMscSynonymauthsResPSynonymDTO) SetEntId(v string) *AlibabaAlihealthDrugMscSynonymauthsResPSynonymDTO {
    s.EntId = &v
    return s
}
func (s *AlibabaAlihealthDrugMscSynonymauthsResPSynonymDTO) SetDictRegionCode(v string) *AlibabaAlihealthDrugMscSynonymauthsResPSynonymDTO {
    s.DictRegionCode = &v
    return s
}
func (s *AlibabaAlihealthDrugMscSynonymauthsResPSynonymDTO) SetRefEntId(v string) *AlibabaAlihealthDrugMscSynonymauthsResPSynonymDTO {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugMscSynonymauthsResPSynonymDTO) SetSynOwnEntId(v string) *AlibabaAlihealthDrugMscSynonymauthsResPSynonymDTO {
    s.SynOwnEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugMscSynonymauthsResPSynonymDTO) SetUserEntId(v string) *AlibabaAlihealthDrugMscSynonymauthsResPSynonymDTO {
    s.UserEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugMscSynonymauthsResPSynonymDTO) SetCrtDate(v string) *AlibabaAlihealthDrugMscSynonymauthsResPSynonymDTO {
    s.CrtDate = &v
    return s
}
func (s *AlibabaAlihealthDrugMscSynonymauthsResPSynonymDTO) SetUserRoleType(v string) *AlibabaAlihealthDrugMscSynonymauthsResPSynonymDTO {
    s.UserRoleType = &v
    return s
}
