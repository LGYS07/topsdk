package domain


type AlibabaAlihealthDrugMscDrugtablePage struct {
    /*
        总数     */
    TotalNum  *int64 `json:"total_num,omitempty" `

    /*
        返回列表     */
    ResultList  *[]AlibabaAlihealthDrugMscDrugtableResDrugDetailInfoDTO `json:"result_list,omitempty" `

    /*
        当前页     */
    Page  *int64 `json:"page,omitempty" `

    /*
        分页大小     */
    PageSize  *int64 `json:"page_size,omitempty" `

}

func (s *AlibabaAlihealthDrugMscDrugtablePage) SetTotalNum(v int64) *AlibabaAlihealthDrugMscDrugtablePage {
    s.TotalNum = &v
    return s
}
func (s *AlibabaAlihealthDrugMscDrugtablePage) SetResultList(v []AlibabaAlihealthDrugMscDrugtableResDrugDetailInfoDTO) *AlibabaAlihealthDrugMscDrugtablePage {
    s.ResultList = &v
    return s
}
func (s *AlibabaAlihealthDrugMscDrugtablePage) SetPage(v int64) *AlibabaAlihealthDrugMscDrugtablePage {
    s.Page = &v
    return s
}
func (s *AlibabaAlihealthDrugMscDrugtablePage) SetPageSize(v int64) *AlibabaAlihealthDrugMscDrugtablePage {
    s.PageSize = &v
    return s
}
