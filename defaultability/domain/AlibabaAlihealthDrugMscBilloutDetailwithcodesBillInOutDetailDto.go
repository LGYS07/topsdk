package domain


type AlibabaAlihealthDrugMscBilloutDetailwithcodesBillInOutDetailDto struct {
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
    BillChkInOutDetailListDTOList  *[]AlibabaAlihealthDrugMscBilloutDetailwithcodesBillchkinoutdetaillistdtolist `json:"bill_chk_in_out_detail_list_d_t_o_list,omitempty" `

}

func (s *AlibabaAlihealthDrugMscBilloutDetailwithcodesBillInOutDetailDto) SetModDate(v string) *AlibabaAlihealthDrugMscBilloutDetailwithcodesBillInOutDetailDto {
    s.ModDate = &v
    return s
}
func (s *AlibabaAlihealthDrugMscBilloutDetailwithcodesBillInOutDetailDto) SetProcessDate(v string) *AlibabaAlihealthDrugMscBilloutDetailwithcodesBillInOutDetailDto {
    s.ProcessDate = &v
    return s
}
func (s *AlibabaAlihealthDrugMscBilloutDetailwithcodesBillInOutDetailDto) SetBillTime(v string) *AlibabaAlihealthDrugMscBilloutDetailwithcodesBillInOutDetailDto {
    s.BillTime = &v
    return s
}
func (s *AlibabaAlihealthDrugMscBilloutDetailwithcodesBillInOutDetailDto) SetToUserId(v string) *AlibabaAlihealthDrugMscBilloutDetailwithcodesBillInOutDetailDto {
    s.ToUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugMscBilloutDetailwithcodesBillInOutDetailDto) SetToEntName(v string) *AlibabaAlihealthDrugMscBilloutDetailwithcodesBillInOutDetailDto {
    s.ToEntName = &v
    return s
}
func (s *AlibabaAlihealthDrugMscBilloutDetailwithcodesBillInOutDetailDto) SetFromUserId(v string) *AlibabaAlihealthDrugMscBilloutDetailwithcodesBillInOutDetailDto {
    s.FromUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugMscBilloutDetailwithcodesBillInOutDetailDto) SetFromEntName(v string) *AlibabaAlihealthDrugMscBilloutDetailwithcodesBillInOutDetailDto {
    s.FromEntName = &v
    return s
}
func (s *AlibabaAlihealthDrugMscBilloutDetailwithcodesBillInOutDetailDto) SetBillTypeName(v string) *AlibabaAlihealthDrugMscBilloutDetailwithcodesBillInOutDetailDto {
    s.BillTypeName = &v
    return s
}
func (s *AlibabaAlihealthDrugMscBilloutDetailwithcodesBillInOutDetailDto) SetBillType(v string) *AlibabaAlihealthDrugMscBilloutDetailwithcodesBillInOutDetailDto {
    s.BillType = &v
    return s
}
func (s *AlibabaAlihealthDrugMscBilloutDetailwithcodesBillInOutDetailDto) SetBillCode(v string) *AlibabaAlihealthDrugMscBilloutDetailwithcodesBillInOutDetailDto {
    s.BillCode = &v
    return s
}
func (s *AlibabaAlihealthDrugMscBilloutDetailwithcodesBillInOutDetailDto) SetBillChkInOutDetailListDTOList(v []AlibabaAlihealthDrugMscBilloutDetailwithcodesBillchkinoutdetaillistdtolist) *AlibabaAlihealthDrugMscBilloutDetailwithcodesBillInOutDetailDto {
    s.BillChkInOutDetailListDTOList = &v
    return s
}
