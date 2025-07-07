package domain


type AlibabaAlihealthDrugMscSearchbillPageInfoDto struct {
    /*
        返回结果     */
    ResultList  *[]AlibabaAlihealthDrugMscSearchbillBillChkInOutDo `json:"result_list,omitempty" `

    /*
        总计     */
    TotalNum  *int64 `json:"total_num,omitempty" `

}

func (s *AlibabaAlihealthDrugMscSearchbillPageInfoDto) SetResultList(v []AlibabaAlihealthDrugMscSearchbillBillChkInOutDo) *AlibabaAlihealthDrugMscSearchbillPageInfoDto {
    s.ResultList = &v
    return s
}
func (s *AlibabaAlihealthDrugMscSearchbillPageInfoDto) SetTotalNum(v int64) *AlibabaAlihealthDrugMscSearchbillPageInfoDto {
    s.TotalNum = &v
    return s
}
