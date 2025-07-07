package domain


type AlibabaAlihealthDrugtraceTopLsydListupoutBillUpOutDetailDo struct {
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
        收货企业     */
    ToUserName  *string `json:"to_user_name,omitempty" `

    /*
        发货企业     */
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

}

func (s *AlibabaAlihealthDrugtraceTopLsydListupoutBillUpOutDetailDo) SetFromEntName(v string) *AlibabaAlihealthDrugtraceTopLsydListupoutBillUpOutDetailDo {
    s.FromEntName = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydListupoutBillUpOutDetailDo) SetCodeCount(v int64) *AlibabaAlihealthDrugtraceTopLsydListupoutBillUpOutDetailDo {
    s.CodeCount = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydListupoutBillUpOutDetailDo) SetExprieDate(v string) *AlibabaAlihealthDrugtraceTopLsydListupoutBillUpOutDetailDo {
    s.ExprieDate = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydListupoutBillUpOutDetailDo) SetProduceEntName(v string) *AlibabaAlihealthDrugtraceTopLsydListupoutBillUpOutDetailDo {
    s.ProduceEntName = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydListupoutBillUpOutDetailDo) SetProduceDate(v string) *AlibabaAlihealthDrugtraceTopLsydListupoutBillUpOutDetailDo {
    s.ProduceDate = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydListupoutBillUpOutDetailDo) SetProduceBatchNo(v string) *AlibabaAlihealthDrugtraceTopLsydListupoutBillUpOutDetailDo {
    s.ProduceBatchNo = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydListupoutBillUpOutDetailDo) SetPkgSpec(v string) *AlibabaAlihealthDrugtraceTopLsydListupoutBillUpOutDetailDo {
    s.PkgSpec = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydListupoutBillUpOutDetailDo) SetPhysicInfo(v string) *AlibabaAlihealthDrugtraceTopLsydListupoutBillUpOutDetailDo {
    s.PhysicInfo = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydListupoutBillUpOutDetailDo) SetPhysicName(v string) *AlibabaAlihealthDrugtraceTopLsydListupoutBillUpOutDetailDo {
    s.PhysicName = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydListupoutBillUpOutDetailDo) SetPrepnCount(v int64) *AlibabaAlihealthDrugtraceTopLsydListupoutBillUpOutDetailDo {
    s.PrepnCount = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydListupoutBillUpOutDetailDo) SetFromRefUserId(v string) *AlibabaAlihealthDrugtraceTopLsydListupoutBillUpOutDetailDo {
    s.FromRefUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydListupoutBillUpOutDetailDo) SetToRefUserId(v string) *AlibabaAlihealthDrugtraceTopLsydListupoutBillUpOutDetailDo {
    s.ToRefUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydListupoutBillUpOutDetailDo) SetBillTime(v string) *AlibabaAlihealthDrugtraceTopLsydListupoutBillUpOutDetailDo {
    s.BillTime = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydListupoutBillUpOutDetailDo) SetBillCode(v string) *AlibabaAlihealthDrugtraceTopLsydListupoutBillUpOutDetailDo {
    s.BillCode = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydListupoutBillUpOutDetailDo) SetBillType(v string) *AlibabaAlihealthDrugtraceTopLsydListupoutBillUpOutDetailDo {
    s.BillType = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydListupoutBillUpOutDetailDo) SetToUserName(v string) *AlibabaAlihealthDrugtraceTopLsydListupoutBillUpOutDetailDo {
    s.ToUserName = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydListupoutBillUpOutDetailDo) SetFromUserName(v string) *AlibabaAlihealthDrugtraceTopLsydListupoutBillUpOutDetailDo {
    s.FromUserName = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydListupoutBillUpOutDetailDo) SetExprieDateFormat(v string) *AlibabaAlihealthDrugtraceTopLsydListupoutBillUpOutDetailDo {
    s.ExprieDateFormat = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydListupoutBillUpOutDetailDo) SetBillTimeFormat(v string) *AlibabaAlihealthDrugtraceTopLsydListupoutBillUpOutDetailDo {
    s.BillTimeFormat = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydListupoutBillUpOutDetailDo) SetBillOutId(v int64) *AlibabaAlihealthDrugtraceTopLsydListupoutBillUpOutDetailDo {
    s.BillOutId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydListupoutBillUpOutDetailDo) SetPrepnUnit(v int64) *AlibabaAlihealthDrugtraceTopLsydListupoutBillUpOutDetailDo {
    s.PrepnUnit = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydListupoutBillUpOutDetailDo) SetPrepnSpec(v string) *AlibabaAlihealthDrugtraceTopLsydListupoutBillUpOutDetailDo {
    s.PrepnSpec = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydListupoutBillUpOutDetailDo) SetDrugEntBaseInfoId(v string) *AlibabaAlihealthDrugtraceTopLsydListupoutBillUpOutDetailDo {
    s.DrugEntBaseInfoId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydListupoutBillUpOutDetailDo) SetProduceDateFormat(v string) *AlibabaAlihealthDrugtraceTopLsydListupoutBillUpOutDetailDo {
    s.ProduceDateFormat = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydListupoutBillUpOutDetailDo) SetStatus(v string) *AlibabaAlihealthDrugtraceTopLsydListupoutBillUpOutDetailDo {
    s.Status = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydListupoutBillUpOutDetailDo) SetToUserId(v string) *AlibabaAlihealthDrugtraceTopLsydListupoutBillUpOutDetailDo {
    s.ToUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydListupoutBillUpOutDetailDo) SetFromUserId(v string) *AlibabaAlihealthDrugtraceTopLsydListupoutBillUpOutDetailDo {
    s.FromUserId = &v
    return s
}
