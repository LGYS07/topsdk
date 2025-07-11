package domain


type AlibabaAlihealthDrugtraceTopYljgQueryListpartsPage struct {
    /*
        总数     */
    TotalNum  *int64 `json:"total_num,omitempty" `

    /*
        返回列表     */
    ResultList  *[]AlibabaAlihealthDrugtraceTopYljgQueryListpartsPEntParDto `json:"result_list,omitempty" `

}

func (s *AlibabaAlihealthDrugtraceTopYljgQueryListpartsPage) SetTotalNum(v int64) *AlibabaAlihealthDrugtraceTopYljgQueryListpartsPage {
    s.TotalNum = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryListpartsPage) SetResultList(v []AlibabaAlihealthDrugtraceTopYljgQueryListpartsPEntParDto) *AlibabaAlihealthDrugtraceTopYljgQueryListpartsPage {
    s.ResultList = &v
    return s
}
