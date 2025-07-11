package domain

import (
	"github.com/LGYS07/topsdk/util"
)

type AlibabaAlihealthDrugtraceTopYljgQueryRelationCodeActiveInfoDto struct {
	/*
	   处理标志     */
	ProcessFlag *string `json:"process_flag,omitempty" `

	/*
	   状态     */
	Status *string `json:"status,omitempty" `

	/*
	   关联关系类型     */
	RelationType *string `json:"relation_type,omitempty" `

	/*
	   紧急人     */
	UserCert *string `json:"user_cert,omitempty" `

	/*
	   生产编号     */
	ProdCode *string `json:"prod_code,omitempty" `

	/*
	   操作人姓名     */
	OperIcName *string `json:"oper_ic_name,omitempty" `

	/*
	   上传文件名     */
	UploadFileName *string `json:"upload_file_name,omitempty" `

	/*
	   上传文件路径     */
	UploadFilePath *string `json:"upload_file_path,omitempty" `

	/*
	   激活时间     */
	ActiveDate *util.LocalTime `json:"active_date,omitempty" `

	/*
	   企业ID     */
	RefEntId *string `json:"ref_ent_id,omitempty" `

	/*
	   旧企业ID     */
	EntId *string `json:"ent_id,omitempty" `

	/*
	   处理日期     */
	ProcessDate *util.LocalTime `json:"process_date,omitempty" `

	/*
	   处理结束时间     */
	ProcessEndDate *util.LocalTime `json:"process_end_date,omitempty" `

	/*
	   操作人编码     */
	OperIcCode *string `json:"oper_ic_code,omitempty" `

	/*
	   上传标识     */
	UploadFlag *string `json:"upload_flag,omitempty" `

	/*
	   激活时间     */
	CrtDate *util.LocalTime `json:"crt_date,omitempty" `

	/*
	   处理数量     */
	ProcessCount *string `json:"process_count,omitempty" `

	/*
	   总激活数量     */
	ActiveCount *int64 `json:"active_count,omitempty" `

	/*
	   最大包装数量     */
	OtherNum *int64 `json:"other_num,omitempty" `

	/*
	   小码数量     */
	SmallNum *int64 `json:"small_num,omitempty" `

	/*
	   关联关系文件上传日期     */
	CrtDateString *string `json:"crt_date_string,omitempty" `

	/*
	   单据id     */
	BillInId *string `json:"bill_in_id,omitempty" `

	/*
	   激活信息id     */
	CodeActiveInfoId *string `json:"code_active_info_id,omitempty" `
}

func (s *AlibabaAlihealthDrugtraceTopYljgQueryRelationCodeActiveInfoDto) SetProcessFlag(v string) *AlibabaAlihealthDrugtraceTopYljgQueryRelationCodeActiveInfoDto {
	s.ProcessFlag = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryRelationCodeActiveInfoDto) SetStatus(v string) *AlibabaAlihealthDrugtraceTopYljgQueryRelationCodeActiveInfoDto {
	s.Status = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryRelationCodeActiveInfoDto) SetRelationType(v string) *AlibabaAlihealthDrugtraceTopYljgQueryRelationCodeActiveInfoDto {
	s.RelationType = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryRelationCodeActiveInfoDto) SetUserCert(v string) *AlibabaAlihealthDrugtraceTopYljgQueryRelationCodeActiveInfoDto {
	s.UserCert = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryRelationCodeActiveInfoDto) SetProdCode(v string) *AlibabaAlihealthDrugtraceTopYljgQueryRelationCodeActiveInfoDto {
	s.ProdCode = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryRelationCodeActiveInfoDto) SetOperIcName(v string) *AlibabaAlihealthDrugtraceTopYljgQueryRelationCodeActiveInfoDto {
	s.OperIcName = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryRelationCodeActiveInfoDto) SetUploadFileName(v string) *AlibabaAlihealthDrugtraceTopYljgQueryRelationCodeActiveInfoDto {
	s.UploadFileName = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryRelationCodeActiveInfoDto) SetUploadFilePath(v string) *AlibabaAlihealthDrugtraceTopYljgQueryRelationCodeActiveInfoDto {
	s.UploadFilePath = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryRelationCodeActiveInfoDto) SetActiveDate(v util.LocalTime) *AlibabaAlihealthDrugtraceTopYljgQueryRelationCodeActiveInfoDto {
	s.ActiveDate = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryRelationCodeActiveInfoDto) SetRefEntId(v string) *AlibabaAlihealthDrugtraceTopYljgQueryRelationCodeActiveInfoDto {
	s.RefEntId = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryRelationCodeActiveInfoDto) SetEntId(v string) *AlibabaAlihealthDrugtraceTopYljgQueryRelationCodeActiveInfoDto {
	s.EntId = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryRelationCodeActiveInfoDto) SetProcessDate(v util.LocalTime) *AlibabaAlihealthDrugtraceTopYljgQueryRelationCodeActiveInfoDto {
	s.ProcessDate = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryRelationCodeActiveInfoDto) SetProcessEndDate(v util.LocalTime) *AlibabaAlihealthDrugtraceTopYljgQueryRelationCodeActiveInfoDto {
	s.ProcessEndDate = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryRelationCodeActiveInfoDto) SetOperIcCode(v string) *AlibabaAlihealthDrugtraceTopYljgQueryRelationCodeActiveInfoDto {
	s.OperIcCode = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryRelationCodeActiveInfoDto) SetUploadFlag(v string) *AlibabaAlihealthDrugtraceTopYljgQueryRelationCodeActiveInfoDto {
	s.UploadFlag = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryRelationCodeActiveInfoDto) SetCrtDate(v util.LocalTime) *AlibabaAlihealthDrugtraceTopYljgQueryRelationCodeActiveInfoDto {
	s.CrtDate = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryRelationCodeActiveInfoDto) SetProcessCount(v string) *AlibabaAlihealthDrugtraceTopYljgQueryRelationCodeActiveInfoDto {
	s.ProcessCount = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryRelationCodeActiveInfoDto) SetActiveCount(v int64) *AlibabaAlihealthDrugtraceTopYljgQueryRelationCodeActiveInfoDto {
	s.ActiveCount = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryRelationCodeActiveInfoDto) SetOtherNum(v int64) *AlibabaAlihealthDrugtraceTopYljgQueryRelationCodeActiveInfoDto {
	s.OtherNum = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryRelationCodeActiveInfoDto) SetSmallNum(v int64) *AlibabaAlihealthDrugtraceTopYljgQueryRelationCodeActiveInfoDto {
	s.SmallNum = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryRelationCodeActiveInfoDto) SetCrtDateString(v string) *AlibabaAlihealthDrugtraceTopYljgQueryRelationCodeActiveInfoDto {
	s.CrtDateString = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryRelationCodeActiveInfoDto) SetBillInId(v string) *AlibabaAlihealthDrugtraceTopYljgQueryRelationCodeActiveInfoDto {
	s.BillInId = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryRelationCodeActiveInfoDto) SetCodeActiveInfoId(v string) *AlibabaAlihealthDrugtraceTopYljgQueryRelationCodeActiveInfoDto {
	s.CodeActiveInfoId = &v
	return s
}
