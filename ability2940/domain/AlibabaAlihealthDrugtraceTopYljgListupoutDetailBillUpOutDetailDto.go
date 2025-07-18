package domain


type AlibabaAlihealthDrugtraceTopYljgListupoutDetailBillUpOutDetailDto struct {
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
    DrugInfosDtoList  *[]AlibabaAlihealthDrugtraceTopYljgListupoutDetailDrugInfosDto `json:"drug_infos_dto_list,omitempty" `

}

func (s *AlibabaAlihealthDrugtraceTopYljgListupoutDetailBillUpOutDetailDto) SetBillCode(v string) *AlibabaAlihealthDrugtraceTopYljgListupoutDetailBillUpOutDetailDto {
    s.BillCode = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgListupoutDetailBillUpOutDetailDto) SetBillTypeName(v string) *AlibabaAlihealthDrugtraceTopYljgListupoutDetailBillUpOutDetailDto {
    s.BillTypeName = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgListupoutDetailBillUpOutDetailDto) SetBillType(v string) *AlibabaAlihealthDrugtraceTopYljgListupoutDetailBillUpOutDetailDto {
    s.BillType = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgListupoutDetailBillUpOutDetailDto) SetEntSendName(v string) *AlibabaAlihealthDrugtraceTopYljgListupoutDetailBillUpOutDetailDto {
    s.EntSendName = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgListupoutDetailBillUpOutDetailDto) SetEntSendId(v string) *AlibabaAlihealthDrugtraceTopYljgListupoutDetailBillUpOutDetailDto {
    s.EntSendId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgListupoutDetailBillUpOutDetailDto) SetEntRecvName(v string) *AlibabaAlihealthDrugtraceTopYljgListupoutDetailBillUpOutDetailDto {
    s.EntRecvName = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgListupoutDetailBillUpOutDetailDto) SetEntRecvId(v string) *AlibabaAlihealthDrugtraceTopYljgListupoutDetailBillUpOutDetailDto {
    s.EntRecvId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgListupoutDetailBillUpOutDetailDto) SetStoreOutDate(v string) *AlibabaAlihealthDrugtraceTopYljgListupoutDetailBillUpOutDetailDto {
    s.StoreOutDate = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgListupoutDetailBillUpOutDetailDto) SetUpdateDate(v string) *AlibabaAlihealthDrugtraceTopYljgListupoutDetailBillUpOutDetailDto {
    s.UpdateDate = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgListupoutDetailBillUpOutDetailDto) SetDrugInfosDtoList(v []AlibabaAlihealthDrugtraceTopYljgListupoutDetailDrugInfosDto) *AlibabaAlihealthDrugtraceTopYljgListupoutDetailBillUpOutDetailDto {
    s.DrugInfosDtoList = &v
    return s
}
