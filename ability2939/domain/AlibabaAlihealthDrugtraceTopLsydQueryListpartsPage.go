package domain


type AlibabaAlihealthDrugtraceTopLsydQueryListpartsPage struct {
    /*
        总数     */
    TotalNum  *int64 `json:"total_num,omitempty" `

    /*
        返回列表     */
    ResultList  *[]AlibabaAlihealthDrugtraceTopLsydQueryListpartsPEntParDto `json:"result_list,omitempty" `

}

func (s *AlibabaAlihealthDrugtraceTopLsydQueryListpartsPage) SetTotalNum(v int64) *AlibabaAlihealthDrugtraceTopLsydQueryListpartsPage {
    s.TotalNum = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryListpartsPage) SetResultList(v []AlibabaAlihealthDrugtraceTopLsydQueryListpartsPEntParDto) *AlibabaAlihealthDrugtraceTopLsydQueryListpartsPage {
    s.ResultList = &v
    return s
}
