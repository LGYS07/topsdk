package domain


type AlibabaAlihealthDrugtraceTopLsydQueryCodedetailCodeFullInfoDto struct {
    /*
        码对象     */
    CodeStatusTypeDTO  *AlibabaAlihealthDrugtraceTopLsydQueryCodedetailCodeStatusTypeDto `json:"code_status_type_d_t_o,omitempty" `

    /*
        追溯码     */
    Code  *string `json:"code,omitempty" `

    /*
        企业信息对象     */
    PUserEntDTO  *AlibabaAlihealthDrugtraceTopLsydQueryCodedetailPUserEntDto `json:"p_user_ent_d_t_o,omitempty" `

    /*
        码等级【1代表最小码 如：申请的包装比例是1:5:10, 对应的码等级就是3、2、1, 代表大码、中码、小码】     */
    PackageLevel  *string `json:"package_level,omitempty" `

    /*
        药品基本信息对象     */
    DrugEntBaseDTO  *AlibabaAlihealthDrugtraceTopLsydQueryCodedetailDrugEntBaseDto `json:"drug_ent_base_d_t_o,omitempty" `

    /*
        码生产信息对象     */
    CodeProduceInfoDTO  *AlibabaAlihealthDrugtraceTopLsydQueryCodedetailCodeProduceInfoDto `json:"code_produce_info_d_t_o,omitempty" `

}

func (s *AlibabaAlihealthDrugtraceTopLsydQueryCodedetailCodeFullInfoDto) SetCodeStatusTypeDTO(v AlibabaAlihealthDrugtraceTopLsydQueryCodedetailCodeStatusTypeDto) *AlibabaAlihealthDrugtraceTopLsydQueryCodedetailCodeFullInfoDto {
    s.CodeStatusTypeDTO = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryCodedetailCodeFullInfoDto) SetCode(v string) *AlibabaAlihealthDrugtraceTopLsydQueryCodedetailCodeFullInfoDto {
    s.Code = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryCodedetailCodeFullInfoDto) SetPUserEntDTO(v AlibabaAlihealthDrugtraceTopLsydQueryCodedetailPUserEntDto) *AlibabaAlihealthDrugtraceTopLsydQueryCodedetailCodeFullInfoDto {
    s.PUserEntDTO = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryCodedetailCodeFullInfoDto) SetPackageLevel(v string) *AlibabaAlihealthDrugtraceTopLsydQueryCodedetailCodeFullInfoDto {
    s.PackageLevel = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryCodedetailCodeFullInfoDto) SetDrugEntBaseDTO(v AlibabaAlihealthDrugtraceTopLsydQueryCodedetailDrugEntBaseDto) *AlibabaAlihealthDrugtraceTopLsydQueryCodedetailCodeFullInfoDto {
    s.DrugEntBaseDTO = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryCodedetailCodeFullInfoDto) SetCodeProduceInfoDTO(v AlibabaAlihealthDrugtraceTopLsydQueryCodedetailCodeProduceInfoDto) *AlibabaAlihealthDrugtraceTopLsydQueryCodedetailCodeFullInfoDto {
    s.CodeProduceInfoDTO = &v
    return s
}
