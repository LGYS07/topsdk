package domain


type AlibabaAlihealthDrugMscListpartsPage struct {
    /*
        总数     */
    TotalNum  *int64 `json:"total_num,omitempty" `

    /*
        返回列表     */
    ResultList  *[]AlibabaAlihealthDrugMscListpartsPEntParDTO `json:"result_list,omitempty" `

}

func (s *AlibabaAlihealthDrugMscListpartsPage) SetTotalNum(v int64) *AlibabaAlihealthDrugMscListpartsPage {
    s.TotalNum = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListpartsPage) SetResultList(v []AlibabaAlihealthDrugMscListpartsPEntParDTO) *AlibabaAlihealthDrugMscListpartsPage {
    s.ResultList = &v
    return s
}
