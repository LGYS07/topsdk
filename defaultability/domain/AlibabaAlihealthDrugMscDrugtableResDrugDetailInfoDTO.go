package domain


type AlibabaAlihealthDrugMscDrugtableResDrugDetailInfoDTO struct {
    /*
        药品类型(详见码表) 1：特殊药品原料药，2：特殊药品制剂，3：普通药品，9：未分类     */
    PhysicType  *string `json:"physic_type,omitempty" `

    /*
        药品通用名称     */
    PhysicName  *string `json:"physic_name,omitempty" `

    /*
        药品子类编码     */
    ProdCode  *string `json:"prod_code,omitempty" `

    /*
        药品详细类型     */
    PhysicDetailType  *string `json:"physic_detail_type,omitempty" `

    /*
        企业id     */
    ProduceRefEntId  *string `json:"produce_ref_ent_id,omitempty" `

    /*
        商品名称     */
    ProdName  *string `json:"prod_name,omitempty" `

    /*
        修改日期     */
    ModDate  *string `json:"mod_date,omitempty" `

    /*
        生产厂企业名称     */
    ProduceEntName  *string `json:"produce_ent_name,omitempty" `

    /*
        制剂单位类型(详见码表)  赋码最小包装内使用单元单位     */
    PrepnUnit  *string `json:"prepn_unit,omitempty" `

    /*
        制剂规格     */
    PrepnSpec  *string `json:"prepn_spec,omitempty" `

    /*
        批准文号     */
    ApproveNo  *string `json:"approve_no,omitempty" `

    /*
        包装单位     */
    PkgUnit  *string `json:"pkg_unit,omitempty" `

    /*
        药品id     */
    DrugEntBaseInfoId  *string `json:"drug_ent_base_info_id,omitempty" `

    /*
        制剂单位描述     */
    PrepnUnitDesc  *string `json:"prepn_unit_desc,omitempty" `

    /*
        包装规格     */
    PkgSpec  *string `json:"pkg_spec,omitempty" `

    /*
        包装单位描述     */
    PkgUnitDesc  *string `json:"pkg_unit_desc,omitempty" `

    /*
        药品信息     */
    PhysicInfo  *string `json:"physic_info,omitempty" `

}

func (s *AlibabaAlihealthDrugMscDrugtableResDrugDetailInfoDTO) SetPhysicType(v string) *AlibabaAlihealthDrugMscDrugtableResDrugDetailInfoDTO {
    s.PhysicType = &v
    return s
}
func (s *AlibabaAlihealthDrugMscDrugtableResDrugDetailInfoDTO) SetPhysicName(v string) *AlibabaAlihealthDrugMscDrugtableResDrugDetailInfoDTO {
    s.PhysicName = &v
    return s
}
func (s *AlibabaAlihealthDrugMscDrugtableResDrugDetailInfoDTO) SetProdCode(v string) *AlibabaAlihealthDrugMscDrugtableResDrugDetailInfoDTO {
    s.ProdCode = &v
    return s
}
func (s *AlibabaAlihealthDrugMscDrugtableResDrugDetailInfoDTO) SetPhysicDetailType(v string) *AlibabaAlihealthDrugMscDrugtableResDrugDetailInfoDTO {
    s.PhysicDetailType = &v
    return s
}
func (s *AlibabaAlihealthDrugMscDrugtableResDrugDetailInfoDTO) SetProduceRefEntId(v string) *AlibabaAlihealthDrugMscDrugtableResDrugDetailInfoDTO {
    s.ProduceRefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugMscDrugtableResDrugDetailInfoDTO) SetProdName(v string) *AlibabaAlihealthDrugMscDrugtableResDrugDetailInfoDTO {
    s.ProdName = &v
    return s
}
func (s *AlibabaAlihealthDrugMscDrugtableResDrugDetailInfoDTO) SetModDate(v string) *AlibabaAlihealthDrugMscDrugtableResDrugDetailInfoDTO {
    s.ModDate = &v
    return s
}
func (s *AlibabaAlihealthDrugMscDrugtableResDrugDetailInfoDTO) SetProduceEntName(v string) *AlibabaAlihealthDrugMscDrugtableResDrugDetailInfoDTO {
    s.ProduceEntName = &v
    return s
}
func (s *AlibabaAlihealthDrugMscDrugtableResDrugDetailInfoDTO) SetPrepnUnit(v string) *AlibabaAlihealthDrugMscDrugtableResDrugDetailInfoDTO {
    s.PrepnUnit = &v
    return s
}
func (s *AlibabaAlihealthDrugMscDrugtableResDrugDetailInfoDTO) SetPrepnSpec(v string) *AlibabaAlihealthDrugMscDrugtableResDrugDetailInfoDTO {
    s.PrepnSpec = &v
    return s
}
func (s *AlibabaAlihealthDrugMscDrugtableResDrugDetailInfoDTO) SetApproveNo(v string) *AlibabaAlihealthDrugMscDrugtableResDrugDetailInfoDTO {
    s.ApproveNo = &v
    return s
}
func (s *AlibabaAlihealthDrugMscDrugtableResDrugDetailInfoDTO) SetPkgUnit(v string) *AlibabaAlihealthDrugMscDrugtableResDrugDetailInfoDTO {
    s.PkgUnit = &v
    return s
}
func (s *AlibabaAlihealthDrugMscDrugtableResDrugDetailInfoDTO) SetDrugEntBaseInfoId(v string) *AlibabaAlihealthDrugMscDrugtableResDrugDetailInfoDTO {
    s.DrugEntBaseInfoId = &v
    return s
}
func (s *AlibabaAlihealthDrugMscDrugtableResDrugDetailInfoDTO) SetPrepnUnitDesc(v string) *AlibabaAlihealthDrugMscDrugtableResDrugDetailInfoDTO {
    s.PrepnUnitDesc = &v
    return s
}
func (s *AlibabaAlihealthDrugMscDrugtableResDrugDetailInfoDTO) SetPkgSpec(v string) *AlibabaAlihealthDrugMscDrugtableResDrugDetailInfoDTO {
    s.PkgSpec = &v
    return s
}
func (s *AlibabaAlihealthDrugMscDrugtableResDrugDetailInfoDTO) SetPkgUnitDesc(v string) *AlibabaAlihealthDrugMscDrugtableResDrugDetailInfoDTO {
    s.PkgUnitDesc = &v
    return s
}
func (s *AlibabaAlihealthDrugMscDrugtableResDrugDetailInfoDTO) SetPhysicInfo(v string) *AlibabaAlihealthDrugMscDrugtableResDrugDetailInfoDTO {
    s.PhysicInfo = &v
    return s
}
