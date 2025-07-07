package domain


type AlibabaAlihealthDrugtraceTopLsydQueryBillstatusPageInfoDto struct {
    /*
        总计     */
    TotalNum  *int64 `json:"total_num,omitempty" `

    /*
        返回列表     */
    ResultList  *[]AlibabaAlihealthDrugtraceTopLsydQueryBillstatusBillDealStatusSearchDo `json:"result_list,omitempty" `

}

func (s *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusPageInfoDto) SetTotalNum(v int64) *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusPageInfoDto {
    s.TotalNum = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusPageInfoDto) SetResultList(v []AlibabaAlihealthDrugtraceTopLsydQueryBillstatusBillDealStatusSearchDo) *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusPageInfoDto {
    s.ResultList = &v
    return s
}
