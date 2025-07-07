package domain


type AlibabaAlihealthDrugtraceTopLsydListupoutDetailDrugInfosDto struct {
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
    CodeInfoListDtoList  *[]AlibabaAlihealthDrugtraceTopLsydListupoutDetailCodeInfoListDto `json:"code_info_list_dto_list,omitempty" `

    /*
        包装单位描述     */
    PkgUnitDesc  *string `json:"pkg_unit_desc,omitempty" `

}

func (s *AlibabaAlihealthDrugtraceTopLsydListupoutDetailDrugInfosDto) SetProduceDate(v string) *AlibabaAlihealthDrugtraceTopLsydListupoutDetailDrugInfosDto {
    s.ProduceDate = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydListupoutDetailDrugInfosDto) SetProductEntName(v string) *AlibabaAlihealthDrugtraceTopLsydListupoutDetailDrugInfosDto {
    s.ProductEntName = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydListupoutDetailDrugInfosDto) SetPackageSpec(v string) *AlibabaAlihealthDrugtraceTopLsydListupoutDetailDrugInfosDto {
    s.PackageSpec = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydListupoutDetailDrugInfosDto) SetProdName(v string) *AlibabaAlihealthDrugtraceTopLsydListupoutDetailDrugInfosDto {
    s.ProdName = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydListupoutDetailDrugInfosDto) SetPhysicName(v string) *AlibabaAlihealthDrugtraceTopLsydListupoutDetailDrugInfosDto {
    s.PhysicName = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydListupoutDetailDrugInfosDto) SetPrepnSpec(v string) *AlibabaAlihealthDrugtraceTopLsydListupoutDetailDrugInfosDto {
    s.PrepnSpec = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydListupoutDetailDrugInfosDto) SetPrepnUnit(v string) *AlibabaAlihealthDrugtraceTopLsydListupoutDetailDrugInfosDto {
    s.PrepnUnit = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydListupoutDetailDrugInfosDto) SetProduceBatchNo(v string) *AlibabaAlihealthDrugtraceTopLsydListupoutDetailDrugInfosDto {
    s.ProduceBatchNo = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydListupoutDetailDrugInfosDto) SetProdSeqNo(v string) *AlibabaAlihealthDrugtraceTopLsydListupoutDetailDrugInfosDto {
    s.ProdSeqNo = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydListupoutDetailDrugInfosDto) SetDrugEntBaseInfoId(v string) *AlibabaAlihealthDrugtraceTopLsydListupoutDetailDrugInfosDto {
    s.DrugEntBaseInfoId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydListupoutDetailDrugInfosDto) SetValidEndDate(v string) *AlibabaAlihealthDrugtraceTopLsydListupoutDetailDrugInfosDto {
    s.ValidEndDate = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydListupoutDetailDrugInfosDto) SetLeastPkgAmount(v string) *AlibabaAlihealthDrugtraceTopLsydListupoutDetailDrugInfosDto {
    s.LeastPkgAmount = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydListupoutDetailDrugInfosDto) SetLeastPrepnAmount(v string) *AlibabaAlihealthDrugtraceTopLsydListupoutDetailDrugInfosDto {
    s.LeastPrepnAmount = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydListupoutDetailDrugInfosDto) SetApprovalNo(v string) *AlibabaAlihealthDrugtraceTopLsydListupoutDetailDrugInfosDto {
    s.ApprovalNo = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydListupoutDetailDrugInfosDto) SetPhysicType(v string) *AlibabaAlihealthDrugtraceTopLsydListupoutDetailDrugInfosDto {
    s.PhysicType = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydListupoutDetailDrugInfosDto) SetPhysicTypeName(v string) *AlibabaAlihealthDrugtraceTopLsydListupoutDetailDrugInfosDto {
    s.PhysicTypeName = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydListupoutDetailDrugInfosDto) SetPreparationsUnit(v string) *AlibabaAlihealthDrugtraceTopLsydListupoutDetailDrugInfosDto {
    s.PreparationsUnit = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydListupoutDetailDrugInfosDto) SetPrepnTypeDesc(v string) *AlibabaAlihealthDrugtraceTopLsydListupoutDetailDrugInfosDto {
    s.PrepnTypeDesc = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydListupoutDetailDrugInfosDto) SetCodeInfoListDtoList(v []AlibabaAlihealthDrugtraceTopLsydListupoutDetailCodeInfoListDto) *AlibabaAlihealthDrugtraceTopLsydListupoutDetailDrugInfosDto {
    s.CodeInfoListDtoList = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydListupoutDetailDrugInfosDto) SetPkgUnitDesc(v string) *AlibabaAlihealthDrugtraceTopLsydListupoutDetailDrugInfosDto {
    s.PkgUnitDesc = &v
    return s
}
