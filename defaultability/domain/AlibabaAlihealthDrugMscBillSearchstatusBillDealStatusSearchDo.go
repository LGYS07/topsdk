package domain


type AlibabaAlihealthDrugMscBillSearchstatusBillDealStatusSearchDo struct {
    /*
        单据号     */
    StoreInoutSeqNo  *string `json:"store_inout_seq_no,omitempty" `

    /*
        药品类型     */
    PhysicType  *string `json:"physic_type,omitempty" `

    /*
        上传文件名     */
    UploadFileName  *string `json:"upload_file_name,omitempty" `

    /*
        发货单位     */
    FromUserName  *string `json:"from_user_name,omitempty" `

    /*
        角色类型     */
    RoleType  *string `json:"role_type,omitempty" `

    /*
        上传日期     */
    CrtDate  *string `json:"crt_date,omitempty" `

    /*
        操作人标识     */
    IcCode  *string `json:"ic_code,omitempty" `

    /*
        文件名     */
    ShortFileName  *string `json:"short_file_name,omitempty" `

    /*
        企业名称     */
    RefUserName  *string `json:"ref_user_name,omitempty" `

    /*
        单据日期     */
    BillTime  *string `json:"bill_time,omitempty" `

    /*
        处理状态  0，处理中 1, 上传成功     3, 处理成功     4, 处理失败     */
    ResultType  *string `json:"result_type,omitempty" `

    /*
        上传标识     */
    UploadFlag  *string `json:"upload_flag,omitempty" `

    /*
        处理结果表状态(暂不用)     */
    ProcessFlag  *string `json:"process_flag,omitempty" `

    /*
        处理日期     */
    ProcessDate  *string `json:"process_date,omitempty" `

    /*
        单号号     */
    BillCode  *string `json:"bill_code,omitempty" `

    /*
        单据类型     */
    BillType  *string `json:"bill_type,omitempty" `

    /*
        收货单位     */
    ToUserName  *string `json:"to_user_name,omitempty" `

    /*
        发货单位主键     */
    FromUserId  *string `json:"from_user_id,omitempty" `

    /*
        发货单位唯一标识     */
    FromRefUserId  *string `json:"from_ref_user_id,omitempty" `

    /*
        收货单位主键     */
    ToUserId  *string `json:"to_user_id,omitempty" `

    /*
        用户唯一标识     */
    RefUserId  *string `json:"ref_user_id,omitempty" `

    /*
        收货单位唯一标识     */
    ToRefUserId  *string `json:"to_ref_user_id,omitempty" `

    /*
        用户主键     */
    UserId  *string `json:"user_id,omitempty" `

    /*
        处理信息     */
    ProcessInfo  *string `json:"process_info,omitempty" `

    /*
        51全部成功 52部分成功     */
    SubProcessFlag  *string `json:"sub_process_flag,omitempty" `

}

