package domain


type AlibabaAlihealthDrugtraceTopYljgDrugtableDrugTableDto struct {
    /*
        制剂类型描述     */
    PrepnTypeDesc  *string `json:"prepn_type_desc,omitempty" `

    /*
        药品类型描述     */
    PhysicTypeDesc  *string `json:"physic_type_desc,omitempty" `

    /*
        药品类型(详见码表) 1：特殊药品原料药，2：特殊药品制剂，3：普通药品，9：未分类     */
    PhysicType  *int64 `json:"physic_type,omitempty" `

    /*
        药品名称     */
    PhysicName  *string `json:"physic_name,omitempty" `

    /*
        药品自类编码     */
    ProdCode  *string `json:"prod_code,omitempty" `

    /*
        药品详细类型     */
    PhysicDetailType  *int64 `json:"physic_detail_type,omitempty" `

    /*
        企业主键     */
    RefEntId  *string `json:"ref_ent_id,omitempty" `

    /*
        商品名称     */
    ProdName  *string `json:"prod_name,omitempty" `

    /*
        修改日期     */
    ModDate  *string `json:"mod_date,omitempty" `

    /*
        企业名称     */
    EntName  *string `json:"ent_name,omitempty" `

    /*
        包装单位描述     */
    PkgUnitDesc  *string `json:"pkg_unit_desc,omitempty" `

    /*
        药品类型详情描述     */
    PhysicDetailTypeDesc  *string `json:"physic_detail_type_desc,omitempty" `

    /*
        制剂单位描述     */
    PrepnUnitDesc  *string `json:"prepn_unit_desc,omitempty" `

    /*
        子列表     */
    SubTypeList  *[]AlibabaAlihealthDrugtraceTopYljgDrugtableSubTypeList `json:"sub_type_list,omitempty" `

}

func (s *AlibabaAlihealthDrugtraceTopYljgDrugtableDrugTableDto) SetPrepnTypeDesc(v string) *AlibabaAlihealthDrugtraceTopYljgDrugtableDrugTableDto {
    s.PrepnTypeDesc = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgDrugtableDrugTableDto) SetPhysicTypeDesc(v string) *AlibabaAlihealthDrugtraceTopYljgDrugtableDrugTableDto {
    s.PhysicTypeDesc = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgDrugtableDrugTableDto) SetPhysicType(v int64) *AlibabaAlihealthDrugtraceTopYljgDrugtableDrugTableDto {
    s.PhysicType = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgDrugtableDrugTableDto) SetPhysicName(v string) *AlibabaAlihealthDrugtraceTopYljgDrugtableDrugTableDto {
    s.PhysicName = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgDrugtableDrugTableDto) SetProdCode(v string) *AlibabaAlihealthDrugtraceTopYljgDrugtableDrugTableDto {
    s.ProdCode = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgDrugtableDrugTableDto) SetPhysicDetailType(v int64) *AlibabaAlihealthDrugtraceTopYljgDrugtableDrugTableDto {
    s.PhysicDetailType = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgDrugtableDrugTableDto) SetRefEntId(v string) *AlibabaAlihealthDrugtraceTopYljgDrugtableDrugTableDto {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgDrugtableDrugTableDto) SetProdName(v string) *AlibabaAlihealthDrugtraceTopYljgDrugtableDrugTableDto {
    s.ProdName = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgDrugtableDrugTableDto) SetModDate(v string) *AlibabaAlihealthDrugtraceTopYljgDrugtableDrugTableDto {
    s.ModDate = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgDrugtableDrugTableDto) SetEntName(v string) *AlibabaAlihealthDrugtraceTopYljgDrugtableDrugTableDto {
    s.EntName = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgDrugtableDrugTableDto) SetPkgUnitDesc(v string) *AlibabaAlihealthDrugtraceTopYljgDrugtableDrugTableDto {
    s.PkgUnitDesc = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgDrugtableDrugTableDto) SetPhysicDetailTypeDesc(v string) *AlibabaAlihealthDrugtraceTopYljgDrugtableDrugTableDto {
    s.PhysicDetailTypeDesc = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgDrugtableDrugTableDto) SetPrepnUnitDesc(v string) *AlibabaAlihealthDrugtraceTopYljgDrugtableDrugTableDto {
    s.PrepnUnitDesc = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgDrugtableDrugTableDto) SetSubTypeList(v []AlibabaAlihealthDrugtraceTopYljgDrugtableSubTypeList) *AlibabaAlihealthDrugtraceTopYljgDrugtableDrugTableDto {
    s.SubTypeList = &v
    return s
}
