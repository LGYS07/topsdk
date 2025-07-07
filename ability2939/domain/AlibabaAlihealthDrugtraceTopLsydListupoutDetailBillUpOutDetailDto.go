package domain


type AlibabaAlihealthDrugtraceTopLsydListupoutDetailBillUpOutDetailDto struct {
    /*
        单据编码     */
    BillCode  *string `json:"bill_code,omitempty" `

    /*
        单据类型描述     */
    BillTypeName  *string `json:"bill_type_name,omitempty" `

    /*
        单据类型     */
    BillType  *string `json:"bill_type,omitempty" `

    /*
        发货企业名称     */
    EntSendName  *string `json:"ent_send_name,omitempty" `

    /*
        发货企业的ref_ent_id     */
    EntSendId  *string `json:"ent_send_id,omitempty" `

    /*
        收货企业名称     */
    EntRecvName  *string `json:"ent_recv_name,omitempty" `

    /*
        收货企业ref_ent_id     */
    EntRecvId  *string `json:"ent_recv_id,omitempty" `

    /*
        单据日期     */
    StoreOutDate  *string `json:"store_out_date,omitempty" `

    /*
        最后更新时间     */
    UpdateDate  *string `json:"update_date,omitempty" `

    /*
        药品信息数据     */
    DrugInfosDtoList  *[]AlibabaAlihealthDrugtraceTopLsydListupoutDetailDrugInfosDto `json:"drug_infos_dto_list,omitempty" `

}

func (s *AlibabaAlihealthDrugtraceTopLsydListupoutDetailBillUpOutDetailDto) SetBillCode(v string) *AlibabaAlihealthDrugtraceTopLsydListupoutDetailBillUpOutDetailDto {
    s.BillCode = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydListupoutDetailBillUpOutDetailDto) SetBillTypeName(v string) *AlibabaAlihealthDrugtraceTopLsydListupoutDetailBillUpOutDetailDto {
    s.BillTypeName = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydListupoutDetailBillUpOutDetailDto) SetBillType(v string) *AlibabaAlihealthDrugtraceTopLsydListupoutDetailBillUpOutDetailDto {
    s.BillType = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydListupoutDetailBillUpOutDetailDto) SetEntSendName(v string) *AlibabaAlihealthDrugtraceTopLsydListupoutDetailBillUpOutDetailDto {
    s.EntSendName = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydListupoutDetailBillUpOutDetailDto) SetEntSendId(v string) *AlibabaAlihealthDrugtraceTopLsydListupoutDetailBillUpOutDetailDto {
    s.EntSendId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydListupoutDetailBillUpOutDetailDto) SetEntRecvName(v string) *AlibabaAlihealthDrugtraceTopLsydListupoutDetailBillUpOutDetailDto {
    s.EntRecvName = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydListupoutDetailBillUpOutDetailDto) SetEntRecvId(v string) *AlibabaAlihealthDrugtraceTopLsydListupoutDetailBillUpOutDetailDto {
    s.EntRecvId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydListupoutDetailBillUpOutDetailDto) SetStoreOutDate(v string) *AlibabaAlihealthDrugtraceTopLsydListupoutDetailBillUpOutDetailDto {
    s.StoreOutDate = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydListupoutDetailBillUpOutDetailDto) SetUpdateDate(v string) *AlibabaAlihealthDrugtraceTopLsydListupoutDetailBillUpOutDetailDto {
    s.UpdateDate = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydListupoutDetailBillUpOutDetailDto) SetDrugInfosDtoList(v []AlibabaAlihealthDrugtraceTopLsydListupoutDetailDrugInfosDto) *AlibabaAlihealthDrugtraceTopLsydListupoutDetailBillUpOutDetailDto {
    s.DrugInfosDtoList = &v
    return s
}
