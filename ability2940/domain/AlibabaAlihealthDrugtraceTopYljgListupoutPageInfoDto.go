package domain


type AlibabaAlihealthDrugtraceTopYljgListupoutPageInfoDto struct {
    /*
        返回列表     */
    ResultList  *[]AlibabaAlihealthDrugtraceTopYljgListupoutBillUpOutDetailDo `json:"result_list,omitempty" `

    /*
        总数     */
    TotalNum  *int64 `json:"total_num,omitempty" `

}

func (s *AlibabaAlihealthDrugtraceTopYljgListupoutPageInfoDto) SetResultList(v []AlibabaAlihealthDrugtraceTopYljgListupoutBillUpOutDetailDo) *AlibabaAlihealthDrugtraceTopYljgListupoutPageInfoDto {
    s.ResultList = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgListupoutPageInfoDto) SetTotalNum(v int64) *AlibabaAlihealthDrugtraceTopYljgListupoutPageInfoDto {
    s.TotalNum = &v
    return s
}
