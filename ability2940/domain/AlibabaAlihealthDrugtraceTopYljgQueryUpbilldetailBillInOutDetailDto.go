package domain


type AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailBillInOutDetailDto struct {
    /*
        修改时间     */
    ModDate  *string `json:"mod_date,omitempty" `

    /*
        处理时间     */
    ProcessDate  *string `json:"process_date,omitempty" `

    /*
        单据日期     */
    BillTime  *string `json:"bill_time,omitempty" `

    /*
        收货企业id     */
    ToUserId  *string `json:"to_user_id,omitempty" `

    /*
        收货企业名称     */
    ToEntName  *string `json:"to_ent_name,omitempty" `

    /*
        发货企业id     */
    FromUserId  *string `json:"from_user_id,omitempty" `

    /*
        发货企业名称     */
    FromEntName  *string `json:"from_ent_name,omitempty" `

    /*
        单据类型名称     */
    BillTypeName  *string `json:"bill_type_name,omitempty" `

    /*
        单据类型     */
    BillType  *string `json:"bill_type,omitempty" `

    /*
        单据号码     */
    BillCode  *string `json:"bill_code,omitempty" `

    /*
        单据详情     */
    BillChkInOutDetailListDTOList  *[]AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailBillchkinoutdetaillistdtolist `json:"bill_chk_in_out_detail_list_d_t_o_list,omitempty" `

}

func (s *AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailBillInOutDetailDto) SetModDate(v string) *AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailBillInOutDetailDto {
    s.ModDate = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailBillInOutDetailDto) SetProcessDate(v string) *AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailBillInOutDetailDto {
    s.ProcessDate = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailBillInOutDetailDto) SetBillTime(v string) *AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailBillInOutDetailDto {
    s.BillTime = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailBillInOutDetailDto) SetToUserId(v string) *AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailBillInOutDetailDto {
    s.ToUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailBillInOutDetailDto) SetToEntName(v string) *AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailBillInOutDetailDto {
    s.ToEntName = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailBillInOutDetailDto) SetFromUserId(v string) *AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailBillInOutDetailDto {
    s.FromUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailBillInOutDetailDto) SetFromEntName(v string) *AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailBillInOutDetailDto {
    s.FromEntName = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailBillInOutDetailDto) SetBillTypeName(v string) *AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailBillInOutDetailDto {
    s.BillTypeName = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailBillInOutDetailDto) SetBillType(v string) *AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailBillInOutDetailDto {
    s.BillType = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailBillInOutDetailDto) SetBillCode(v string) *AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailBillInOutDetailDto {
    s.BillCode = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailBillInOutDetailDto) SetBillChkInOutDetailListDTOList(v []AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailBillchkinoutdetaillistdtolist) *AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailBillInOutDetailDto {
    s.BillChkInOutDetailListDTOList = &v
    return s
}
