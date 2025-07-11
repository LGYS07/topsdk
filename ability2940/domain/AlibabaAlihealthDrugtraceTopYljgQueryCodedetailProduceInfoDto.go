package domain


type AlibabaAlihealthDrugtraceTopYljgQueryCodedetailProduceInfoDto struct {
    /*
        生产日期     */
    ProduceDateStr  *string `json:"produce_date_str,omitempty" `

    /*
        最小包装数量     */
    PkgAmount  *string `json:"pkg_amount,omitempty" `

    /*
        有效期至     */
    ExpireDate  *string `json:"expire_date,omitempty" `

    /*
        批次号     */
    BatchNo  *string `json:"batch_no,omitempty" `

}

func (s *AlibabaAlihealthDrugtraceTopYljgQueryCodedetailProduceInfoDto) SetProduceDateStr(v string) *AlibabaAlihealthDrugtraceTopYljgQueryCodedetailProduceInfoDto {
    s.ProduceDateStr = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryCodedetailProduceInfoDto) SetPkgAmount(v string) *AlibabaAlihealthDrugtraceTopYljgQueryCodedetailProduceInfoDto {
    s.PkgAmount = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryCodedetailProduceInfoDto) SetExpireDate(v string) *AlibabaAlihealthDrugtraceTopYljgQueryCodedetailProduceInfoDto {
    s.ExpireDate = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryCodedetailProduceInfoDto) SetBatchNo(v string) *AlibabaAlihealthDrugtraceTopYljgQueryCodedetailProduceInfoDto {
    s.BatchNo = &v
    return s
}
