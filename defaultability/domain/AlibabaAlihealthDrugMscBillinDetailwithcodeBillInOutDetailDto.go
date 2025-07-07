package domain


type AlibabaAlihealthDrugMscBillinDetailwithcodeBillInOutDetailDto struct {
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
    BillChkInOutDetailListDTOList  *[]AlibabaAlihealthDrugMscBillinDetailwithcodeBillchkinoutdetaillistdtolist `json:"bill_chk_in_out_detail_list_d_t_o_list,omitempty" `

}

func (s *AlibabaAlihealthDrugMscBillinDetailwithcodeBillInOutDetailDto) SetModDate(v string) *AlibabaAlihealthDrugMscBillinDetailwithcodeBillInOutDetailDto {
    s.ModDate = &v
    return s
}
func (s *AlibabaAlihealthDrugMscBillinDetailwithcodeBillInOutDetailDto) SetProcessDate(v string) *AlibabaAlihealthDrugMscBillinDetailwithcodeBillInOutDetailDto {
    s.ProcessDate = &v
    return s
}
func (s *AlibabaAlihealthDrugMscBillinDetailwithcodeBillInOutDetailDto) SetBillTime(v string) *AlibabaAlihealthDrugMscBillinDetailwithcodeBillInOutDetailDto {
    s.BillTime = &v
    return s
}
func (s *AlibabaAlihealthDrugMscBillinDetailwithcodeBillInOutDetailDto) SetToUserId(v string) *AlibabaAlihealthDrugMscBillinDetailwithcodeBillInOutDetailDto {
    s.ToUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugMscBillinDetailwithcodeBillInOutDetailDto) SetToEntName(v string) *AlibabaAlihealthDrugMscBillinDetailwithcodeBillInOutDetailDto {
    s.ToEntName = &v
    return s
}
func (s *AlibabaAlihealthDrugMscBillinDetailwithcodeBillInOutDetailDto) SetFromUserId(v string) *AlibabaAlihealthDrugMscBillinDetailwithcodeBillInOutDetailDto {
    s.FromUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugMscBillinDetailwithcodeBillInOutDetailDto) SetFromEntName(v string) *AlibabaAlihealthDrugMscBillinDetailwithcodeBillInOutDetailDto {
    s.FromEntName = &v
    return s
}
func (s *AlibabaAlihealthDrugMscBillinDetailwithcodeBillInOutDetailDto) SetBillTypeName(v string) *AlibabaAlihealthDrugMscBillinDetailwithcodeBillInOutDetailDto {
    s.BillTypeName = &v
    return s
}
func (s *AlibabaAlihealthDrugMscBillinDetailwithcodeBillInOutDetailDto) SetBillType(v string) *AlibabaAlihealthDrugMscBillinDetailwithcodeBillInOutDetailDto {
    s.BillType = &v
    return s
}
func (s *AlibabaAlihealthDrugMscBillinDetailwithcodeBillInOutDetailDto) SetBillCode(v string) *AlibabaAlihealthDrugMscBillinDetailwithcodeBillInOutDetailDto {
    s.BillCode = &v
    return s
}
func (s *AlibabaAlihealthDrugMscBillinDetailwithcodeBillInOutDetailDto) SetBillChkInOutDetailListDTOList(v []AlibabaAlihealthDrugMscBillinDetailwithcodeBillchkinoutdetaillistdtolist) *AlibabaAlihealthDrugMscBillinDetailwithcodeBillInOutDetailDto {
    s.BillChkInOutDetailListDTOList = &v
    return s
}
