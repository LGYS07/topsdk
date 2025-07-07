package domain


type AlibabaAlihealthDrugtraceTopLsydQueryRelationBaseInfoDto struct {
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

func (s *AlibabaAlihealthDrugtraceTopLsydQueryRelationBaseInfoDto) SetPhysicInfo(v string) *AlibabaAlihealthDrugtraceTopLsydQueryRelationBaseInfoDto {
    s.PhysicInfo = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryRelationBaseInfoDto) SetRefEntId(v string) *AlibabaAlihealthDrugtraceTopLsydQueryRelationBaseInfoDto {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryRelationBaseInfoDto) SetPkgSpec(v string) *AlibabaAlihealthDrugtraceTopLsydQueryRelationBaseInfoDto {
    s.PkgSpec = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryRelationBaseInfoDto) SetPrepnSpec(v string) *AlibabaAlihealthDrugtraceTopLsydQueryRelationBaseInfoDto {
    s.PrepnSpec = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryRelationBaseInfoDto) SetPrepnType(v string) *AlibabaAlihealthDrugtraceTopLsydQueryRelationBaseInfoDto {
    s.PrepnType = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryRelationBaseInfoDto) SetPhysicName(v string) *AlibabaAlihealthDrugtraceTopLsydQueryRelationBaseInfoDto {
    s.PhysicName = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryRelationBaseInfoDto) SetPkgRatio(v string) *AlibabaAlihealthDrugtraceTopLsydQueryRelationBaseInfoDto {
    s.PkgRatio = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryRelationBaseInfoDto) SetExprieDate(v string) *AlibabaAlihealthDrugtraceTopLsydQueryRelationBaseInfoDto {
    s.ExprieDate = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryRelationBaseInfoDto) SetProduceBatchNo(v string) *AlibabaAlihealthDrugtraceTopLsydQueryRelationBaseInfoDto {
    s.ProduceBatchNo = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryRelationBaseInfoDto) SetProduceDate(v string) *AlibabaAlihealthDrugtraceTopLsydQueryRelationBaseInfoDto {
    s.ProduceDate = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryRelationBaseInfoDto) SetSubTypeNo(v string) *AlibabaAlihealthDrugtraceTopLsydQueryRelationBaseInfoDto {
    s.SubTypeNo = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryRelationBaseInfoDto) SetProductCode(v string) *AlibabaAlihealthDrugtraceTopLsydQueryRelationBaseInfoDto {
    s.ProductCode = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryRelationBaseInfoDto) SetProdId(v string) *AlibabaAlihealthDrugtraceTopLsydQueryRelationBaseInfoDto {
    s.ProdId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryRelationBaseInfoDto) SetApproveNo(v string) *AlibabaAlihealthDrugtraceTopLsydQueryRelationBaseInfoDto {
    s.ApproveNo = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryRelationBaseInfoDto) SetPhysicType(v string) *AlibabaAlihealthDrugtraceTopLsydQueryRelationBaseInfoDto {
    s.PhysicType = &v
    return s
}