func (s *AlibabaAlihealthDrugMscBillSearchstatusBillDealStatusSearchDo) SetStoreInoutSeqNo(v string) *AlibabaAlihealthDrugMscBillSearchstatusBillDealStatusSearchDo {
    s.StoreInoutSeqNo = &v
    return s
}
func (s *AlibabaAlihealthDrugMscBillSearchstatusBillDealStatusSearchDo) SetPhysicType(v string) *AlibabaAlihealthDrugMscBillSearchstatusBillDealStatusSearchDo {
    s.PhysicType = &v
    return s
}
func (s *AlibabaAlihealthDrugMscBillSearchstatusBillDealStatusSearchDo) SetUploadFileName(v string) *AlibabaAlihealthDrugMscBillSearchstatusBillDealStatusSearchDo {
    s.UploadFileName = &v
    return s
}
func (s *AlibabaAlihealthDrugMscBillSearchstatusBillDealStatusSearchDo) SetFromUserName(v string) *AlibabaAlihealthDrugMscBillSearchstatusBillDealStatusSearchDo {
    s.FromUserName = &v
    return s
}
func (s *AlibabaAlihealthDrugMscBillSearchstatusBillDealStatusSearchDo) SetRoleType(v string) *AlibabaAlihealthDrugMscBillSearchstatusBillDealStatusSearchDo {
    s.RoleType = &v
    return s
}
func (s *AlibabaAlihealthDrugMscBillSearchstatusBillDealStatusSearchDo) SetCrtDate(v string) *AlibabaAlihealthDrugMscBillSearchstatusBillDealStatusSearchDo {
    s.CrtDate = &v
    return s
}
func (s *AlibabaAlihealthDrugMscBillSearchstatusBillDealStatusSearchDo) SetIcCode(v string) *AlibabaAlihealthDrugMscBillSearchstatusBillDealStatusSearchDo {
    s.IcCode = &v
    return s
}
func (s *AlibabaAlihealthDrugMscBillSearchstatusBillDealStatusSearchDo) SetShortFileName(v string) *AlibabaAlihealthDrugMscBillSearchstatusBillDealStatusSearchDo {
    s.ShortFileName = &v
    return s
}
func (s *AlibabaAlihealthDrugMscBillSearchstatusBillDealStatusSearchDo) SetRefUserName(v string) *AlibabaAlihealthDrugMscBillSearchstatusBillDealStatusSearchDo {
    s.RefUserName = &v
    return s
}
func (s *AlibabaAlihealthDrugMscBillSearchstatusBillDealStatusSearchDo) SetBillTime(v string) *AlibabaAlihealthDrugMscBillSearchstatusBillDealStatusSearchDo {
    s.BillTime = &v
    return s
}
func (s *AlibabaAlihealthDrugMscBillSearchstatusBillDealStatusSearchDo) SetResultType(v string) *AlibabaAlihealthDrugMscBillSearchstatusBillDealStatusSearchDo {
    s.ResultType = &v
    return s
}
func (s *AlibabaAlihealthDrugMscBillSearchstatusBillDealStatusSearchDo) SetUploadFlag(v string) *AlibabaAlihealthDrugMscBillSearchstatusBillDealStatusSearchDo {
    s.UploadFlag = &v
    return s
}
func (s *AlibabaAlihealthDrugMscBillSearchstatusBillDealStatusSearchDo) SetProcessFlag(v string) *AlibabaAlihealthDrugMscBillSearchstatusBillDealStatusSearchDo {
    s.ProcessFlag = &v
    return s
}
func (s *AlibabaAlihealthDrugMscBillSearchstatusBillDealStatusSearchDo) SetProcessDate(v string) *AlibabaAlihealthDrugMscBillSearchstatusBillDealStatusSearchDo {
    s.ProcessDate = &v
    return s
}
func (s *AlibabaAlihealthDrugMscBillSearchstatusBillDealStatusSearchDo) SetBillCode(v string) *AlibabaAlihealthDrugMscBillSearchstatusBillDealStatusSearchDo {
    s.BillCode = &v
    return s
}
func (s *AlibabaAlihealthDrugMscBillSearchstatusBillDealStatusSearchDo) SetBillType(v string) *AlibabaAlihealthDrugMscBillSearchstatusBillDealStatusSearchDo {
    s.BillType = &v
    return s
}
func (s *AlibabaAlihealthDrugMscBillSearchstatusBillDealStatusSearchDo) SetToUserName(v string) *AlibabaAlihealthDrugMscBillSearchstatusBillDealStatusSearchDo {
    s.ToUserName = &v
    return s
}
func (s *AlibabaAlihealthDrugMscBillSearchstatusBillDealStatusSearchDo) SetFromUserId(v string) *AlibabaAlihealthDrugMscBillSearchstatusBillDealStatusSearchDo {
    s.FromUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugMscBillSearchstatusBillDealStatusSearchDo) SetFromRefUserId(v string) *AlibabaAlihealthDrugMscBillSearchstatusBillDealStatusSearchDo {
    s.FromRefUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugMscBillSearchstatusBillDealStatusSearchDo) SetToUserId(v string) *AlibabaAlihealthDrugMscBillSearchstatusBillDealStatusSearchDo {
    s.ToUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugMscBillSearchstatusBillDealStatusSearchDo) SetRefUserId(v string) *AlibabaAlihealthDrugMscBillSearchstatusBillDealStatusSearchDo {
    s.RefUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugMscBillSearchstatusBillDealStatusSearchDo) SetToRefUserId(v string) *AlibabaAlihealthDrugMscBillSearchstatusBillDealStatusSearchDo {
    s.ToRefUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugMscBillSearchstatusBillDealStatusSearchDo) SetUserId(v string) *AlibabaAlihealthDrugMscBillSearchstatusBillDealStatusSearchDo {
    s.UserId = &v
    return s
}
func (s *AlibabaAlihealthDrugMscBillSearchstatusBillDealStatusSearchDo) SetProcessInfo(v string) *AlibabaAlihealthDrugMscBillSearchstatusBillDealStatusSearchDo {
    s.ProcessInfo = &v
    return s
}
func (s *AlibabaAlihealthDrugMscBillSearchstatusBillDealStatusSearchDo) SetSubProcessFlag(v string) *AlibabaAlihealthDrugMscBillSearchstatusBillDealStatusSearchDo {
    s.SubProcessFlag = &v
    return s
}
