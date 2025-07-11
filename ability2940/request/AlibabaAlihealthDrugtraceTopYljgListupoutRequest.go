package request


type AlibabaAlihealthDrugtraceTopYljgListupoutRequest struct {
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

func (s *AlibabaAlihealthDrugtraceTopYljgListupoutRequest) SetRefEntId(v string) *AlibabaAlihealthDrugtraceTopYljgListupoutRequest {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgListupoutRequest) SetBeginDate(v string) *AlibabaAlihealthDrugtraceTopYljgListupoutRequest {
    s.BeginDate = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgListupoutRequest) SetEndDate(v string) *AlibabaAlihealthDrugtraceTopYljgListupoutRequest {
    s.EndDate = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgListupoutRequest) SetFromUserId(v string) *AlibabaAlihealthDrugtraceTopYljgListupoutRequest {
    s.FromUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgListupoutRequest) SetProduceBatchNo(v string) *AlibabaAlihealthDrugtraceTopYljgListupoutRequest {
    s.ProduceBatchNo = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgListupoutRequest) SetDrugEntBaseInfoId(v string) *AlibabaAlihealthDrugtraceTopYljgListupoutRequest {
    s.DrugEntBaseInfoId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgListupoutRequest) SetBillType(v string) *AlibabaAlihealthDrugtraceTopYljgListupoutRequest {
    s.BillType = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgListupoutRequest) SetPhysicType(v string) *AlibabaAlihealthDrugtraceTopYljgListupoutRequest {
    s.PhysicType = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgListupoutRequest) SetStatus(v string) *AlibabaAlihealthDrugtraceTopYljgListupoutRequest {
    s.Status = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgListupoutRequest) SetBillCode(v string) *AlibabaAlihealthDrugtraceTopYljgListupoutRequest {
    s.BillCode = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgListupoutRequest) SetPageSize(v int64) *AlibabaAlihealthDrugtraceTopYljgListupoutRequest {
    s.PageSize = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgListupoutRequest) SetPage(v int64) *AlibabaAlihealthDrugtraceTopYljgListupoutRequest {
    s.Page = &v
    return s
}

func (req *AlibabaAlihealthDrugtraceTopYljgListupoutRequest) ToMap() map[string]interface{} {
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

func (req *AlibabaAlihealthDrugtraceTopYljgListupoutRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}