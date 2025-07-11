package domain


type AlibabaAlihealthDrugtraceTopYljgQueryRelationBaseInfosDto struct {
    /*
        药品基础信息     */
    BaseInfoList  *[]AlibabaAlihealthDrugtraceTopYljgQueryRelationBaseInfoDto `json:"base_info_list,omitempty" `

}

func (s *AlibabaAlihealthDrugtraceTopYljgQueryRelationBaseInfosDto) SetBaseInfoList(v []AlibabaAlihealthDrugtraceTopYljgQueryRelationBaseInfoDto) *AlibabaAlihealthDrugtraceTopYljgQueryRelationBaseInfosDto {
    s.BaseInfoList = &v
    return s
}
