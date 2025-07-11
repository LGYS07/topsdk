package domain


type AlibabaAlihealthDrugtraceTopYljgQueryCodedetailDrugEntBaseDto struct {
    /*
        药品类型描述     */
    PhysicTypeDesc  *string `json:"physic_type_desc,omitempty" `

    /*
        药品名称     */
    PhysicName  *string `json:"physic_name,omitempty" `

    /*
        有效期     */
    Exprie  *string `json:"exprie,omitempty" `

    /*
        药品id     */
    DrugEntBaseInfoId  *string `json:"drug_ent_base_info_id,omitempty" `

    /*
        批准文号     */
    ApprovalLicenceNo  *string `json:"approval_licence_no,omitempty" `

    /*
        包装规格     */
    PkgSpecCrit  *string `json:"pkg_spec_crit,omitempty" `

    /*
        制剂规格     */
    PrepnSpec  *string `json:"prepn_spec,omitempty" `

    /*
        剂型描述     */
    PrepnTypeDesc  *string `json:"prepn_type_desc,omitempty" `

}

func (s *AlibabaAlihealthDrugtraceTopYljgQueryCodedetailDrugEntBaseDto) SetPhysicTypeDesc(v string) *AlibabaAlihealthDrugtraceTopYljgQueryCodedetailDrugEntBaseDto {
    s.PhysicTypeDesc = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryCodedetailDrugEntBaseDto) SetPhysicName(v string) *AlibabaAlihealthDrugtraceTopYljgQueryCodedetailDrugEntBaseDto {
    s.PhysicName = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryCodedetailDrugEntBaseDto) SetExprie(v string) *AlibabaAlihealthDrugtraceTopYljgQueryCodedetailDrugEntBaseDto {
    s.Exprie = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryCodedetailDrugEntBaseDto) SetDrugEntBaseInfoId(v string) *AlibabaAlihealthDrugtraceTopYljgQueryCodedetailDrugEntBaseDto {
    s.DrugEntBaseInfoId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryCodedetailDrugEntBaseDto) SetApprovalLicenceNo(v string) *AlibabaAlihealthDrugtraceTopYljgQueryCodedetailDrugEntBaseDto {
    s.ApprovalLicenceNo = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryCodedetailDrugEntBaseDto) SetPkgSpecCrit(v string) *AlibabaAlihealthDrugtraceTopYljgQueryCodedetailDrugEntBaseDto {
    s.PkgSpecCrit = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryCodedetailDrugEntBaseDto) SetPrepnSpec(v string) *AlibabaAlihealthDrugtraceTopYljgQueryCodedetailDrugEntBaseDto {
    s.PrepnSpec = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryCodedetailDrugEntBaseDto) SetPrepnTypeDesc(v string) *AlibabaAlihealthDrugtraceTopYljgQueryCodedetailDrugEntBaseDto {
    s.PrepnTypeDesc = &v
    return s
}
