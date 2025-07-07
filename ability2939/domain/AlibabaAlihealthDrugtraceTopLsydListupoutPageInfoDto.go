package domain


type AlibabaAlihealthDrugtraceTopLsydListupoutPageInfoDto struct {
    /*
        返回列表     */
    ResultList  *[]AlibabaAlihealthDrugtraceTopLsydListupoutBillUpOutDetailDo `json:"result_list,omitempty" `

    /*
        总数     */
    TotalNum  *int64 `json:"total_num,omitempty" `

}

func (s *AlibabaAlihealthDrugtraceTopLsydListupoutPageInfoDto) SetResultList(v []AlibabaAlihealthDrugtraceTopLsydListupoutBillUpOutDetailDo) *AlibabaAlihealthDrugtraceTopLsydListupoutPageInfoDto {
    s.ResultList = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydListupoutPageInfoDto) SetTotalNum(v int64) *AlibabaAlihealthDrugtraceTopLsydListupoutPageInfoDto {
    s.TotalNum = &v
    return s
}
