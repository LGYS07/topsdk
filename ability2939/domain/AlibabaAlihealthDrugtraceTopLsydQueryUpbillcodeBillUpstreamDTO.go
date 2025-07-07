package domain

import (
        "topsdk/util"
    )

type AlibabaAlihealthDrugtraceTopLsydQueryUpbillcodeBillUpstreamDTO struct {
    /*
        发货企业名称     */
    FromUserName  *string `json:"from_user_name,omitempty" `

    /*
        单据时间     */
    BillTime  *util.LocalTime `json:"bill_time,omitempty" `

    /*
        货主     */
    RefUserId  *string `json:"ref_user_id,omitempty" `

    /*
        发货企业ID     */
    FromUserId  *string `json:"from_user_id,omitempty" `

    /*
        单据类型     */
    BillType  *string `json:"bill_type,omitempty" `

    /*
        收货企业名称     */
    ToUserName  *string `json:"to_user_name,omitempty" `

    /*
        单据号     */
    BillCode  *string `json:"bill_code,omitempty" `

    /*
        收货企业ID     */
    ToUserId  *string `json:"to_user_id,omitempty" `

    /*
        收货企业REF_ENT_ID     */
    ToRefUserId  *string `json:"to_ref_user_id,omitempty" `

    /*
        发货企业REF_ENT_ID     */
    FromRefUserId  *string `json:"from_ref_user_id,omitempty" `

}

func (s *AlibabaAlihealthDrugtraceTopLsydQueryUpbillcodeBillUpstreamDTO) SetFromUserName(v string) *AlibabaAlihealthDrugtraceTopLsydQueryUpbillcodeBillUpstreamDTO {
    s.FromUserName = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryUpbillcodeBillUpstreamDTO) SetBillTime(v util.LocalTime) *AlibabaAlihealthDrugtraceTopLsydQueryUpbillcodeBillUpstreamDTO {
    s.BillTime = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryUpbillcodeBillUpstreamDTO) SetRefUserId(v string) *AlibabaAlihealthDrugtraceTopLsydQueryUpbillcodeBillUpstreamDTO {
    s.RefUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryUpbillcodeBillUpstreamDTO) SetFromUserId(v string) *AlibabaAlihealthDrugtraceTopLsydQueryUpbillcodeBillUpstreamDTO {
    s.FromUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryUpbillcodeBillUpstreamDTO) SetBillType(v string) *AlibabaAlihealthDrugtraceTopLsydQueryUpbillcodeBillUpstreamDTO {
    s.BillType = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryUpbillcodeBillUpstreamDTO) SetToUserName(v string) *AlibabaAlihealthDrugtraceTopLsydQueryUpbillcodeBillUpstreamDTO {
    s.ToUserName = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryUpbillcodeBillUpstreamDTO) SetBillCode(v string) *AlibabaAlihealthDrugtraceTopLsydQueryUpbillcodeBillUpstreamDTO {
    s.BillCode = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryUpbillcodeBillUpstreamDTO) SetToUserId(v string) *AlibabaAlihealthDrugtraceTopLsydQueryUpbillcodeBillUpstreamDTO {
    s.ToUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryUpbillcodeBillUpstreamDTO) SetToRefUserId(v string) *AlibabaAlihealthDrugtraceTopLsydQueryUpbillcodeBillUpstreamDTO {
    s.ToRefUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryUpbillcodeBillUpstreamDTO) SetFromRefUserId(v string) *AlibabaAlihealthDrugtraceTopLsydQueryUpbillcodeBillUpstreamDTO {
    s.FromRefUserId = &v
    return s
}
