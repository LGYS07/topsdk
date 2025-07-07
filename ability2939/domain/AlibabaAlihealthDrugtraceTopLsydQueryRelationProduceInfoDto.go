package domain

import (
	"github.com/LGYS07/topsdk/util"
)

type AlibabaAlihealthDrugtraceTopLsydQueryRelationProduceInfoDto struct {
	/*
	   生产日期     */
	ProduceDate *util.LocalTime `json:"produce_date,omitempty" `

	/*
	   有效期     */
	ExpireDate *string `json:"expire_date,omitempty" `

	/*
	   批次号     */
	BatchNo *string `json:"batch_no,omitempty" `

	/*
	   码     */
	Code *string `json:"code,omitempty" `

	/*
	   最小包装数量     */
	PkgAmount *string `json:"pkg_amount,omitempty" `
}

func (s *AlibabaAlihealthDrugtraceTopLsydQueryRelationProduceInfoDto) SetProduceDate(v util.LocalTime) *AlibabaAlihealthDrugtraceTopLsydQueryRelationProduceInfoDto {
	s.ProduceDate = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryRelationProduceInfoDto) SetExpireDate(v string) *AlibabaAlihealthDrugtraceTopLsydQueryRelationProduceInfoDto {
	s.ExpireDate = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryRelationProduceInfoDto) SetBatchNo(v string) *AlibabaAlihealthDrugtraceTopLsydQueryRelationProduceInfoDto {
	s.BatchNo = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryRelationProduceInfoDto) SetCode(v string) *AlibabaAlihealthDrugtraceTopLsydQueryRelationProduceInfoDto {
	s.Code = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryRelationProduceInfoDto) SetPkgAmount(v string) *AlibabaAlihealthDrugtraceTopLsydQueryRelationProduceInfoDto {
	s.PkgAmount = &v
	return s
}
