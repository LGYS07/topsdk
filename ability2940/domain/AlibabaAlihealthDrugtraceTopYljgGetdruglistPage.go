package domain


type AlibabaAlihealthDrugtraceTopYljgGetdruglistPage struct {
    /*
        总计     */
    TotalNum  *int64 `json:"total_num,omitempty" `

    /*
        结果列表     */
    ResultList  *[]AlibabaAlihealthDrugtraceTopYljgGetdruglistDrugDetailInfoDto `json:"result_list,omitempty" `

    /*
        当前页     */
    Page  *int64 `json:"page,omitempty" `

    /*
        页大小     */
    PageSize  *int64 `json:"page_size,omitempty" `

}

func (s *AlibabaAlihealthDrugtraceTopYljgGetdruglistPage) SetTotalNum(v int64) *AlibabaAlihealthDrugtraceTopYljgGetdruglistPage {
    s.TotalNum = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgGetdruglistPage) SetResultList(v []AlibabaAlihealthDrugtraceTopYljgGetdruglistDrugDetailInfoDto) *AlibabaAlihealthDrugtraceTopYljgGetdruglistPage {
    s.ResultList = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgGetdruglistPage) SetPage(v int64) *AlibabaAlihealthDrugtraceTopYljgGetdruglistPage {
    s.Page = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgGetdruglistPage) SetPageSize(v int64) *AlibabaAlihealthDrugtraceTopYljgGetdruglistPage {
    s.PageSize = &v
    return s
}
