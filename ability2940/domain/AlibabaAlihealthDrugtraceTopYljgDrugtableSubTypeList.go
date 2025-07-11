package domain


type AlibabaAlihealthDrugtraceTopYljgDrugtableSubTypeList struct {
    /*
        制剂单位     */
    PrepnUnit  *string `json:"prepn_unit,omitempty" `

    /*
        包装规格     */
    PackageSpec  *string `json:"package_spec,omitempty" `

    /*
        码列表     */
    CodeResList  *[]AlibabaAlihealthDrugtraceTopYljgDrugtableCodeResList `json:"code_res_list,omitempty" `

    /*
        制剂规格     */
    PrepnSpec  *string `json:"prepn_spec,omitempty" `

    /*
        企业药品ID     */
    ProdSeqNo  *string `json:"prod_seq_no,omitempty" `

    /*
        批准文号     */
    ApproveNo  *string `json:"approve_no,omitempty" `

    /*
        药品详情类型     */
    PhysicDetailType  *string `json:"physic_detail_type,omitempty" `

    /*
        包装单位     */
    PackUnit  *string `json:"pack_unit,omitempty" `

    /*
        药品ID     */
    DrugEntBaseInfoId  *string `json:"drug_ent_base_info_id,omitempty" `

    /*
        包装单位     */
    PackUnitName  *string `json:"pack_unit_name,omitempty" `

    /*
        制剂描述     */
    PrepnDesc  *string `json:"prepn_desc,omitempty" `

    /*
        制剂单位描述     */
    PrepnUnitName  *string `json:"prepn_unit_name,omitempty" `

    /*
        子类型     */
    SubTypeNo  *string `json:"sub_type_no,omitempty" `

}

func (s *AlibabaAlihealthDrugtraceTopYljgDrugtableSubTypeList) SetPrepnUnit(v string) *AlibabaAlihealthDrugtraceTopYljgDrugtableSubTypeList {
    s.PrepnUnit = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgDrugtableSubTypeList) SetPackageSpec(v string) *AlibabaAlihealthDrugtraceTopYljgDrugtableSubTypeList {
    s.PackageSpec = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgDrugtableSubTypeList) SetCodeResList(v []AlibabaAlihealthDrugtraceTopYljgDrugtableCodeResList) *AlibabaAlihealthDrugtraceTopYljgDrugtableSubTypeList {
    s.CodeResList = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgDrugtableSubTypeList) SetPrepnSpec(v string) *AlibabaAlihealthDrugtraceTopYljgDrugtableSubTypeList {
    s.PrepnSpec = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgDrugtableSubTypeList) SetProdSeqNo(v string) *AlibabaAlihealthDrugtraceTopYljgDrugtableSubTypeList {
    s.ProdSeqNo = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgDrugtableSubTypeList) SetApproveNo(v string) *AlibabaAlihealthDrugtraceTopYljgDrugtableSubTypeList {
    s.ApproveNo = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgDrugtableSubTypeList) SetPhysicDetailType(v string) *AlibabaAlihealthDrugtraceTopYljgDrugtableSubTypeList {
    s.PhysicDetailType = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgDrugtableSubTypeList) SetPackUnit(v string) *AlibabaAlihealthDrugtraceTopYljgDrugtableSubTypeList {
    s.PackUnit = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgDrugtableSubTypeList) SetDrugEntBaseInfoId(v string) *AlibabaAlihealthDrugtraceTopYljgDrugtableSubTypeList {
    s.DrugEntBaseInfoId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgDrugtableSubTypeList) SetPackUnitName(v string) *AlibabaAlihealthDrugtraceTopYljgDrugtableSubTypeList {
    s.PackUnitName = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgDrugtableSubTypeList) SetPrepnDesc(v string) *AlibabaAlihealthDrugtraceTopYljgDrugtableSubTypeList {
    s.PrepnDesc = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgDrugtableSubTypeList) SetPrepnUnitName(v string) *AlibabaAlihealthDrugtraceTopYljgDrugtableSubTypeList {
    s.PrepnUnitName = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgDrugtableSubTypeList) SetSubTypeNo(v string) *AlibabaAlihealthDrugtraceTopYljgDrugtableSubTypeList {
    s.SubTypeNo = &v
    return s
}
