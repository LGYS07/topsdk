package domain


type AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailBillInOutDetailDto struct {
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
    BillChkInOutDetailListDTOList  *[]AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailBillchkinoutdetaillistdtolist `json:"bill_chk_in_out_detail_list_d_t_o_list,omitempty" `

}

func (s *AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailBillInOutDetailDto) SetModDate(v string) *AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailBillInOutDetailDto {
    s.ModDate = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailBillInOutDetailDto) SetProcessDate(v string) *AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailBillInOutDetailDto {
    s.ProcessDate = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailBillInOutDetailDto) SetBillTime(v string) *AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailBillInOutDetailDto {
    s.BillTime = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailBillInOutDetailDto) SetToUserId(v string) *AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailBillInOutDetailDto {
    s.ToUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailBillInOutDetailDto) SetToEntName(v string) *AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailBillInOutDetailDto {
    s.ToEntName = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailBillInOutDetailDto) SetFromUserId(v string) *AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailBillInOutDetailDto {
    s.FromUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailBillInOutDetailDto) SetFromEntName(v string) *AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailBillInOutDetailDto {
    s.FromEntName = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailBillInOutDetailDto) SetBillTypeName(v string) *AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailBillInOutDetailDto {
    s.BillTypeName = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailBillInOutDetailDto) SetBillType(v string) *AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailBillInOutDetailDto {
    s.BillType = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailBillInOutDetailDto) SetBillCode(v string) *AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailBillInOutDetailDto {
    s.BillCode = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailBillInOutDetailDto) SetBillChkInOutDetailListDTOList(v []AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailBillchkinoutdetaillistdtolist) *AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailBillInOutDetailDto {
    s.BillChkInOutDetailListDTOList = &v
    return s
}
