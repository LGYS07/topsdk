package domain


type AlibabaAlihealthDrugMscSearchbillDetailBillchkinoutdetaillistdtolist struct {
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
    CodeAndParentList  *[]AlibabaAlihealthDrugMscSearchbillDetailCodeandparentlist `json:"code_and_parent_list,omitempty" `

    /*
        国药准字     */
    ApproveNo  *string `json:"approve_no,omitempty" `

}

func (s *AlibabaAlihealthDrugMscSearchbillDetailBillchkinoutdetaillistdtolist) SetExpiredDate(v string) *AlibabaAlihealthDrugMscSearchbillDetailBillchkinoutdetaillistdtolist {
    s.ExpiredDate = &v
    return s
}
func (s *AlibabaAlihealthDrugMscSearchbillDetailBillchkinoutdetaillistdtolist) SetProduceEntName(v string) *AlibabaAlihealthDrugMscSearchbillDetailBillchkinoutdetaillistdtolist {
    s.ProduceEntName = &v
    return s
}
func (s *AlibabaAlihealthDrugMscSearchbillDetailBillchkinoutdetaillistdtolist) SetProdCode(v string) *AlibabaAlihealthDrugMscSearchbillDetailBillchkinoutdetaillistdtolist {
    s.ProdCode = &v
    return s
}
func (s *AlibabaAlihealthDrugMscSearchbillDetailBillchkinoutdetaillistdtolist) SetProductCode(v string) *AlibabaAlihealthDrugMscSearchbillDetailBillchkinoutdetaillistdtolist {
    s.ProductCode = &v
    return s
}
func (s *AlibabaAlihealthDrugMscSearchbillDetailBillchkinoutdetaillistdtolist) SetProduceDate(v string) *AlibabaAlihealthDrugMscSearchbillDetailBillchkinoutdetaillistdtolist {
    s.ProduceDate = &v
    return s
}
func (s *AlibabaAlihealthDrugMscSearchbillDetailBillchkinoutdetaillistdtolist) SetProductBatchNo(v string) *AlibabaAlihealthDrugMscSearchbillDetailBillchkinoutdetaillistdtolist {
    s.ProductBatchNo = &v
    return s
}
func (s *AlibabaAlihealthDrugMscSearchbillDetailBillchkinoutdetaillistdtolist) SetDrugEntBaseInfoId(v string) *AlibabaAlihealthDrugMscSearchbillDetailBillchkinoutdetaillistdtolist {
    s.DrugEntBaseInfoId = &v
    return s
}
func (s *AlibabaAlihealthDrugMscSearchbillDetailBillchkinoutdetaillistdtolist) SetPhysicName(v string) *AlibabaAlihealthDrugMscSearchbillDetailBillchkinoutdetaillistdtolist {
    s.PhysicName = &v
    return s
}
func (s *AlibabaAlihealthDrugMscSearchbillDetailBillchkinoutdetaillistdtolist) SetPreparationsUnit(v string) *AlibabaAlihealthDrugMscSearchbillDetailBillchkinoutdetaillistdtolist {
    s.PreparationsUnit = &v
    return s
}
func (s *AlibabaAlihealthDrugMscSearchbillDetailBillchkinoutdetaillistdtolist) SetTempPkgSpec(v string) *AlibabaAlihealthDrugMscSearchbillDetailBillchkinoutdetaillistdtolist {
    s.TempPkgSpec = &v
    return s
}
func (s *AlibabaAlihealthDrugMscSearchbillDetailBillchkinoutdetaillistdtolist) SetMinPreparationsCount(v string) *AlibabaAlihealthDrugMscSearchbillDetailBillchkinoutdetaillistdtolist {
    s.MinPreparationsCount = &v
    return s
}
func (s *AlibabaAlihealthDrugMscSearchbillDetailBillchkinoutdetaillistdtolist) SetMinPkgCount(v string) *AlibabaAlihealthDrugMscSearchbillDetailBillchkinoutdetaillistdtolist {
    s.MinPkgCount = &v
    return s
}
func (s *AlibabaAlihealthDrugMscSearchbillDetailBillchkinoutdetaillistdtolist) SetPhysicTypeName(v string) *AlibabaAlihealthDrugMscSearchbillDetailBillchkinoutdetaillistdtolist {
    s.PhysicTypeName = &v
    return s
}
func (s *AlibabaAlihealthDrugMscSearchbillDetailBillchkinoutdetaillistdtolist) SetPhysicType(v string) *AlibabaAlihealthDrugMscSearchbillDetailBillchkinoutdetaillistdtolist {
    s.PhysicType = &v
    return s
}
func (s *AlibabaAlihealthDrugMscSearchbillDetailBillchkinoutdetaillistdtolist) SetCodeAndParentList(v []AlibabaAlihealthDrugMscSearchbillDetailCodeandparentlist) *AlibabaAlihealthDrugMscSearchbillDetailBillchkinoutdetaillistdtolist {
    s.CodeAndParentList = &v
    return s
}
func (s *AlibabaAlihealthDrugMscSearchbillDetailBillchkinoutdetaillistdtolist) SetApproveNo(v string) *AlibabaAlihealthDrugMscSearchbillDetailBillchkinoutdetaillistdtolist {
    s.ApproveNo = &v
    return s
}
