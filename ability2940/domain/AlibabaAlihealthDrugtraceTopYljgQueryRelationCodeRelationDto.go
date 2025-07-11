package domain


type AlibabaAlihealthDrugtraceTopYljgQueryRelationCodeRelationDto struct {
    /*
        激活信息     */
    CodeActiveInfoDto  *AlibabaAlihealthDrugtraceTopYljgQueryRelationCodeActiveInfoDto `json:"code_active_info_dto,omitempty" `

    /*
        码关联关系     */
    CodeRelationList  *[]AlibabaAlihealthDrugtraceTopYljgQueryRelationCodeInfo `json:"code_relation_list,omitempty" `

    /*
        是否是最小包装     */
    IsSmallest  *string `json:"is_smallest,omitempty" `

    /*
        药品包装信息     */
    PkgInfoDto  *AlibabaAlihealthDrugtraceTopYljgQueryRelationPkgInfoDto `json:"pkg_info_dto,omitempty" `

    /*
        药品基础信息     */
    BaseInfosDto  *AlibabaAlihealthDrugtraceTopYljgQueryRelationBaseInfosDto `json:"base_infos_dto,omitempty" `

    /*
        生产信息     */
    ProduceInfoList  *[]AlibabaAlihealthDrugtraceTopYljgQueryRelationProduceInfoDto `json:"produce_info_list,omitempty" `

}

func (s *AlibabaAlihealthDrugtraceTopYljgQueryRelationCodeRelationDto) SetCodeActiveInfoDto(v AlibabaAlihealthDrugtraceTopYljgQueryRelationCodeActiveInfoDto) *AlibabaAlihealthDrugtraceTopYljgQueryRelationCodeRelationDto {
    s.CodeActiveInfoDto = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryRelationCodeRelationDto) SetCodeRelationList(v []AlibabaAlihealthDrugtraceTopYljgQueryRelationCodeInfo) *AlibabaAlihealthDrugtraceTopYljgQueryRelationCodeRelationDto {
    s.CodeRelationList = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryRelationCodeRelationDto) SetIsSmallest(v string) *AlibabaAlihealthDrugtraceTopYljgQueryRelationCodeRelationDto {
    s.IsSmallest = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryRelationCodeRelationDto) SetPkgInfoDto(v AlibabaAlihealthDrugtraceTopYljgQueryRelationPkgInfoDto) *AlibabaAlihealthDrugtraceTopYljgQueryRelationCodeRelationDto {
    s.PkgInfoDto = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryRelationCodeRelationDto) SetBaseInfosDto(v AlibabaAlihealthDrugtraceTopYljgQueryRelationBaseInfosDto) *AlibabaAlihealthDrugtraceTopYljgQueryRelationCodeRelationDto {
    s.BaseInfosDto = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryRelationCodeRelationDto) SetProduceInfoList(v []AlibabaAlihealthDrugtraceTopYljgQueryRelationProduceInfoDto) *AlibabaAlihealthDrugtraceTopYljgQueryRelationCodeRelationDto {
    s.ProduceInfoList = &v
    return s
}
