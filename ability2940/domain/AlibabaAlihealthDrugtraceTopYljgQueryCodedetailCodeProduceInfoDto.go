package domain


type AlibabaAlihealthDrugtraceTopYljgQueryCodedetailCodeProduceInfoDto struct {
    /*
        生产信息集合     */
    ProduceInfoList  *[]AlibabaAlihealthDrugtraceTopYljgQueryCodedetailProduceInfoDto `json:"produce_info_list,omitempty" `

}

func (s *AlibabaAlihealthDrugtraceTopYljgQueryCodedetailCodeProduceInfoDto) SetProduceInfoList(v []AlibabaAlihealthDrugtraceTopYljgQueryCodedetailProduceInfoDto) *AlibabaAlihealthDrugtraceTopYljgQueryCodedetailCodeProduceInfoDto {
    s.ProduceInfoList = &v
    return s
}
