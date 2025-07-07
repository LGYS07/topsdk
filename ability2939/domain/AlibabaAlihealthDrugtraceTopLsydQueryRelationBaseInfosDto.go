package domain


type AlibabaAlihealthDrugtraceTopLsydQueryRelationBaseInfosDto struct {
    /*
        药品基础信息     */
    BaseInfoList  *[]AlibabaAlihealthDrugtraceTopLsydQueryRelationBaseInfoDto `json:"base_info_list,omitempty" `

}

func (s *AlibabaAlihealthDrugtraceTopLsydQueryRelationBaseInfosDto) SetBaseInfoList(v []AlibabaAlihealthDrugtraceTopLsydQueryRelationBaseInfoDto) *AlibabaAlihealthDrugtraceTopLsydQueryRelationBaseInfosDto {
    s.BaseInfoList = &v
    return s
}
