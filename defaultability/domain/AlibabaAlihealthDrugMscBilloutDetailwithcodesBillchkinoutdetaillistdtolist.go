package domain


type AlibabaAlihealthDrugMscBilloutDetailwithcodesBillchkinoutdetaillistdtolist struct {
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
        码列表     */
    CodeAndParentList  *[]AlibabaAlihealthDrugMscBilloutDetailwithcodesCodeandparentlist `json:"code_and_parent_list,omitempty" `

    /*
        国药准字     */
    ApproveNo  *string `json:"approve_no,omitempty" `

}

func (s *AlibabaAlihealthDrugMscBilloutDetailwithcodesBillchkinoutdetaillistdtolist) SetExpiredDate(v string) *AlibabaAlihealthDrugMscBilloutDetailwithcodesBillchkinoutdetaillistdtolist {
    s.ExpiredDate = &v
    return s
}
func (s *AlibabaAlihealthDrugMscBilloutDetailwithcodesBillchkinoutdetaillistdtolist) SetProduceEntName(v string) *AlibabaAlihealthDrugMscBilloutDetailwithcodesBillchkinoutdetaillistdtolist {
    s.ProduceEntName = &v
    return s
}
func (s *AlibabaAlihealthDrugMscBilloutDetailwithcodesBillchkinoutdetaillistdtolist) SetProdCode(v string) *AlibabaAlihealthDrugMscBilloutDetailwithcodesBillchkinoutdetaillistdtolist {
    s.ProdCode = &v
    return s
}
func (s *AlibabaAlihealthDrugMscBilloutDetailwithcodesBillchkinoutdetaillistdtolist) SetProductCode(v string) *AlibabaAlihealthDrugMscBilloutDetailwithcodesBillchkinoutdetaillistdtolist {
    s.ProductCode = &v
    return s
}
func (s *AlibabaAlihealthDrugMscBilloutDetailwithcodesBillchkinoutdetaillistdtolist) SetProduceDate(v string) *AlibabaAlihealthDrugMscBilloutDetailwithcodesBillchkinoutdetaillistdtolist {
    s.ProduceDate = &v
    return s
}
func (s *AlibabaAlihealthDrugMscBilloutDetailwithcodesBillchkinoutdetaillistdtolist) SetProductBatchNo(v string) *AlibabaAlihealthDrugMscBilloutDetailwithcodesBillchkinoutdetaillistdtolist {
    s.ProductBatchNo = &v
    return s
}
func (s *AlibabaAlihealthDrugMscBilloutDetailwithcodesBillchkinoutdetaillistdtolist) SetDrugEntBaseInfoId(v string) *AlibabaAlihealthDrugMscBilloutDetailwithcodesBillchkinoutdetaillistdtolist {
    s.DrugEntBaseInfoId = &v
    return s
}
func (s *AlibabaAlihealthDrugMscBilloutDetailwithcodesBillchkinoutdetaillistdtolist) SetPhysicName(v string) *AlibabaAlihealthDrugMscBilloutDetailwithcodesBillchkinoutdetaillistdtolist {
    s.PhysicName = &v
    return s
}
func (s *AlibabaAlihealthDrugMscBilloutDetailwithcodesBillchkinoutdetaillistdtolist) SetPreparationsUnit(v string) *AlibabaAlihealthDrugMscBilloutDetailwithcodesBillchkinoutdetaillistdtolist {
    s.PreparationsUnit = &v
    return s
}
func (s *AlibabaAlihealthDrugMscBilloutDetailwithcodesBillchkinoutdetaillistdtolist) SetTempPkgSpec(v string) *AlibabaAlihealthDrugMscBilloutDetailwithcodesBillchkinoutdetaillistdtolist {
    s.TempPkgSpec = &v
    return s
}
func (s *AlibabaAlihealthDrugMscBilloutDetailwithcodesBillchkinoutdetaillistdtolist) SetMinPreparationsCount(v string) *AlibabaAlihealthDrugMscBilloutDetailwithcodesBillchkinoutdetaillistdtolist {
    s.MinPreparationsCount = &v
    return s
}
func (s *AlibabaAlihealthDrugMscBilloutDetailwithcodesBillchkinoutdetaillistdtolist) SetMinPkgCount(v string) *AlibabaAlihealthDrugMscBilloutDetailwithcodesBillchkinoutdetaillistdtolist {
    s.MinPkgCount = &v
    return s
}
func (s *AlibabaAlihealthDrugMscBilloutDetailwithcodesBillchkinoutdetaillistdtolist) SetPhysicTypeName(v string) *AlibabaAlihealthDrugMscBilloutDetailwithcodesBillchkinoutdetaillistdtolist {
    s.PhysicTypeName = &v
    return s
}
func (s *AlibabaAlihealthDrugMscBilloutDetailwithcodesBillchkinoutdetaillistdtolist) SetPhysicType(v string) *AlibabaAlihealthDrugMscBilloutDetailwithcodesBillchkinoutdetaillistdtolist {
    s.PhysicType = &v
    return s
}
func (s *AlibabaAlihealthDrugMscBilloutDetailwithcodesBillchkinoutdetaillistdtolist) SetCodeAndParentList(v []AlibabaAlihealthDrugMscBilloutDetailwithcodesCodeandparentlist) *AlibabaAlihealthDrugMscBilloutDetailwithcodesBillchkinoutdetaillistdtolist {
    s.CodeAndParentList = &v
    return s
}
func (s *AlibabaAlihealthDrugMscBilloutDetailwithcodesBillchkinoutdetaillistdtolist) SetApproveNo(v string) *AlibabaAlihealthDrugMscBilloutDetailwithcodesBillchkinoutdetaillistdtolist {
    s.ApproveNo = &v
    return s
}
