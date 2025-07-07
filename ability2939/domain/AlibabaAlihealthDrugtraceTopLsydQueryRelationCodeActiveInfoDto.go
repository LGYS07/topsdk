package domain

import (
        "topsdk/util"
    )

type AlibabaAlihealthDrugtraceTopLsydQueryRelationCodeActiveInfoDto struct {
    /*
        处理标志     */
    ProcessFlag  *string `json:"process_flag,omitempty" `

    /*
        状态     */
    Status  *string `json:"status,omitempty" `

    /*
        关联关系类型     */
    RelationType  *string `json:"relation_type,omitempty" `

    /*
        紧急人     */
    UserCert  *string `json:"user_cert,omitempty" `

    /*
        生产编号     */
    ProdCode  *string `json:"prod_code,omitempty" `

    /*
        操作人姓名     */
    OperIcName  *string `json:"oper_ic_name,omitempty" `

    /*
        上传文件名     */
    UploadFileName  *string `json:"upload_file_name,omitempty" `

    /*
        上传文件路径     */
    UploadFilePath  *string `json:"upload_file_path,omitempty" `

    /*
        激活时间     */
    ActiveDate  *util.LocalTime `json:"active_date,omitempty" `

    /*
        企业ID     */
    RefEntId  *string `json:"ref_ent_id,omitempty" `

    /*
        旧企业ID     */
    EntId  *string `json:"ent_id,omitempty" `

    /*
        处理日期     */
    ProcessDate  *util.LocalTime `json:"process_date,omitempty" `

    /*
        处理结束时间     */
    ProcessEndDate  *util.LocalTime `json:"process_end_date,omitempty" `

    /*
        操作人编码     */
    OperIcCode  *string `json:"oper_ic_code,omitempty" `

    /*
        上传标识     */
    UploadFlag  *string `json:"upload_flag,omitempty" `

    /*
        激活时间     */
    CrtDate  *util.LocalTime `json:"crt_date,omitempty" `

    /*
        处理数量     */
    ProcessCount  *string `json:"process_count,omitempty" `

    /*
        总激活数量     */
    ActiveCount  *int64 `json:"active_count,omitempty" `

    /*
        最大包装数量     */
    OtherNum  *int64 `json:"other_num,omitempty" `

    /*
        小码数量     */
    SmallNum  *int64 `json:"small_num,omitempty" `

    /*
        关联关系文件上传日期     */
    CrtDateString  *string `json:"crt_date_string,omitempty" `

    /*
        单据id     */
    BillInId  *string `json:"bill_in_id,omitempty" `

    /*
        激活信息id     */
    CodeActiveInfoId  *string `json:"code_active_info_id,omitempty" `

}

func (s *AlibabaAlihealthDrugtraceTopLsydQueryRelationCodeActiveInfoDto) SetProcessFlag(v string) *AlibabaAlihealthDrugtraceTopLsydQueryRelationCodeActiveInfoDto {
    s.ProcessFlag = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryRelationCodeActiveInfoDto) SetStatus(v string) *AlibabaAlihealthDrugtraceTopLsydQueryRelationCodeActiveInfoDto {
    s.Status = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryRelationCodeActiveInfoDto) SetRelationType(v string) *AlibabaAlihealthDrugtraceTopLsydQueryRelationCodeActiveInfoDto {
    s.RelationType = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryRelationCodeActiveInfoDto) SetUserCert(v string) *AlibabaAlihealthDrugtraceTopLsydQueryRelationCodeActiveInfoDto {
    s.UserCert = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryRelationCodeActiveInfoDto) SetProdCode(v string) *AlibabaAlihealthDrugtraceTopLsydQueryRelationCodeActiveInfoDto {
    s.ProdCode = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryRelationCodeActiveInfoDto) SetOperIcName(v string) *AlibabaAlihealthDrugtraceTopLsydQueryRelationCodeActiveInfoDto {
    s.OperIcName = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryRelationCodeActiveInfoDto) SetUploadFileName(v string) *AlibabaAlihealthDrugtraceTopLsydQueryRelationCodeActiveInfoDto {
    s.UploadFileName = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryRelationCodeActiveInfoDto) SetUploadFilePath(v string) *AlibabaAlihealthDrugtraceTopLsydQueryRelationCodeActiveInfoDto {
    s.UploadFilePath = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryRelationCodeActiveInfoDto) SetActiveDate(v util.LocalTime) *AlibabaAlihealthDrugtraceTopLsydQueryRelationCodeActiveInfoDto {
    s.ActiveDate = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryRelationCodeActiveInfoDto) SetRefEntId(v string) *AlibabaAlihealthDrugtraceTopLsydQueryRelationCodeActiveInfoDto {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryRelationCodeActiveInfoDto) SetEntId(v string) *AlibabaAlihealthDrugtraceTopLsydQueryRelationCodeActiveInfoDto {
    s.EntId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryRelationCodeActiveInfoDto) SetProcessDate(v util.LocalTime) *AlibabaAlihealthDrugtraceTopLsydQueryRelationCodeActiveInfoDto {
    s.ProcessDate = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryRelationCodeActiveInfoDto) SetProcessEndDate(v util.LocalTime) *AlibabaAlihealthDrugtraceTopLsydQueryRelationCodeActiveInfoDto {
    s.ProcessEndDate = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryRelationCodeActiveInfoDto) SetOperIcCode(v string) *AlibabaAlihealthDrugtraceTopLsydQueryRelationCodeActiveInfoDto {
    s.OperIcCode = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryRelationCodeActiveInfoDto) SetUploadFlag(v string) *AlibabaAlihealthDrugtraceTopLsydQueryRelationCodeActiveInfoDto {
    s.UploadFlag = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryRelationCodeActiveInfoDto) SetCrtDate(v util.LocalTime) *AlibabaAlihealthDrugtraceTopLsydQueryRelationCodeActiveInfoDto {
    s.CrtDate = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryRelationCodeActiveInfoDto) SetProcessCount(v string) *AlibabaAlihealthDrugtraceTopLsydQueryRelationCodeActiveInfoDto {
    s.ProcessCount = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryRelationCodeActiveInfoDto) SetActiveCount(v int64) *AlibabaAlihealthDrugtraceTopLsydQueryRelationCodeActiveInfoDto {
    s.ActiveCount = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryRelationCodeActiveInfoDto) SetOtherNum(v int64) *AlibabaAlihealthDrugtraceTopLsydQueryRelationCodeActiveInfoDto {
    s.OtherNum = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryRelationCodeActiveInfoDto) SetSmallNum(v int64) *AlibabaAlihealthDrugtraceTopLsydQueryRelationCodeActiveInfoDto {
    s.SmallNum = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryRelationCodeActiveInfoDto) SetCrtDateString(v string) *AlibabaAlihealthDrugtraceTopLsydQueryRelationCodeActiveInfoDto {
    s.CrtDateString = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryRelationCodeActiveInfoDto) SetBillInId(v string) *AlibabaAlihealthDrugtraceTopLsydQueryRelationCodeActiveInfoDto {
    s.BillInId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryRelationCodeActiveInfoDto) SetCodeActiveInfoId(v string) *AlibabaAlihealthDrugtraceTopLsydQueryRelationCodeActiveInfoDto {
    s.CodeActiveInfoId = &v
    return s
}
