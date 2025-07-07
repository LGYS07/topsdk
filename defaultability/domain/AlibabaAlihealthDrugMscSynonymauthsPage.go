package domain


type AlibabaAlihealthDrugMscSynonymauthsPage struct {
    /*
        总数     */
    TotalNum  *int64 `json:"total_num,omitempty" `

    /*
        返回列表     */
    ResultList  *[]AlibabaAlihealthDrugMscSynonymauthsResPSynonymDTO `json:"result_list,omitempty" `

}

func (s *AlibabaAlihealthDrugMscSynonymauthsPage) SetTotalNum(v int64) *AlibabaAlihealthDrugMscSynonymauthsPage {
    s.TotalNum = &v
    return s
}
func (s *AlibabaAlihealthDrugMscSynonymauthsPage) SetResultList(v []AlibabaAlihealthDrugMscSynonymauthsResPSynonymDTO) *AlibabaAlihealthDrugMscSynonymauthsPage {
    s.ResultList = &v
    return s
}
