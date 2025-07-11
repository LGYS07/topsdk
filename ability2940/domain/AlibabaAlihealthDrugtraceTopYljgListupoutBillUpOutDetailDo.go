package domain


type AlibabaAlihealthDrugtraceTopYljgListupoutBillUpOutDetailDo struct {
    /*
        发货单位     */
    FromEntName  *string `json:"from_ent_name,omitempty" `

    /*
        最小码量     */
    CodeCount  *int64 `json:"code_count,omitempty" `

    /*
        失效日期     */
    ExprieDate  *string `json:"exprie_date,omitempty" `

    /*
        厂商     */
    ProduceEntName  *string `json:"produce_ent_name,omitempty" `

    /*
        生产日期     */
    ProduceDate  *string `json:"produce_date,omitempty" `

    /*
        生产批号     */
    ProduceBatchNo  *string `json:"produce_batch_no,omitempty" `

    /*
        包装规格     */
    PkgSpec  *string `json:"pkg_spec,omitempty" `

    /*
        药品信息     */
    PhysicInfo  *string `json:"physic_info,omitempty" `

    /*
        药品名称     */
    PhysicName  *string `json:"physic_name,omitempty" `

    /*
        制剂数量     */
    PrepnCount  *int64 `json:"prepn_count,omitempty" `

    /*
        发货单位REF_ENT_ID     */
    FromRefUserId  *string `json:"from_ref_user_id,omitempty" `

    /*
        收货单位REF_ENT_ID     */
    ToRefUserId  *string `json:"to_ref_user_id,omitempty" `

    /*
        单据时间     */
    BillTime  *string `json:"bill_time,omitempty" `

    /*
        单据码     */
    BillCode  *string `json:"bill_code,omitempty" `

    /*
        单据类型     */
    BillType  *string `json:"bill_type,omitempty" `

    /*
        发货企业     */
    ToUserName  *string `json:"to_user_name,omitempty" `

    /*
        收货企业     */
    FromUserName  *string `json:"from_user_name,omitempty" `

    /*
        失效日期格式化     */
    ExprieDateFormat  *string `json:"exprie_date_format,omitempty" `

    /*
        单据时间格式化     */
    BillTimeFormat  *string `json:"bill_time_format,omitempty" `

    /*
        单据ID     */
    BillOutId  *int64 `json:"bill_out_id,omitempty" `

    /*
        制剂单位     */
    PrepnUnit  *int64 `json:"prepn_unit,omitempty" `

    /*
        制剂规格     */
    PrepnSpec  *string `json:"prepn_spec,omitempty" `

    /*
        药品ID     */
    DrugEntBaseInfoId  *string `json:"drug_ent_base_info_id,omitempty" `

    /*
        生产日期格式化     */
    ProduceDateFormat  *string `json:"produce_date_format,omitempty" `

    /*
        确认状态1未确认2已确认     */
    Status  *string `json:"status,omitempty" `

    /*
        收货企业ent_id     */
    ToUserId  *string `json:"to_user_id,omitempty" `

    /*
        发货企业ent_id     */
    FromUserId  *string `json:"from_user_id,omitempty" `

    /*
        委托企业refEntId     */
    AssRefEntId  *string `json:"ass_ref_ent_id,omitempty" `

    /*
        委托企业entId     */
    AssEntId  *string `json:"ass_ent_id,omitempty" `

}

