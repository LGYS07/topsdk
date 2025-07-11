package domain


type AlibabaAlihealthDrugtraceTopYljgQueryRelationBaseInfoDto struct {
    /*
        药品信息     */
    PhysicInfo  *string `json:"physic_info,omitempty" `

    /*
        企业id     */
    RefEntId  *string `json:"ref_ent_id,omitempty" `

    /*
        药品包装规格     */
    PkgSpec  *string `json:"pkg_spec,omitempty" `

    /*
        药品制剂规格     */
    PrepnSpec  *string `json:"prepn_spec,omitempty" `

    /*
        药品制剂类型     */
    PrepnType  *string `json:"prepn_type,omitempty" `

    /*
        药品通用名称     */
    PhysicName  *string `json:"physic_name,omitempty" `

    /*
        药品包装比例     */
    PkgRatio  *string `json:"pkg_ratio,omitempty" `

    /*
        药品有效期至     */
    ExprieDate  *string `json:"exprie_date,omitempty" `

    /*
        药品生产批次号     */
    ProduceBatchNo  *string `json:"produce_batch_no,omitempty" `

    /*
        药品生产日期     */
    ProduceDate  *string `json:"produce_date,omitempty" `

    /*
        药品自类编码     */
    SubTypeNo  *string `json:"sub_type_no,omitempty" `

    /*
        药品编号     */
    ProductCode  *string `json:"product_code,omitempty" `

    /*
        药品ID     */
    ProdId  *string `json:"prod_id,omitempty" `

    /*
        批准文号     */
    ApproveNo  *string `json:"approve_no,omitempty" `

    /*
        药品类型     */
    PhysicType  *string `json:"physic_type,omitempty" `

}

func (s *AlibabaAlihealthDrugtraceTopYljgQueryRelationBaseInfoDto) SetPhysicInfo(v string) *AlibabaAlihealthDrugtraceTopYljgQueryRelationBaseInfoDto {
    s.PhysicInfo = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryRelationBaseInfoDto) SetRefEntId(v string) *AlibabaAlihealthDrugtraceTopYljgQueryRelationBaseInfoDto {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryRelationBaseInfoDto) SetPkgSpec(v string) *AlibabaAlihealthDrugtraceTopYljgQueryRelationBaseInfoDto {
    s.PkgSpec = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryRelationBaseInfoDto) SetPrepnSpec(v string) *AlibabaAlihealthDrugtraceTopYljgQueryRelationBaseInfoDto {
    s.PrepnSpec = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryRelationBaseInfoDto) SetPrepnType(v string) *AlibabaAlihealthDrugtraceTopYljgQueryRelationBaseInfoDto {
    s.PrepnType = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryRelationBaseInfoDto) SetPhysicName(v string) *AlibabaAlihealthDrugtraceTopYljgQueryRelationBaseInfoDto {
    s.PhysicName = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryRelationBaseInfoDto) SetPkgRatio(v string) *AlibabaAlihealthDrugtraceTopYljgQueryRelationBaseInfoDto {
    s.PkgRatio = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryRelationBaseInfoDto) SetExprieDate(v string) *AlibabaAlihealthDrugtraceTopYljgQueryRelationBaseInfoDto {
    s.ExprieDate = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryRelationBaseInfoDto) SetProduceBatchNo(v string) *AlibabaAlihealthDrugtraceTopYljgQueryRelationBaseInfoDto {
    s.ProduceBatchNo = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryRelationBaseInfoDto) SetProduceDate(v string) *AlibabaAlihealthDrugtraceTopYljgQueryRelationBaseInfoDto {
    s.ProduceDate = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryRelationBaseInfoDto) SetSubTypeNo(v string) *AlibabaAlihealthDrugtraceTopYljgQueryRelationBaseInfoDto {
    s.SubTypeNo = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryRelationBaseInfoDto) SetProductCode(v string) *AlibabaAlihealthDrugtraceTopYljgQueryRelationBaseInfoDto {
    s.ProductCode = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryRelationBaseInfoDto) SetProdId(v string) *AlibabaAlihealthDrugtraceTopYljgQueryRelationBaseInfoDto {
    s.ProdId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryRelationBaseInfoDto) SetApproveNo(v string) *AlibabaAlihealthDrugtraceTopYljgQueryRelationBaseInfoDto {
    s.ApproveNo = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryRelationBaseInfoDto) SetPhysicType(v string) *AlibabaAlihealthDrugtraceTopYljgQueryRelationBaseInfoDto {
    s.PhysicType = &v
    return s
}
