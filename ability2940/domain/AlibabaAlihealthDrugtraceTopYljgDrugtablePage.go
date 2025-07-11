package domain


type AlibabaAlihealthDrugtraceTopYljgDrugtablePage struct {
    /*
        总计     */
    TotalNum  *int64 `json:"total_num,omitempty" `

    /*
        返回列表     */
    ResultList  *[]AlibabaAlihealthDrugtraceTopYljgDrugtableDrugTableDto `json:"result_list,omitempty" `

    /*
        当前页     */
    Page  *int64 `json:"page,omitempty" `

    /*
        分页大小     */
    PageSize  *int64 `json:"page_size,omitempty" `

}

func (s *AlibabaAlihealthDrugtraceTopYljgDrugtablePage) SetTotalNum(v int64) *AlibabaAlihealthDrugtraceTopYljgDrugtablePage {
    s.TotalNum = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgDrugtablePage) SetResultList(v []AlibabaAlihealthDrugtraceTopYljgDrugtableDrugTableDto) *AlibabaAlihealthDrugtraceTopYljgDrugtablePage {
    s.ResultList = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgDrugtablePage) SetPage(v int64) *AlibabaAlihealthDrugtraceTopYljgDrugtablePage {
    s.Page = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgDrugtablePage) SetPageSize(v int64) *AlibabaAlihealthDrugtraceTopYljgDrugtablePage {
    s.PageSize = &v
    return s
}
