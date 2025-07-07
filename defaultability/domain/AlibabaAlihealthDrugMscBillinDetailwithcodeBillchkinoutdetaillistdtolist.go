package domain


type AlibabaAlihealthDrugMscBillinDetailwithcodeBillchkinoutdetaillistdtolist struct {
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
    CodeAndParentList  *[]AlibabaAlihealthDrugMscBillinDetailwithcodeCodeandparentlist `json:"code_and_parent_list,omitempty" `

    /*
        国药准字     */
    ApproveNo  *string `json:"approve_no,omitempty" `

}

func (s *AlibabaAlihealthDrugMscBillinDetailwithcodeBillchkinoutdetaillistdtolist) SetExpiredDate(v string) *AlibabaAlihealthDrugMscBillinDetailwithcodeBillchkinoutdetaillistdtolist {
    s.ExpiredDate = &v
    return s
}
func (s *AlibabaAlihealthDrugMscBillinDetailwithcodeBillchkinoutdetaillistdtolist) SetProduceEntName(v string) *AlibabaAlihealthDrugMscBillinDetailwithcodeBillchkinoutdetaillistdtolist {
    s.ProduceEntName = &v
    return s
}
func (s *AlibabaAlihealthDrugMscBillinDetailwithcodeBillchkinoutdetaillistdtolist) SetProdCode(v string) *AlibabaAlihealthDrugMscBillinDetailwithcodeBillchkinoutdetaillistdtolist {
    s.ProdCode = &v
    return s
}
func (s *AlibabaAlihealthDrugMscBillinDetailwithcodeBillchkinoutdetaillistdtolist) SetProductCode(v string) *AlibabaAlihealthDrugMscBillinDetailwithcodeBillchkinoutdetaillistdtolist {
    s.ProductCode = &v
    return s
}
func (s *AlibabaAlihealthDrugMscBillinDetailwithcodeBillchkinoutdetaillistdtolist) SetProduceDate(v string) *AlibabaAlihealthDrugMscBillinDetailwithcodeBillchkinoutdetaillistdtolist {
    s.ProduceDate = &v
    return s
}
func (s *AlibabaAlihealthDrugMscBillinDetailwithcodeBillchkinoutdetaillistdtolist) SetProductBatchNo(v string) *AlibabaAlihealthDrugMscBillinDetailwithcodeBillchkinoutdetaillistdtolist {
    s.ProductBatchNo = &v
    return s
}
func (s *AlibabaAlihealthDrugMscBillinDetailwithcodeBillchkinoutdetaillistdtolist) SetDrugEntBaseInfoId(v string) *AlibabaAlihealthDrugMscBillinDetailwithcodeBillchkinoutdetaillistdtolist {
    s.DrugEntBaseInfoId = &v
    return s
}
func (s *AlibabaAlihealthDrugMscBillinDetailwithcodeBillchkinoutdetaillistdtolist) SetPhysicName(v string) *AlibabaAlihealthDrugMscBillinDetailwithcodeBillchkinoutdetaillistdtolist {
    s.PhysicName = &v
    return s
}
func (s *AlibabaAlihealthDrugMscBillinDetailwithcodeBillchkinoutdetaillistdtolist) SetPreparationsUnit(v string) *AlibabaAlihealthDrugMscBillinDetailwithcodeBillchkinoutdetaillistdtolist {
    s.PreparationsUnit = &v
    return s
}
func (s *AlibabaAlihealthDrugMscBillinDetailwithcodeBillchkinoutdetaillistdtolist) SetTempPkgSpec(v string) *AlibabaAlihealthDrugMscBillinDetailwithcodeBillchkinoutdetaillistdtolist {
    s.TempPkgSpec = &v
    return s
}
func (s *AlibabaAlihealthDrugMscBillinDetailwithcodeBillchkinoutdetaillistdtolist) SetMinPreparationsCount(v string) *AlibabaAlihealthDrugMscBillinDetailwithcodeBillchkinoutdetaillistdtolist {
    s.MinPreparationsCount = &v
    return s
}
func (s *AlibabaAlihealthDrugMscBillinDetailwithcodeBillchkinoutdetaillistdtolist) SetMinPkgCount(v string) *AlibabaAlihealthDrugMscBillinDetailwithcodeBillchkinoutdetaillistdtolist {
    s.MinPkgCount = &v
    return s
}
func (s *AlibabaAlihealthDrugMscBillinDetailwithcodeBillchkinoutdetaillistdtolist) SetPhysicTypeName(v string) *AlibabaAlihealthDrugMscBillinDetailwithcodeBillchkinoutdetaillistdtolist {
    s.PhysicTypeName = &v
    return s
}
func (s *AlibabaAlihealthDrugMscBillinDetailwithcodeBillchkinoutdetaillistdtolist) SetPhysicType(v string) *AlibabaAlihealthDrugMscBillinDetailwithcodeBillchkinoutdetaillistdtolist {
    s.PhysicType = &v
    return s
}
func (s *AlibabaAlihealthDrugMscBillinDetailwithcodeBillchkinoutdetaillistdtolist) SetCodeAndParentList(v []AlibabaAlihealthDrugMscBillinDetailwithcodeCodeandparentlist) *AlibabaAlihealthDrugMscBillinDetailwithcodeBillchkinoutdetaillistdtolist {
    s.CodeAndParentList = &v
    return s
}
func (s *AlibabaAlihealthDrugMscBillinDetailwithcodeBillchkinoutdetaillistdtolist) SetApproveNo(v string) *AlibabaAlihealthDrugMscBillinDetailwithcodeBillchkinoutdetaillistdtolist {
    s.ApproveNo = &v
    return s
}
