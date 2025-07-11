package domain


type AlibabaAlihealthDrugtraceTopYljgQueryBillstatusPageInfoDto struct {
    /*
        总计     */
    TotalNum  *int64 `json:"total_num,omitempty" `

    /*
        返回列表     */
    ResultList  *[]AlibabaAlihealthDrugtraceTopYljgQueryBillstatusBillDealStatusSearchDo `json:"result_list,omitempty" `

}

func (s *AlibabaAlihealthDrugtraceTopYljgQueryBillstatusPageInfoDto) SetTotalNum(v int64) *AlibabaAlihealthDrugtraceTopYljgQueryBillstatusPageInfoDto {
    s.TotalNum = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryBillstatusPageInfoDto) SetResultList(v []AlibabaAlihealthDrugtraceTopYljgQueryBillstatusBillDealStatusSearchDo) *AlibabaAlihealthDrugtraceTopYljgQueryBillstatusPageInfoDto {
    s.ResultList = &v
    return s
}
