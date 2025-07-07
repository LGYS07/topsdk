package domain


type AlibabaAlihealthDrugtraceTopLsydQueryRelationCodeRelationDto struct {
    /*
        激活信息     */
    CodeActiveInfoDto  *AlibabaAlihealthDrugtraceTopLsydQueryRelationCodeActiveInfoDto `json:"code_active_info_dto,omitempty" `

    /*
        码关联关系     */
    CodeRelationList  *[]AlibabaAlihealthDrugtraceTopLsydQueryRelationCodeInfo `json:"code_relation_list,omitempty" `

    /*
        是否是最小包装     */
    IsSmallest  *string `json:"is_smallest,omitempty" `

    /*
        药品包装信息     */
    PkgInfoDto  *AlibabaAlihealthDrugtraceTopLsydQueryRelationPkgInfoDto `json:"pkg_info_dto,omitempty" `

    /*
        药品基础信息     */
    BaseInfosDto  *AlibabaAlihealthDrugtraceTopLsydQueryRelationBaseInfosDto `json:"base_infos_dto,omitempty" `

    /*
        生产信息     */
    ProduceInfoList  *[]AlibabaAlihealthDrugtraceTopLsydQueryRelationProduceInfoDto `json:"produce_info_list,omitempty" `

}

func (s *AlibabaAlihealthDrugtraceTopLsydQueryRelationCodeRelationDto) SetCodeActiveInfoDto(v AlibabaAlihealthDrugtraceTopLsydQueryRelationCodeActiveInfoDto) *AlibabaAlihealthDrugtraceTopLsydQueryRelationCodeRelationDto {
    s.CodeActiveInfoDto = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryRelationCodeRelationDto) SetCodeRelationList(v []AlibabaAlihealthDrugtraceTopLsydQueryRelationCodeInfo) *AlibabaAlihealthDrugtraceTopLsydQueryRelationCodeRelationDto {
    s.CodeRelationList = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryRelationCodeRelationDto) SetIsSmallest(v string) *AlibabaAlihealthDrugtraceTopLsydQueryRelationCodeRelationDto {
    s.IsSmallest = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryRelationCodeRelationDto) SetPkgInfoDto(v AlibabaAlihealthDrugtraceTopLsydQueryRelationPkgInfoDto) *AlibabaAlihealthDrugtraceTopLsydQueryRelationCodeRelationDto {
    s.PkgInfoDto = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryRelationCodeRelationDto) SetBaseInfosDto(v AlibabaAlihealthDrugtraceTopLsydQueryRelationBaseInfosDto) *AlibabaAlihealthDrugtraceTopLsydQueryRelationCodeRelationDto {
    s.BaseInfosDto = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryRelationCodeRelationDto) SetProduceInfoList(v []AlibabaAlihealthDrugtraceTopLsydQueryRelationProduceInfoDto) *AlibabaAlihealthDrugtraceTopLsydQueryRelationCodeRelationDto {
    s.ProduceInfoList = &v
    return s
}
