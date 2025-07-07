package domain


type AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailBillchkinoutdetaillistdtolist struct {
    /*
        有效期至     */
    ExpiredDate  *string `json:"expired_date,omitempty" `

    /*
        生产企业名称     */
    ProduceEntName  *string `json:"produce_ent_name,omitempty" `

    /*
        子类编码     */
    ProdCode  *string `json:"prod_code,omitempty" `

    /*
        子类编码前7位     */
    ProductCode  *string `json:"product_code,omitempty" `

    /*
        生产日期     */
    ProduceDate  *string `json:"produce_date,omitempty" `

    /*
        批次号     */
    ProductBatchNo  *string `json:"product_batch_no,omitempty" `

    /*
        药品id     */
    DrugEntBaseInfoId  *string `json:"drug_ent_base_info_id,omitempty" `

    /*
        药品名称     */
    PhysicName  *string `json:"physic_name,omitempty" `

    /*
        制剂单位     */
    PreparationsUnit  *string `json:"preparations_unit,omitempty" `

    /*
        包装规格     */
    TempPkgSpec  *string `json:"temp_pkg_spec,omitempty" `

    /*
        最小制剂数量     */
    MinPreparationsCount  *string `json:"min_preparations_count,omitempty" `

    /*
        最小包装数量     */
    MinPkgCount  *string `json:"min_pkg_count,omitempty" `

    /*
        药品类型名称     */
    PhysicTypeName  *string `json:"physic_type_name,omitempty" `

    /*
        药品类型编码     */
    PhysicType  *string `json:"physic_type,omitempty" `

    /*
        码列表（预留属性，暂无返回）     */
    CodeAndParentList  *[]AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailCodeandparentlist `json:"code_and_parent_list,omitempty" `

    /*
        国药准字     */
    ApproveNo  *string `json:"approve_no,omitempty" `

}

func (s *AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailBillchkinoutdetaillistdtolist) SetExpiredDate(v string) *AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailBillchkinoutdetaillistdtolist {
    s.ExpiredDate = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailBillchkinoutdetaillistdtolist) SetProduceEntName(v string) *AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailBillchkinoutdetaillistdtolist {
    s.ProduceEntName = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailBillchkinoutdetaillistdtolist) SetProdCode(v string) *AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailBillchkinoutdetaillistdtolist {
    s.ProdCode = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailBillchkinoutdetaillistdtolist) SetProductCode(v string) *AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailBillchkinoutdetaillistdtolist {
    s.ProductCode = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailBillchkinoutdetaillistdtolist) SetProduceDate(v string) *AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailBillchkinoutdetaillistdtolist {
    s.ProduceDate = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailBillchkinoutdetaillistdtolist) SetProductBatchNo(v string) *AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailBillchkinoutdetaillistdtolist {
    s.ProductBatchNo = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailBillchkinoutdetaillistdtolist) SetDrugEntBaseInfoId(v string) *AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailBillchkinoutdetaillistdtolist {
    s.DrugEntBaseInfoId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailBillchkinoutdetaillistdtolist) SetPhysicName(v string) *AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailBillchkinoutdetaillistdtolist {
    s.PhysicName = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailBillchkinoutdetaillistdtolist) SetPreparationsUnit(v string) *AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailBillchkinoutdetaillistdtolist {
    s.PreparationsUnit = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailBillchkinoutdetaillistdtolist) SetTempPkgSpec(v string) *AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailBillchkinoutdetaillistdtolist {
    s.TempPkgSpec = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailBillchkinoutdetaillistdtolist) SetMinPreparationsCount(v string) *AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailBillchkinoutdetaillistdtolist {
    s.MinPreparationsCount = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailBillchkinoutdetaillistdtolist) SetMinPkgCount(v string) *AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailBillchkinoutdetaillistdtolist {
    s.MinPkgCount = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailBillchkinoutdetaillistdtolist) SetPhysicTypeName(v string) *AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailBillchkinoutdetaillistdtolist {
    s.PhysicTypeName = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailBillchkinoutdetaillistdtolist) SetPhysicType(v string) *AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailBillchkinoutdetaillistdtolist {
    s.PhysicType = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailBillchkinoutdetaillistdtolist) SetCodeAndParentList(v []AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailCodeandparentlist) *AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailBillchkinoutdetaillistdtolist {
    s.CodeAndParentList = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailBillchkinoutdetaillistdtolist) SetApproveNo(v string) *AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailBillchkinoutdetaillistdtolist {
    s.ApproveNo = &v
    return s
}
