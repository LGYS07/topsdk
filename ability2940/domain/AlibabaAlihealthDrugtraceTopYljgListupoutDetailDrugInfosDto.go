package domain


type AlibabaAlihealthDrugtraceTopYljgListupoutDetailDrugInfosDto struct {
    /*
        生产日期     */
    ProduceDate  *string `json:"produce_date,omitempty" `

    /*
        生产企业名称     */
    ProductEntName  *string `json:"product_ent_name,omitempty" `

    /*
        产品包装规格     */
    PackageSpec  *string `json:"package_spec,omitempty" `

    /*
        药品商品名     */
    ProdName  *string `json:"prod_name,omitempty" `

    /*
        药品通用名     */
    PhysicName  *string `json:"physic_name,omitempty" `

    /*
        制剂规格     */
    PrepnSpec  *string `json:"prepn_spec,omitempty" `

    /*
        制剂单位编码     */
    PrepnUnit  *string `json:"prepn_unit,omitempty" `

    /*
        批次号     */
    ProduceBatchNo  *string `json:"produce_batch_no,omitempty" `

    /*
        药品标识     */
    ProdSeqNo  *string `json:"prod_seq_no,omitempty" `

    /*
        药品标识     */
    DrugEntBaseInfoId  *string `json:"drug_ent_base_info_id,omitempty" `

    /*
        有效期至     */
    ValidEndDate  *string `json:"valid_end_date,omitempty" `

    /*
        按最小包装单位统计数量     */
    LeastPkgAmount  *string `json:"least_pkg_amount,omitempty" `

    /*
        按最小制剂单位统计数量     */
    LeastPrepnAmount  *string `json:"least_prepn_amount,omitempty" `

    /*
        批准文号     */
    ApprovalNo  *string `json:"approval_no,omitempty" `

    /*
        药品类型     */
    PhysicType  *string `json:"physic_type,omitempty" `

    /*
        药品类型描述     */
    PhysicTypeName  *string `json:"physic_type_name,omitempty" `

    /*
        制剂单位     */
    PreparationsUnit  *string `json:"preparations_unit,omitempty" `

    /*
        制剂规格描述     */
    PrepnTypeDesc  *string `json:"prepn_type_desc,omitempty" `

    /*
        码信息     */
    CodeInfoListDtoList  *[]AlibabaAlihealthDrugtraceTopYljgListupoutDetailCodeInfoListDto `json:"code_info_list_dto_list,omitempty" `

    /*
        包装单位描述     */
    PkgUnitDesc  *string `json:"pkg_unit_desc,omitempty" `

}

func (s *AlibabaAlihealthDrugtraceTopYljgListupoutDetailDrugInfosDto) SetProduceDate(v string) *AlibabaAlihealthDrugtraceTopYljgListupoutDetailDrugInfosDto {
    s.ProduceDate = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgListupoutDetailDrugInfosDto) SetProductEntName(v string) *AlibabaAlihealthDrugtraceTopYljgListupoutDetailDrugInfosDto {
    s.ProductEntName = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgListupoutDetailDrugInfosDto) SetPackageSpec(v string) *AlibabaAlihealthDrugtraceTopYljgListupoutDetailDrugInfosDto {
    s.PackageSpec = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgListupoutDetailDrugInfosDto) SetProdName(v string) *AlibabaAlihealthDrugtraceTopYljgListupoutDetailDrugInfosDto {
    s.ProdName = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgListupoutDetailDrugInfosDto) SetPhysicName(v string) *AlibabaAlihealthDrugtraceTopYljgListupoutDetailDrugInfosDto {
    s.PhysicName = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgListupoutDetailDrugInfosDto) SetPrepnSpec(v string) *AlibabaAlihealthDrugtraceTopYljgListupoutDetailDrugInfosDto {
    s.PrepnSpec = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgListupoutDetailDrugInfosDto) SetPrepnUnit(v string) *AlibabaAlihealthDrugtraceTopYljgListupoutDetailDrugInfosDto {
    s.PrepnUnit = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgListupoutDetailDrugInfosDto) SetProduceBatchNo(v string) *AlibabaAlihealthDrugtraceTopYljgListupoutDetailDrugInfosDto {
    s.ProduceBatchNo = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgListupoutDetailDrugInfosDto) SetProdSeqNo(v string) *AlibabaAlihealthDrugtraceTopYljgListupoutDetailDrugInfosDto {
    s.ProdSeqNo = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgListupoutDetailDrugInfosDto) SetDrugEntBaseInfoId(v string) *AlibabaAlihealthDrugtraceTopYljgListupoutDetailDrugInfosDto {
    s.DrugEntBaseInfoId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgListupoutDetailDrugInfosDto) SetValidEndDate(v string) *AlibabaAlihealthDrugtraceTopYljgListupoutDetailDrugInfosDto {
    s.ValidEndDate = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgListupoutDetailDrugInfosDto) SetLeastPkgAmount(v string) *AlibabaAlihealthDrugtraceTopYljgListupoutDetailDrugInfosDto {
    s.LeastPkgAmount = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgListupoutDetailDrugInfosDto) SetLeastPrepnAmount(v string) *AlibabaAlihealthDrugtraceTopYljgListupoutDetailDrugInfosDto {
    s.LeastPrepnAmount = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgListupoutDetailDrugInfosDto) SetApprovalNo(v string) *AlibabaAlihealthDrugtraceTopYljgListupoutDetailDrugInfosDto {
    s.ApprovalNo = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgListupoutDetailDrugInfosDto) SetPhysicType(v string) *AlibabaAlihealthDrugtraceTopYljgListupoutDetailDrugInfosDto {
    s.PhysicType = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgListupoutDetailDrugInfosDto) SetPhysicTypeName(v string) *AlibabaAlihealthDrugtraceTopYljgListupoutDetailDrugInfosDto {
    s.PhysicTypeName = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgListupoutDetailDrugInfosDto) SetPreparationsUnit(v string) *AlibabaAlihealthDrugtraceTopYljgListupoutDetailDrugInfosDto {
    s.PreparationsUnit = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgListupoutDetailDrugInfosDto) SetPrepnTypeDesc(v string) *AlibabaAlihealthDrugtraceTopYljgListupoutDetailDrugInfosDto {
    s.PrepnTypeDesc = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgListupoutDetailDrugInfosDto) SetCodeInfoListDtoList(v []AlibabaAlihealthDrugtraceTopYljgListupoutDetailCodeInfoListDto) *AlibabaAlihealthDrugtraceTopYljgListupoutDetailDrugInfosDto {
    s.CodeInfoListDtoList = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgListupoutDetailDrugInfosDto) SetPkgUnitDesc(v string) *AlibabaAlihealthDrugtraceTopYljgListupoutDetailDrugInfosDto {
    s.PkgUnitDesc = &v
    return s
}
