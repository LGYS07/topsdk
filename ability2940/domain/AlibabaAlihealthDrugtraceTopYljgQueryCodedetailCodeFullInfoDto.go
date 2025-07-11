package domain


type AlibabaAlihealthDrugtraceTopYljgQueryCodedetailCodeFullInfoDto struct {
    /*
        码对象     */
    CodeStatusTypeDTO  *AlibabaAlihealthDrugtraceTopYljgQueryCodedetailCodeStatusTypeDto `json:"code_status_type_d_t_o,omitempty" `

    /*
        追溯码     */
    Code  *string `json:"code,omitempty" `

    /*
        企业信息对象     */
    PUserEntDTO  *AlibabaAlihealthDrugtraceTopYljgQueryCodedetailPUserEntDto `json:"p_user_ent_d_t_o,omitempty" `

    /*
        码等级【1代表最小码 如：申请的包装比例是1:5:10, 对应的码等级就是3、2、1, 代表大码、中码、小码】     */
    PackageLevel  *string `json:"package_level,omitempty" `

    /*
        药品基本信息对象     */
    DrugEntBaseDTO  *AlibabaAlihealthDrugtraceTopYljgQueryCodedetailDrugEntBaseDto `json:"drug_ent_base_d_t_o,omitempty" `

    /*
        码生产信息对象     */
    CodeProduceInfoDTO  *AlibabaAlihealthDrugtraceTopYljgQueryCodedetailCodeProduceInfoDto `json:"code_produce_info_d_t_o,omitempty" `

}

func (s *AlibabaAlihealthDrugtraceTopYljgQueryCodedetailCodeFullInfoDto) SetCodeStatusTypeDTO(v AlibabaAlihealthDrugtraceTopYljgQueryCodedetailCodeStatusTypeDto) *AlibabaAlihealthDrugtraceTopYljgQueryCodedetailCodeFullInfoDto {
    s.CodeStatusTypeDTO = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryCodedetailCodeFullInfoDto) SetCode(v string) *AlibabaAlihealthDrugtraceTopYljgQueryCodedetailCodeFullInfoDto {
    s.Code = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryCodedetailCodeFullInfoDto) SetPUserEntDTO(v AlibabaAlihealthDrugtraceTopYljgQueryCodedetailPUserEntDto) *AlibabaAlihealthDrugtraceTopYljgQueryCodedetailCodeFullInfoDto {
    s.PUserEntDTO = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryCodedetailCodeFullInfoDto) SetPackageLevel(v string) *AlibabaAlihealthDrugtraceTopYljgQueryCodedetailCodeFullInfoDto {
    s.PackageLevel = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryCodedetailCodeFullInfoDto) SetDrugEntBaseDTO(v AlibabaAlihealthDrugtraceTopYljgQueryCodedetailDrugEntBaseDto) *AlibabaAlihealthDrugtraceTopYljgQueryCodedetailCodeFullInfoDto {
    s.DrugEntBaseDTO = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryCodedetailCodeFullInfoDto) SetCodeProduceInfoDTO(v AlibabaAlihealthDrugtraceTopYljgQueryCodedetailCodeProduceInfoDto) *AlibabaAlihealthDrugtraceTopYljgQueryCodedetailCodeFullInfoDto {
    s.CodeProduceInfoDTO = &v
    return s
}
