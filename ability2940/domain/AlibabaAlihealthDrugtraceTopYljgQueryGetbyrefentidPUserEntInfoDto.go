package domain


type AlibabaAlihealthDrugtraceTopYljgQueryGetbyrefentidPUserEntInfoDto struct {
    /*
        企业id     */
    EntId  *string `json:"ent_id,omitempty" `

    /*
        企业唯一标识     */
    RefEntId  *string `json:"ref_ent_id,omitempty" `

    /*
        是否入网     */
    IsNetwork  *string `json:"is_network,omitempty" `

    /*
        企业名称     */
    EntName  *string `json:"ent_name,omitempty" `

    /*
        所属管理机构     */
    DirectManage  *string `json:"direct_manage,omitempty" `

    /*
        是否法人     */
    LegalOrgFlag  *string `json:"legal_org_flag,omitempty" `

    /*
        机构代码     */
    OrgCode  *string `json:"org_code,omitempty" `

    /*
        注册地编码     */
    RegRegionCode  *string `json:"reg_region_code,omitempty" `

    /*
        注册地明细     */
    RegRegionDetail  *string `json:"reg_region_detail,omitempty" `

    /*
        所在地编码     */
    DictRegionCode  *string `json:"dict_region_code,omitempty" `

    /*
        所在地明细     */
    DictRegionDetail  *string `json:"dict_region_detail,omitempty" `

    /*
        状态1.使用中0.已废除     */
    Status  *string `json:"status,omitempty" `

    /*
        企业类型[1,生产企业 2批发企业 3医疗机构 4药店 5物流]     */
    UserRoleTypeStr  *string `json:"user_role_type_str,omitempty" `

    /*
        企业类型编码     */
    UserRoleType  *string `json:"user_role_type,omitempty" `

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

func (s *AlibabaAlihealthDrugtraceTopYljgQueryGetbyrefentidPUserEntInfoDto) SetEntId(v string) *AlibabaAlihealthDrugtraceTopYljgQueryGetbyrefentidPUserEntInfoDto {
    s.EntId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryGetbyrefentidPUserEntInfoDto) SetRefEntId(v string) *AlibabaAlihealthDrugtraceTopYljgQueryGetbyrefentidPUserEntInfoDto {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryGetbyrefentidPUserEntInfoDto) SetIsNetwork(v string) *AlibabaAlihealthDrugtraceTopYljgQueryGetbyrefentidPUserEntInfoDto {
    s.IsNetwork = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryGetbyrefentidPUserEntInfoDto) SetEntName(v string) *AlibabaAlihealthDrugtraceTopYljgQueryGetbyrefentidPUserEntInfoDto {
    s.EntName = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryGetbyrefentidPUserEntInfoDto) SetDirectManage(v string) *AlibabaAlihealthDrugtraceTopYljgQueryGetbyrefentidPUserEntInfoDto {
    s.DirectManage = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryGetbyrefentidPUserEntInfoDto) SetLegalOrgFlag(v string) *AlibabaAlihealthDrugtraceTopYljgQueryGetbyrefentidPUserEntInfoDto {
    s.LegalOrgFlag = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryGetbyrefentidPUserEntInfoDto) SetOrgCode(v string) *AlibabaAlihealthDrugtraceTopYljgQueryGetbyrefentidPUserEntInfoDto {
    s.OrgCode = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryGetbyrefentidPUserEntInfoDto) SetRegRegionCode(v string) *AlibabaAlihealthDrugtraceTopYljgQueryGetbyrefentidPUserEntInfoDto {
    s.RegRegionCode = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryGetbyrefentidPUserEntInfoDto) SetRegRegionDetail(v string) *AlibabaAlihealthDrugtraceTopYljgQueryGetbyrefentidPUserEntInfoDto {
    s.RegRegionDetail = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryGetbyrefentidPUserEntInfoDto) SetDictRegionCode(v string) *AlibabaAlihealthDrugtraceTopYljgQueryGetbyrefentidPUserEntInfoDto {
    s.DictRegionCode = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryGetbyrefentidPUserEntInfoDto) SetDictRegionDetail(v string) *AlibabaAlihealthDrugtraceTopYljgQueryGetbyrefentidPUserEntInfoDto {
    s.DictRegionDetail = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryGetbyrefentidPUserEntInfoDto) SetStatus(v string) *AlibabaAlihealthDrugtraceTopYljgQueryGetbyrefentidPUserEntInfoDto {
    s.Status = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryGetbyrefentidPUserEntInfoDto) SetUserRoleTypeStr(v string) *AlibabaAlihealthDrugtraceTopYljgQueryGetbyrefentidPUserEntInfoDto {
    s.UserRoleTypeStr = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryGetbyrefentidPUserEntInfoDto) SetUserRoleType(v string) *AlibabaAlihealthDrugtraceTopYljgQueryGetbyrefentidPUserEntInfoDto {
    s.UserRoleType = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryGetbyrefentidPUserEntInfoDto) SetEntOrgType(v string) *AlibabaAlihealthDrugtraceTopYljgQueryGetbyrefentidPUserEntInfoDto {
    s.EntOrgType = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryGetbyrefentidPUserEntInfoDto) SetProvName(v string) *AlibabaAlihealthDrugtraceTopYljgQueryGetbyrefentidPUserEntInfoDto {
    s.ProvName = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryGetbyrefentidPUserEntInfoDto) SetAreaName(v string) *AlibabaAlihealthDrugtraceTopYljgQueryGetbyrefentidPUserEntInfoDto {
    s.AreaName = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryGetbyrefentidPUserEntInfoDto) SetCityName(v string) *AlibabaAlihealthDrugtraceTopYljgQueryGetbyrefentidPUserEntInfoDto {
    s.CityName = &v
    return s
}
