package domain


type AlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidPUserEntInfoDto struct {
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

func (s *AlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidPUserEntInfoDto) SetEntId(v string) *AlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidPUserEntInfoDto {
    s.EntId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidPUserEntInfoDto) SetRefEntId(v string) *AlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidPUserEntInfoDto {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidPUserEntInfoDto) SetIsNetwork(v string) *AlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidPUserEntInfoDto {
    s.IsNetwork = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidPUserEntInfoDto) SetEntName(v string) *AlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidPUserEntInfoDto {
    s.EntName = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidPUserEntInfoDto) SetDirectManage(v string) *AlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidPUserEntInfoDto {
    s.DirectManage = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidPUserEntInfoDto) SetLegalOrgFlag(v string) *AlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidPUserEntInfoDto {
    s.LegalOrgFlag = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidPUserEntInfoDto) SetOrgCode(v string) *AlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidPUserEntInfoDto {
    s.OrgCode = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidPUserEntInfoDto) SetRegRegionCode(v string) *AlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidPUserEntInfoDto {
    s.RegRegionCode = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidPUserEntInfoDto) SetRegRegionDetail(v string) *AlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidPUserEntInfoDto {
    s.RegRegionDetail = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidPUserEntInfoDto) SetDictRegionCode(v string) *AlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidPUserEntInfoDto {
    s.DictRegionCode = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidPUserEntInfoDto) SetDictRegionDetail(v string) *AlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidPUserEntInfoDto {
    s.DictRegionDetail = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidPUserEntInfoDto) SetStatus(v string) *AlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidPUserEntInfoDto {
    s.Status = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidPUserEntInfoDto) SetUserRoleTypeStr(v string) *AlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidPUserEntInfoDto {
    s.UserRoleTypeStr = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidPUserEntInfoDto) SetUserRoleType(v string) *AlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidPUserEntInfoDto {
    s.UserRoleType = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidPUserEntInfoDto) SetEntOrgType(v string) *AlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidPUserEntInfoDto {
    s.EntOrgType = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidPUserEntInfoDto) SetProvName(v string) *AlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidPUserEntInfoDto {
    s.ProvName = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidPUserEntInfoDto) SetAreaName(v string) *AlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidPUserEntInfoDto {
    s.AreaName = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidPUserEntInfoDto) SetCityName(v string) *AlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidPUserEntInfoDto {
    s.CityName = &v
    return s
}
