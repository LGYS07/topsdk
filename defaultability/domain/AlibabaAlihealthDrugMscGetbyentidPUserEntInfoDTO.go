package domain


type AlibabaAlihealthDrugMscGetbyentidPUserEntInfoDTO struct {
    /*
        所在地编码     */
    DictRegionCode  *string `json:"dict_region_code,omitempty" `

    /*
        所在地明细     */
    RegRegionDetail  *string `json:"reg_region_detail,omitempty" `

    /*
        注册地编码     */
    RegRegionCode  *string `json:"reg_region_code,omitempty" `

    /*
        注册地明细     */
    OrgCode  *string `json:"org_code,omitempty" `

    /*
        是否法人     */
    LegalOrgFlag  *string `json:"legal_org_flag,omitempty" `

    /*
        所属管理机构     */
    DirectManage  *string `json:"direct_manage,omitempty" `

    /*
        企业名称     */
    EntName  *string `json:"ent_name,omitempty" `

    /*
        是否入网     */
    IsNetwork  *string `json:"is_network,omitempty" `

    /*
        企业唯一标识     */
    RefEntId  *string `json:"ref_ent_id,omitempty" `

    /*
        企业id     */
    EntId  *string `json:"ent_id,omitempty" `

    /*
        所在地明细     */
    DictRegionDetail  *string `json:"dict_region_detail,omitempty" `

    /*
        状态1.使用中0.已废除     */
    Status  *string `json:"status,omitempty" `

    /*
        拼音缩写     */
    EntCapitalName  *string `json:"ent_capital_name,omitempty" `

    /*
        企业类型     */
    UserRoleType  *string `json:"user_role_type,omitempty" `

    /*
        企业类型编码     */
    UserRoleTypeStr  *string `json:"user_role_type_str,omitempty" `

    /*
        企业机构详细类别     */
    EntOrgType  *string `json:"ent_org_type,omitempty" `

    /*
        省     */
    ProvName  *string `json:"prov_name,omitempty" `

    /*
        市     */
    AreaName  *string `json:"area_name,omitempty" `

    /*
        县     */
    CityName  *string `json:"city_name,omitempty" `

}

func (s *AlibabaAlihealthDrugMscGetbyentidPUserEntInfoDTO) SetDictRegionCode(v string) *AlibabaAlihealthDrugMscGetbyentidPUserEntInfoDTO {
    s.DictRegionCode = &v
    return s
}
func (s *AlibabaAlihealthDrugMscGetbyentidPUserEntInfoDTO) SetRegRegionDetail(v string) *AlibabaAlihealthDrugMscGetbyentidPUserEntInfoDTO {
    s.RegRegionDetail = &v
    return s
}
func (s *AlibabaAlihealthDrugMscGetbyentidPUserEntInfoDTO) SetRegRegionCode(v string) *AlibabaAlihealthDrugMscGetbyentidPUserEntInfoDTO {
    s.RegRegionCode = &v
    return s
}
func (s *AlibabaAlihealthDrugMscGetbyentidPUserEntInfoDTO) SetOrgCode(v string) *AlibabaAlihealthDrugMscGetbyentidPUserEntInfoDTO {
    s.OrgCode = &v
    return s
}
func (s *AlibabaAlihealthDrugMscGetbyentidPUserEntInfoDTO) SetLegalOrgFlag(v string) *AlibabaAlihealthDrugMscGetbyentidPUserEntInfoDTO {
    s.LegalOrgFlag = &v
    return s
}
func (s *AlibabaAlihealthDrugMscGetbyentidPUserEntInfoDTO) SetDirectManage(v string) *AlibabaAlihealthDrugMscGetbyentidPUserEntInfoDTO {
    s.DirectManage = &v
    return s
}
func (s *AlibabaAlihealthDrugMscGetbyentidPUserEntInfoDTO) SetEntName(v string) *AlibabaAlihealthDrugMscGetbyentidPUserEntInfoDTO {
    s.EntName = &v
    return s
}
func (s *AlibabaAlihealthDrugMscGetbyentidPUserEntInfoDTO) SetIsNetwork(v string) *AlibabaAlihealthDrugMscGetbyentidPUserEntInfoDTO {
    s.IsNetwork = &v
    return s
}
func (s *AlibabaAlihealthDrugMscGetbyentidPUserEntInfoDTO) SetRefEntId(v string) *AlibabaAlihealthDrugMscGetbyentidPUserEntInfoDTO {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugMscGetbyentidPUserEntInfoDTO) SetEntId(v string) *AlibabaAlihealthDrugMscGetbyentidPUserEntInfoDTO {
    s.EntId = &v
    return s
}
func (s *AlibabaAlihealthDrugMscGetbyentidPUserEntInfoDTO) SetDictRegionDetail(v string) *AlibabaAlihealthDrugMscGetbyentidPUserEntInfoDTO {
    s.DictRegionDetail = &v
    return s
}
func (s *AlibabaAlihealthDrugMscGetbyentidPUserEntInfoDTO) SetStatus(v string) *AlibabaAlihealthDrugMscGetbyentidPUserEntInfoDTO {
    s.Status = &v
    return s
}
func (s *AlibabaAlihealthDrugMscGetbyentidPUserEntInfoDTO) SetEntCapitalName(v string) *AlibabaAlihealthDrugMscGetbyentidPUserEntInfoDTO {
    s.EntCapitalName = &v
    return s
}
func (s *AlibabaAlihealthDrugMscGetbyentidPUserEntInfoDTO) SetUserRoleType(v string) *AlibabaAlihealthDrugMscGetbyentidPUserEntInfoDTO {
    s.UserRoleType = &v
    return s
}
func (s *AlibabaAlihealthDrugMscGetbyentidPUserEntInfoDTO) SetUserRoleTypeStr(v string) *AlibabaAlihealthDrugMscGetbyentidPUserEntInfoDTO {
    s.UserRoleTypeStr = &v
    return s
}
func (s *AlibabaAlihealthDrugMscGetbyentidPUserEntInfoDTO) SetEntOrgType(v string) *AlibabaAlihealthDrugMscGetbyentidPUserEntInfoDTO {
    s.EntOrgType = &v
    return s
}
func (s *AlibabaAlihealthDrugMscGetbyentidPUserEntInfoDTO) SetProvName(v string) *AlibabaAlihealthDrugMscGetbyentidPUserEntInfoDTO {
    s.ProvName = &v
    return s
}
func (s *AlibabaAlihealthDrugMscGetbyentidPUserEntInfoDTO) SetAreaName(v string) *AlibabaAlihealthDrugMscGetbyentidPUserEntInfoDTO {
    s.AreaName = &v
    return s
}
func (s *AlibabaAlihealthDrugMscGetbyentidPUserEntInfoDTO) SetCityName(v string) *AlibabaAlihealthDrugMscGetbyentidPUserEntInfoDTO {
    s.CityName = &v
    return s
}
