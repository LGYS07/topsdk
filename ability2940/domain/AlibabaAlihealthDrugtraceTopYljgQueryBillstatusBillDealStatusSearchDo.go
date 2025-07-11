package domain


type AlibabaAlihealthDrugtraceTopYljgQueryBillstatusBillDealStatusSearchDo struct {
    /*
        出入库号     */
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
        创建日期     */
    CrtDate  *string `json:"crt_date,omitempty" `

    /*
        IC码     */
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

func (s *AlibabaAlihealthDrugtraceTopYljgQueryBillstatusBillDealStatusSearchDo) SetStoreInoutSeqNo(v string) *AlibabaAlihealthDrugtraceTopYljgQueryBillstatusBillDealStatusSearchDo {
    s.StoreInoutSeqNo = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryBillstatusBillDealStatusSearchDo) SetPhysicType(v string) *AlibabaAlihealthDrugtraceTopYljgQueryBillstatusBillDealStatusSearchDo {
    s.PhysicType = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryBillstatusBillDealStatusSearchDo) SetUploadFileName(v string) *AlibabaAlihealthDrugtraceTopYljgQueryBillstatusBillDealStatusSearchDo {
    s.UploadFileName = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryBillstatusBillDealStatusSearchDo) SetFromUserName(v string) *AlibabaAlihealthDrugtraceTopYljgQueryBillstatusBillDealStatusSearchDo {
    s.FromUserName = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryBillstatusBillDealStatusSearchDo) SetRoleType(v string) *AlibabaAlihealthDrugtraceTopYljgQueryBillstatusBillDealStatusSearchDo {
    s.RoleType = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryBillstatusBillDealStatusSearchDo) SetCrtDate(v string) *AlibabaAlihealthDrugtraceTopYljgQueryBillstatusBillDealStatusSearchDo {
    s.CrtDate = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryBillstatusBillDealStatusSearchDo) SetIcCode(v string) *AlibabaAlihealthDrugtraceTopYljgQueryBillstatusBillDealStatusSearchDo {
    s.IcCode = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryBillstatusBillDealStatusSearchDo) SetShortFileName(v string) *AlibabaAlihealthDrugtraceTopYljgQueryBillstatusBillDealStatusSearchDo {
    s.ShortFileName = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryBillstatusBillDealStatusSearchDo) SetRefUserName(v string) *AlibabaAlihealthDrugtraceTopYljgQueryBillstatusBillDealStatusSearchDo {
    s.RefUserName = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryBillstatusBillDealStatusSearchDo) SetBillTime(v string) *AlibabaAlihealthDrugtraceTopYljgQueryBillstatusBillDealStatusSearchDo {
    s.BillTime = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryBillstatusBillDealStatusSearchDo) SetResultType(v string) *AlibabaAlihealthDrugtraceTopYljgQueryBillstatusBillDealStatusSearchDo {
    s.ResultType = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryBillstatusBillDealStatusSearchDo) SetUploadFlag(v string) *AlibabaAlihealthDrugtraceTopYljgQueryBillstatusBillDealStatusSearchDo {
    s.UploadFlag = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryBillstatusBillDealStatusSearchDo) SetProcessFlag(v string) *AlibabaAlihealthDrugtraceTopYljgQueryBillstatusBillDealStatusSearchDo {
    s.ProcessFlag = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryBillstatusBillDealStatusSearchDo) SetProcessDate(v string) *AlibabaAlihealthDrugtraceTopYljgQueryBillstatusBillDealStatusSearchDo {
    s.ProcessDate = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryBillstatusBillDealStatusSearchDo) SetBillCode(v string) *AlibabaAlihealthDrugtraceTopYljgQueryBillstatusBillDealStatusSearchDo {
    s.BillCode = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryBillstatusBillDealStatusSearchDo) SetBillType(v string) *AlibabaAlihealthDrugtraceTopYljgQueryBillstatusBillDealStatusSearchDo {
    s.BillType = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryBillstatusBillDealStatusSearchDo) SetToUserName(v string) *AlibabaAlihealthDrugtraceTopYljgQueryBillstatusBillDealStatusSearchDo {
    s.ToUserName = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryBillstatusBillDealStatusSearchDo) SetFromUserId(v string) *AlibabaAlihealthDrugtraceTopYljgQueryBillstatusBillDealStatusSearchDo {
    s.FromUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryBillstatusBillDealStatusSearchDo) SetFromRefUserId(v string) *AlibabaAlihealthDrugtraceTopYljgQueryBillstatusBillDealStatusSearchDo {
    s.FromRefUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryBillstatusBillDealStatusSearchDo) SetToUserId(v string) *AlibabaAlihealthDrugtraceTopYljgQueryBillstatusBillDealStatusSearchDo {
    s.ToUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryBillstatusBillDealStatusSearchDo) SetRefUserId(v string) *AlibabaAlihealthDrugtraceTopYljgQueryBillstatusBillDealStatusSearchDo {
    s.RefUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryBillstatusBillDealStatusSearchDo) SetToRefUserId(v string) *AlibabaAlihealthDrugtraceTopYljgQueryBillstatusBillDealStatusSearchDo {
    s.ToRefUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryBillstatusBillDealStatusSearchDo) SetUserId(v string) *AlibabaAlihealthDrugtraceTopYljgQueryBillstatusBillDealStatusSearchDo {
    s.UserId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryBillstatusBillDealStatusSearchDo) SetProcessInfo(v string) *AlibabaAlihealthDrugtraceTopYljgQueryBillstatusBillDealStatusSearchDo {
    s.ProcessInfo = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryBillstatusBillDealStatusSearchDo) SetSubProcessFlag(v string) *AlibabaAlihealthDrugtraceTopYljgQueryBillstatusBillDealStatusSearchDo {
    s.SubProcessFlag = &v
    return s
}
