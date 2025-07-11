package domain


type AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailBillchkinoutdetaillistdtolist struct {
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
    CodeAndParentList  *[]AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailCodeandparentlist `json:"code_and_parent_list,omitempty" `

    /*
        国药准字     */
    ApproveNo  *string `json:"approve_no,omitempty" `

}

func (s *AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailBillchkinoutdetaillistdtolist) SetExpiredDate(v string) *AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailBillchkinoutdetaillistdtolist {
    s.ExpiredDate = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailBillchkinoutdetaillistdtolist) SetProduceEntName(v string) *AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailBillchkinoutdetaillistdtolist {
    s.ProduceEntName = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailBillchkinoutdetaillistdtolist) SetProdCode(v string) *AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailBillchkinoutdetaillistdtolist {
    s.ProdCode = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailBillchkinoutdetaillistdtolist) SetProductCode(v string) *AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailBillchkinoutdetaillistdtolist {
    s.ProductCode = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailBillchkinoutdetaillistdtolist) SetProduceDate(v string) *AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailBillchkinoutdetaillistdtolist {
    s.ProduceDate = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailBillchkinoutdetaillistdtolist) SetProductBatchNo(v string) *AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailBillchkinoutdetaillistdtolist {
    s.ProductBatchNo = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailBillchkinoutdetaillistdtolist) SetDrugEntBaseInfoId(v string) *AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailBillchkinoutdetaillistdtolist {
    s.DrugEntBaseInfoId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailBillchkinoutdetaillistdtolist) SetPhysicName(v string) *AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailBillchkinoutdetaillistdtolist {
    s.PhysicName = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailBillchkinoutdetaillistdtolist) SetPreparationsUnit(v string) *AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailBillchkinoutdetaillistdtolist {
    s.PreparationsUnit = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailBillchkinoutdetaillistdtolist) SetTempPkgSpec(v string) *AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailBillchkinoutdetaillistdtolist {
    s.TempPkgSpec = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailBillchkinoutdetaillistdtolist) SetMinPreparationsCount(v string) *AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailBillchkinoutdetaillistdtolist {
    s.MinPreparationsCount = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailBillchkinoutdetaillistdtolist) SetMinPkgCount(v string) *AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailBillchkinoutdetaillistdtolist {
    s.MinPkgCount = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailBillchkinoutdetaillistdtolist) SetPhysicTypeName(v string) *AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailBillchkinoutdetaillistdtolist {
    s.PhysicTypeName = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailBillchkinoutdetaillistdtolist) SetPhysicType(v string) *AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailBillchkinoutdetaillistdtolist {
    s.PhysicType = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailBillchkinoutdetaillistdtolist) SetCodeAndParentList(v []AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailCodeandparentlist) *AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailBillchkinoutdetaillistdtolist {
    s.CodeAndParentList = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailBillchkinoutdetaillistdtolist) SetApproveNo(v string) *AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailBillchkinoutdetaillistdtolist {
    s.ApproveNo = &v
    return s
}