func (s *AlibabaAlihealthDrugtraceTopYljgListupoutBillUpOutDetailDo) SetFromEntName(v string) *AlibabaAlihealthDrugtraceTopYljgListupoutBillUpOutDetailDo {
    s.FromEntName = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgListupoutBillUpOutDetailDo) SetCodeCount(v int64) *AlibabaAlihealthDrugtraceTopYljgListupoutBillUpOutDetailDo {
    s.CodeCount = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgListupoutBillUpOutDetailDo) SetExprieDate(v string) *AlibabaAlihealthDrugtraceTopYljgListupoutBillUpOutDetailDo {
    s.ExprieDate = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgListupoutBillUpOutDetailDo) SetProduceEntName(v string) *AlibabaAlihealthDrugtraceTopYljgListupoutBillUpOutDetailDo {
    s.ProduceEntName = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgListupoutBillUpOutDetailDo) SetProduceDate(v string) *AlibabaAlihealthDrugtraceTopYljgListupoutBillUpOutDetailDo {
    s.ProduceDate = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgListupoutBillUpOutDetailDo) SetProduceBatchNo(v string) *AlibabaAlihealthDrugtraceTopYljgListupoutBillUpOutDetailDo {
    s.ProduceBatchNo = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgListupoutBillUpOutDetailDo) SetPkgSpec(v string) *AlibabaAlihealthDrugtraceTopYljgListupoutBillUpOutDetailDo {
    s.PkgSpec = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgListupoutBillUpOutDetailDo) SetPhysicInfo(v string) *AlibabaAlihealthDrugtraceTopYljgListupoutBillUpOutDetailDo {
    s.PhysicInfo = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgListupoutBillUpOutDetailDo) SetPhysicName(v string) *AlibabaAlihealthDrugtraceTopYljgListupoutBillUpOutDetailDo {
    s.PhysicName = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgListupoutBillUpOutDetailDo) SetPrepnCount(v int64) *AlibabaAlihealthDrugtraceTopYljgListupoutBillUpOutDetailDo {
    s.PrepnCount = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgListupoutBillUpOutDetailDo) SetFromRefUserId(v string) *AlibabaAlihealthDrugtraceTopYljgListupoutBillUpOutDetailDo {
    s.FromRefUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgListupoutBillUpOutDetailDo) SetToRefUserId(v string) *AlibabaAlihealthDrugtraceTopYljgListupoutBillUpOutDetailDo {
    s.ToRefUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgListupoutBillUpOutDetailDo) SetBillTime(v string) *AlibabaAlihealthDrugtraceTopYljgListupoutBillUpOutDetailDo {
    s.BillTime = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgListupoutBillUpOutDetailDo) SetBillCode(v string) *AlibabaAlihealthDrugtraceTopYljgListupoutBillUpOutDetailDo {
    s.BillCode = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgListupoutBillUpOutDetailDo) SetBillType(v string) *AlibabaAlihealthDrugtraceTopYljgListupoutBillUpOutDetailDo {
    s.BillType = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgListupoutBillUpOutDetailDo) SetToUserName(v string) *AlibabaAlihealthDrugtraceTopYljgListupoutBillUpOutDetailDo {
    s.ToUserName = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgListupoutBillUpOutDetailDo) SetFromUserName(v string) *AlibabaAlihealthDrugtraceTopYljgListupoutBillUpOutDetailDo {
    s.FromUserName = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgListupoutBillUpOutDetailDo) SetExprieDateFormat(v string) *AlibabaAlihealthDrugtraceTopYljgListupoutBillUpOutDetailDo {
    s.ExprieDateFormat = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgListupoutBillUpOutDetailDo) SetBillTimeFormat(v string) *AlibabaAlihealthDrugtraceTopYljgListupoutBillUpOutDetailDo {
    s.BillTimeFormat = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgListupoutBillUpOutDetailDo) SetBillOutId(v int64) *AlibabaAlihealthDrugtraceTopYljgListupoutBillUpOutDetailDo {
    s.BillOutId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgListupoutBillUpOutDetailDo) SetPrepnUnit(v int64) *AlibabaAlihealthDrugtraceTopYljgListupoutBillUpOutDetailDo {
    s.PrepnUnit = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgListupoutBillUpOutDetailDo) SetPrepnSpec(v string) *AlibabaAlihealthDrugtraceTopYljgListupoutBillUpOutDetailDo {
    s.PrepnSpec = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgListupoutBillUpOutDetailDo) SetDrugEntBaseInfoId(v string) *AlibabaAlihealthDrugtraceTopYljgListupoutBillUpOutDetailDo {
    s.DrugEntBaseInfoId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgListupoutBillUpOutDetailDo) SetProduceDateFormat(v string) *AlibabaAlihealthDrugtraceTopYljgListupoutBillUpOutDetailDo {
    s.ProduceDateFormat = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgListupoutBillUpOutDetailDo) SetStatus(v string) *AlibabaAlihealthDrugtraceTopYljgListupoutBillUpOutDetailDo {
    s.Status = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgListupoutBillUpOutDetailDo) SetToUserId(v string) *AlibabaAlihealthDrugtraceTopYljgListupoutBillUpOutDetailDo {
    s.ToUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgListupoutBillUpOutDetailDo) SetFromUserId(v string) *AlibabaAlihealthDrugtraceTopYljgListupoutBillUpOutDetailDo {
    s.FromUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgListupoutBillUpOutDetailDo) SetAssRefEntId(v string) *AlibabaAlihealthDrugtraceTopYljgListupoutBillUpOutDetailDo {
    s.AssRefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgListupoutBillUpOutDetailDo) SetAssEntId(v string) *AlibabaAlihealthDrugtraceTopYljgListupoutBillUpOutDetailDo {
    s.AssEntId = &v
    return s
}
