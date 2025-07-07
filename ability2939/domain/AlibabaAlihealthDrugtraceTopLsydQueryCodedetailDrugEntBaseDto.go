package domain


type AlibabaAlihealthDrugtraceTopLsydQueryCodedetailDrugEntBaseDto struct {
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

func (s *AlibabaAlihealthDrugtraceTopLsydQueryCodedetailDrugEntBaseDto) SetPhysicTypeDesc(v string) *AlibabaAlihealthDrugtraceTopLsydQueryCodedetailDrugEntBaseDto {
    s.PhysicTypeDesc = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryCodedetailDrugEntBaseDto) SetPhysicName(v string) *AlibabaAlihealthDrugtraceTopLsydQueryCodedetailDrugEntBaseDto {
    s.PhysicName = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryCodedetailDrugEntBaseDto) SetExprie(v string) *AlibabaAlihealthDrugtraceTopLsydQueryCodedetailDrugEntBaseDto {
    s.Exprie = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryCodedetailDrugEntBaseDto) SetDrugEntBaseInfoId(v string) *AlibabaAlihealthDrugtraceTopLsydQueryCodedetailDrugEntBaseDto {
    s.DrugEntBaseInfoId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryCodedetailDrugEntBaseDto) SetApprovalLicenceNo(v string) *AlibabaAlihealthDrugtraceTopLsydQueryCodedetailDrugEntBaseDto {
    s.ApprovalLicenceNo = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryCodedetailDrugEntBaseDto) SetPkgSpecCrit(v string) *AlibabaAlihealthDrugtraceTopLsydQueryCodedetailDrugEntBaseDto {
    s.PkgSpecCrit = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryCodedetailDrugEntBaseDto) SetPrepnSpec(v string) *AlibabaAlihealthDrugtraceTopLsydQueryCodedetailDrugEntBaseDto {
    s.PrepnSpec = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryCodedetailDrugEntBaseDto) SetPrepnTypeDesc(v string) *AlibabaAlihealthDrugtraceTopLsydQueryCodedetailDrugEntBaseDto {
    s.PrepnTypeDesc = &v
    return s
}
