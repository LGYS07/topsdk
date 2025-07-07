package request


type AlibabaAlihealthDrugtraceTopLsydListupoutRequest struct {
    /*
        企业ID     */
    RefEntId  *string `json:"ref_ent_id" required:"true" `
    /*
        开始日期（不写时分秒）     */
    BeginDate  *string `json:"begin_date" required:"true" `
    /*
        结束日期（不写时分秒）     */
    EndDate  *string `json:"end_date" required:"true" `
    /*
        发货单位     */
    FromUserId  *string `json:"from_user_id,omitempty" required:"false" `
    /*
        生产批号     */
    ProduceBatchNo  *string `json:"produce_batch_no,omitempty" required:"false" `
    /*
        药品ID     */
    DrugEntBaseInfoId  *string `json:"drug_ent_base_info_id,omitempty" required:"false" `
    /*
        单据类型     */
    BillType  *string `json:"bill_type,omitempty" required:"false" `
    /*
        药品类型     */
    PhysicType  *string `json:"physic_type,omitempty" required:"false" `
    /*
        状态     */
    Status  *string `json:"status,omitempty" required:"false" `
    /*
        单据号     */
    BillCode  *string `json:"bill_code,omitempty" required:"false" `
    /*
        页大小     */
    PageSize  *int64 `json:"page_size" required:"true" `
    /*
        页码     */
    Page  *int64 `json:"page" required:"true" `
}

func (s *AlibabaAlihealthDrugtraceTopLsydListupoutRequest) SetRefEntId(v string) *AlibabaAlihealthDrugtraceTopLsydListupoutRequest {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydListupoutRequest) SetBeginDate(v string) *AlibabaAlihealthDrugtraceTopLsydListupoutRequest {
    s.BeginDate = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydListupoutRequest) SetEndDate(v string) *AlibabaAlihealthDrugtraceTopLsydListupoutRequest {
    s.EndDate = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydListupoutRequest) SetFromUserId(v string) *AlibabaAlihealthDrugtraceTopLsydListupoutRequest {
    s.FromUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydListupoutRequest) SetProduceBatchNo(v string) *AlibabaAlihealthDrugtraceTopLsydListupoutRequest {
    s.ProduceBatchNo = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydListupoutRequest) SetDrugEntBaseInfoId(v string) *AlibabaAlihealthDrugtraceTopLsydListupoutRequest {
    s.DrugEntBaseInfoId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydListupoutRequest) SetBillType(v string) *AlibabaAlihealthDrugtraceTopLsydListupoutRequest {
    s.BillType = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydListupoutRequest) SetPhysicType(v string) *AlibabaAlihealthDrugtraceTopLsydListupoutRequest {
    s.PhysicType = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydListupoutRequest) SetStatus(v string) *AlibabaAlihealthDrugtraceTopLsydListupoutRequest {
    s.Status = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydListupoutRequest) SetBillCode(v string) *AlibabaAlihealthDrugtraceTopLsydListupoutRequest {
    s.BillCode = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydListupoutRequest) SetPageSize(v int64) *AlibabaAlihealthDrugtraceTopLsydListupoutRequest {
    s.PageSize = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydListupoutRequest) SetPage(v int64) *AlibabaAlihealthDrugtraceTopLsydListupoutRequest {
    s.Page = &v
    return s
}

func (req *AlibabaAlihealthDrugtraceTopLsydListupoutRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.RefEntId != nil) {
        paramMap["ref_ent_id"] = *req.RefEntId
    }
    if(req.BeginDate != nil) {
        paramMap["begin_date"] = *req.BeginDate
    }
    if(req.EndDate != nil) {
        paramMap["end_date"] = *req.EndDate
    }
    if(req.FromUserId != nil) {
        paramMap["from_user_id"] = *req.FromUserId
    }
    if(req.ProduceBatchNo != nil) {
        paramMap["produce_batch_no"] = *req.ProduceBatchNo
    }
    if(req.DrugEntBaseInfoId != nil) {
        paramMap["drug_ent_base_info_id"] = *req.DrugEntBaseInfoId
    }
    if(req.BillType != nil) {
        paramMap["bill_type"] = *req.BillType
    }
    if(req.PhysicType != nil) {
        paramMap["physic_type"] = *req.PhysicType
    }
    if(req.Status != nil) {
        paramMap["status"] = *req.Status
    }
    if(req.BillCode != nil) {
        paramMap["bill_code"] = *req.BillCode
    }
    if(req.PageSize != nil) {
        paramMap["page_size"] = *req.PageSize
    }
    if(req.Page != nil) {
        paramMap["page"] = *req.Page
    }
    return paramMap
}

func (req *AlibabaAlihealthDrugtraceTopLsydListupoutRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}