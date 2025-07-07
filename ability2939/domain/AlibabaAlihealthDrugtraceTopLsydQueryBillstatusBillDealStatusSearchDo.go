package domain


type AlibabaAlihealthDrugtraceTopLsydQueryBillstatusBillDealStatusSearchDo struct {
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

func (s *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusBillDealStatusSearchDo) SetStoreInoutSeqNo(v string) *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusBillDealStatusSearchDo {
    s.StoreInoutSeqNo = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusBillDealStatusSearchDo) SetPhysicType(v string) *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusBillDealStatusSearchDo {
    s.PhysicType = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusBillDealStatusSearchDo) SetUploadFileName(v string) *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusBillDealStatusSearchDo {
    s.UploadFileName = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusBillDealStatusSearchDo) SetFromUserName(v string) *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusBillDealStatusSearchDo {
    s.FromUserName = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusBillDealStatusSearchDo) SetRoleType(v string) *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusBillDealStatusSearchDo {
    s.RoleType = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusBillDealStatusSearchDo) SetCrtDate(v string) *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusBillDealStatusSearchDo {
    s.CrtDate = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusBillDealStatusSearchDo) SetIcCode(v string) *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusBillDealStatusSearchDo {
    s.IcCode = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusBillDealStatusSearchDo) SetShortFileName(v string) *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusBillDealStatusSearchDo {
    s.ShortFileName = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusBillDealStatusSearchDo) SetRefUserName(v string) *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusBillDealStatusSearchDo {
    s.RefUserName = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusBillDealStatusSearchDo) SetBillTime(v string) *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusBillDealStatusSearchDo {
    s.BillTime = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusBillDealStatusSearchDo) SetResultType(v string) *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusBillDealStatusSearchDo {
    s.ResultType = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusBillDealStatusSearchDo) SetUploadFlag(v string) *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusBillDealStatusSearchDo {
    s.UploadFlag = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusBillDealStatusSearchDo) SetProcessFlag(v string) *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusBillDealStatusSearchDo {
    s.ProcessFlag = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusBillDealStatusSearchDo) SetProcessDate(v string) *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusBillDealStatusSearchDo {
    s.ProcessDate = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusBillDealStatusSearchDo) SetBillCode(v string) *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusBillDealStatusSearchDo {
    s.BillCode = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusBillDealStatusSearchDo) SetBillType(v string) *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusBillDealStatusSearchDo {
    s.BillType = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusBillDealStatusSearchDo) SetToUserName(v string) *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusBillDealStatusSearchDo {
    s.ToUserName = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusBillDealStatusSearchDo) SetFromUserId(v string) *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusBillDealStatusSearchDo {
    s.FromUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusBillDealStatusSearchDo) SetFromRefUserId(v string) *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusBillDealStatusSearchDo {
    s.FromRefUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusBillDealStatusSearchDo) SetToUserId(v string) *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusBillDealStatusSearchDo {
    s.ToUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusBillDealStatusSearchDo) SetRefUserId(v string) *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusBillDealStatusSearchDo {
    s.RefUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusBillDealStatusSearchDo) SetToRefUserId(v string) *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusBillDealStatusSearchDo {
    s.ToRefUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusBillDealStatusSearchDo) SetUserId(v string) *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusBillDealStatusSearchDo {
    s.UserId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusBillDealStatusSearchDo) SetProcessInfo(v string) *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusBillDealStatusSearchDo {
    s.ProcessInfo = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusBillDealStatusSearchDo) SetSubProcessFlag(v string) *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusBillDealStatusSearchDo {
    s.SubProcessFlag = &v
    return s
}
