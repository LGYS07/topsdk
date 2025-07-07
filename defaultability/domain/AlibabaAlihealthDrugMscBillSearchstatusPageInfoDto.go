package domain


type AlibabaAlihealthDrugMscBillSearchstatusPageInfoDto struct {
    /*
        总计     */
    TotalNum  *int64 `json:"total_num,omitempty" `

    /*
        返回列表     */
    ResultList  *[]AlibabaAlihealthDrugMscBillSearchstatusBillDealStatusSearchDo `json:"result_list,omitempty" `

}

func (s *AlibabaAlihealthDrugMscBillSearchstatusPageInfoDto) SetTotalNum(v int64) *AlibabaAlihealthDrugMscBillSearchstatusPageInfoDto {
    s.TotalNum = &v
    return s
}
func (s *AlibabaAlihealthDrugMscBillSearchstatusPageInfoDto) SetResultList(v []AlibabaAlihealthDrugMscBillSearchstatusBillDealStatusSearchDo) *AlibabaAlihealthDrugMscBillSearchstatusPageInfoDto {
    s.ResultList = &v
    return s
}
