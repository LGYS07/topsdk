package domain

import (
	"github.com/LGYS07/topsdk/util"
)

type AlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeBillUpstreamDTO struct {
	/*
	   发货企业名称     */
	FromUserName *string `json:"from_user_name,omitempty" `

	/*
	   单据时间     */
	BillTime *util.LocalTime `json:"bill_time,omitempty" `

	/*
	   货主     */
	RefUserId *string `json:"ref_user_id,omitempty" `

	/*
	   发货企业ID     */
	FromUserId *string `json:"from_user_id,omitempty" `

	/*
	   单据类型     */
	BillType *string `json:"bill_type,omitempty" `

	/*
	   收货企业名称     */
	ToUserName *string `json:"to_user_name,omitempty" `

	/*
	   单据号     */
	BillCode *string `json:"bill_code,omitempty" `

	/*
	   收货企业ID     */
	ToUserId *string `json:"to_user_id,omitempty" `

	/*
	   收货企业REF_ENT_ID     */
	ToRefUserId *string `json:"to_ref_user_id,omitempty" `

	/*
	   发货企业REF_ENT_ID     */
	FromRefUserId *string `json:"from_ref_user_id,omitempty" `

	/*
	   委托企业refEntId     */
	AssRefEntId *string `json:"ass_ref_ent_id,omitempty" `

	/*
	   委托企业entId     */
	AssEntId *string `json:"ass_ent_id,omitempty" `

	/*
	   委托企业entId对应的名称     */
	AssEntName *string `json:"ass_ent_name,omitempty" `
}

func (s *AlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeBillUpstreamDTO) SetFromUserName(v string) *AlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeBillUpstreamDTO {
	s.FromUserName = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeBillUpstreamDTO) SetBillTime(v util.LocalTime) *AlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeBillUpstreamDTO {
	s.BillTime = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeBillUpstreamDTO) SetRefUserId(v string) *AlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeBillUpstreamDTO {
	s.RefUserId = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeBillUpstreamDTO) SetFromUserId(v string) *AlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeBillUpstreamDTO {
	s.FromUserId = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeBillUpstreamDTO) SetBillType(v string) *AlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeBillUpstreamDTO {
	s.BillType = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeBillUpstreamDTO) SetToUserName(v string) *AlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeBillUpstreamDTO {
	s.ToUserName = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeBillUpstreamDTO) SetBillCode(v string) *AlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeBillUpstreamDTO {
	s.BillCode = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeBillUpstreamDTO) SetToUserId(v string) *AlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeBillUpstreamDTO {
	s.ToUserId = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeBillUpstreamDTO) SetToRefUserId(v string) *AlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeBillUpstreamDTO {
	s.ToRefUserId = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeBillUpstreamDTO) SetFromRefUserId(v string) *AlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeBillUpstreamDTO {
	s.FromRefUserId = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeBillUpstreamDTO) SetAssRefEntId(v string) *AlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeBillUpstreamDTO {
	s.AssRefEntId = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeBillUpstreamDTO) SetAssEntId(v string) *AlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeBillUpstreamDTO {
	s.AssEntId = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeBillUpstreamDTO) SetAssEntName(v string) *AlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeBillUpstreamDTO {
	s.AssEntName = &v
	return s
}
