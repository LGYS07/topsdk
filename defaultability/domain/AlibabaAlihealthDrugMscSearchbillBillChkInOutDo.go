package domain


type AlibabaAlihealthDrugMscSearchbillBillChkInOutDo struct {
    /*
        单据类型     */
    BillType  *string `json:"bill_type,omitempty" `

    /*
        单号号码     */
    BillCode  *string `json:"bill_code,omitempty" `

    /*
        发货企业ID     */
    FromUserId  *string `json:"from_user_id,omitempty" `

    /*
        企业名称     */
    RefUserName  *string `json:"ref_user_name,omitempty" `

    /*
        企业ID     */
    RefUserId  *string `json:"ref_user_id,omitempty" `

    /*
        生产日期     */
    ProduceDate  *string `json:"produce_date,omitempty" `

    /*
        上传文件名     */
    UploadFileName  *string `json:"upload_file_name,omitempty" `

    /*
        发货单位     */
    FromUserName  *string `json:"from_user_name,omitempty" `

    /*
        收货单位     */
    ToUserId  *string `json:"to_user_id,omitempty" `

    /*
        生产企业ID     */
    ProduceEntId  *string `json:"produce_ent_id,omitempty" `

    /*
        单据时间     */
    BillTime  *string `json:"bill_time,omitempty" `

    /*
        角色类型     */
    UserRoleType  *string `json:"user_role_type,omitempty" `

    /*
        处理日期     */
    ProcessDate  *string `json:"process_date,omitempty" `

    /*
        单据ID     */
    BillId  *string `json:"bill_id,omitempty" `

    /*
        收货单位     */
    ToUserName  *string `json:"to_user_name,omitempty" `

    /*
        代理企业     */
    AgentUserName  *string `json:"agent_user_name,omitempty" `

    /*
        代理企业ID     */
    AgentRefUserId  *string `json:"agent_ref_user_id,omitempty" `

    /*
        单据类型     */
    BillTypeName  *string `json:"bill_type_name,omitempty" `

    /*
        收货单位ID     */
    ToRefUserId  *string `json:"to_ref_user_id,omitempty" `

    /*
        发货单位ID     */
    FromRefUserId  *string `json:"from_ref_user_id,omitempty" `

    /*
        51全部成功 52部分成功     */
    SubProcessFlag  *string `json:"sub_process_flag,omitempty" `

}

func (s *AlibabaAlihealthDrugMscSearchbillBillChkInOutDo) SetBillType(v string) *AlibabaAlihealthDrugMscSearchbillBillChkInOutDo {
    s.BillType = &v
    return s
}
func (s *AlibabaAlihealthDrugMscSearchbillBillChkInOutDo) SetBillCode(v string) *AlibabaAlihealthDrugMscSearchbillBillChkInOutDo {
    s.BillCode = &v
    return s
}
func (s *AlibabaAlihealthDrugMscSearchbillBillChkInOutDo) SetFromUserId(v string) *AlibabaAlihealthDrugMscSearchbillBillChkInOutDo {
    s.FromUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugMscSearchbillBillChkInOutDo) SetRefUserName(v string) *AlibabaAlihealthDrugMscSearchbillBillChkInOutDo {
    s.RefUserName = &v
    return s
}
func (s *AlibabaAlihealthDrugMscSearchbillBillChkInOutDo) SetRefUserId(v string) *AlibabaAlihealthDrugMscSearchbillBillChkInOutDo {
    s.RefUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugMscSearchbillBillChkInOutDo) SetProduceDate(v string) *AlibabaAlihealthDrugMscSearchbillBillChkInOutDo {
    s.ProduceDate = &v
    return s
}
func (s *AlibabaAlihealthDrugMscSearchbillBillChkInOutDo) SetUploadFileName(v string) *AlibabaAlihealthDrugMscSearchbillBillChkInOutDo {
    s.UploadFileName = &v
    return s
}
func (s *AlibabaAlihealthDrugMscSearchbillBillChkInOutDo) SetFromUserName(v string) *AlibabaAlihealthDrugMscSearchbillBillChkInOutDo {
    s.FromUserName = &v
    return s
}
func (s *AlibabaAlihealthDrugMscSearchbillBillChkInOutDo) SetToUserId(v string) *AlibabaAlihealthDrugMscSearchbillBillChkInOutDo {
    s.ToUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugMscSearchbillBillChkInOutDo) SetProduceEntId(v string) *AlibabaAlihealthDrugMscSearchbillBillChkInOutDo {
    s.ProduceEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugMscSearchbillBillChkInOutDo) SetBillTime(v string) *AlibabaAlihealthDrugMscSearchbillBillChkInOutDo {
    s.BillTime = &v
    return s
}
func (s *AlibabaAlihealthDrugMscSearchbillBillChkInOutDo) SetUserRoleType(v string) *AlibabaAlihealthDrugMscSearchbillBillChkInOutDo {
    s.UserRoleType = &v
    return s
}
func (s *AlibabaAlihealthDrugMscSearchbillBillChkInOutDo) SetProcessDate(v string) *AlibabaAlihealthDrugMscSearchbillBillChkInOutDo {
    s.ProcessDate = &v
    return s
}
func (s *AlibabaAlihealthDrugMscSearchbillBillChkInOutDo) SetBillId(v string) *AlibabaAlihealthDrugMscSearchbillBillChkInOutDo {
    s.BillId = &v
    return s
}
func (s *AlibabaAlihealthDrugMscSearchbillBillChkInOutDo) SetToUserName(v string) *AlibabaAlihealthDrugMscSearchbillBillChkInOutDo {
    s.ToUserName = &v
    return s
}
func (s *AlibabaAlihealthDrugMscSearchbillBillChkInOutDo) SetAgentUserName(v string) *AlibabaAlihealthDrugMscSearchbillBillChkInOutDo {
    s.AgentUserName = &v
    return s
}
func (s *AlibabaAlihealthDrugMscSearchbillBillChkInOutDo) SetAgentRefUserId(v string) *AlibabaAlihealthDrugMscSearchbillBillChkInOutDo {
    s.AgentRefUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugMscSearchbillBillChkInOutDo) SetBillTypeName(v string) *AlibabaAlihealthDrugMscSearchbillBillChkInOutDo {
    s.BillTypeName = &v
    return s
}
func (s *AlibabaAlihealthDrugMscSearchbillBillChkInOutDo) SetToRefUserId(v string) *AlibabaAlihealthDrugMscSearchbillBillChkInOutDo {
    s.ToRefUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugMscSearchbillBillChkInOutDo) SetFromRefUserId(v string) *AlibabaAlihealthDrugMscSearchbillBillChkInOutDo {
    s.FromRefUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugMscSearchbillBillChkInOutDo) SetSubProcessFlag(v string) *AlibabaAlihealthDrugMscSearchbillBillChkInOutDo {
    s.SubProcessFlag = &v
    return s
}
