package domain


type AlibabaAlihealthDrugtraceTopLsydQueryCodedetailCodeProduceInfoDto struct {
    /*
        生产信息集合     */
    ProduceInfoList  *[]AlibabaAlihealthDrugtraceTopLsydQueryCodedetailProduceInfoDto `json:"produce_info_list,omitempty" `

}

func (s *AlibabaAlihealthDrugtraceTopLsydQueryCodedetailCodeProduceInfoDto) SetProduceInfoList(v []AlibabaAlihealthDrugtraceTopLsydQueryCodedetailProduceInfoDto) *AlibabaAlihealthDrugtraceTopLsydQueryCodedetailCodeProduceInfoDto {
    s.ProduceInfoList = &v
    return s
}
