package request


type AlibabaAlihealthDrugtraceTopYljgQueryBillstatusRequest struct {
    /*
        企业ID     */
    RefEntId  *string `json:"ref_ent_id" required:"true" `
    /*
        开始日期（没有时分秒，【单据创建时间】）     */
    BeginDate  *string `json:"begin_date" required:"true" `
    /*
        结束日期（没有时分秒，【单据创建时间】）     */
    EndDate  *string `json:"end_date" required:"true" `
    /*
        单据类型 A：全部 AI：全部入库 AO：全部出库     */
    BillType  *string `json:"bill_type" required:"true" `
    /*
        单据号     */
    BillCode  *string `json:"bill_code,omitempty" required:"false" `
    /*
        药品类型     */
    DrugType  *string `json:"drug_type,omitempty" required:"false" `
    /*
        状态  0, 上传成功     3, 处理成功     4, 处理失败     */
    DealStatus  *string `json:"deal_status,omitempty" required:"false" `
    /*
        发货商     */
    FromUserId  *string `json:"from_user_id,omitempty" required:"false" `
    /*
        收货商     */
    ToUserId  *string `json:"to_user_id,omitempty" required:"false" `
    /*
        代理商     */
    AgentRefUserId  *string `json:"agent_ref_user_id,omitempty" required:"false" `
    /*
        页大小     */
    PageSize  *int64 `json:"page_size" required:"true" `
    /*
        页码     */
    Page  *int64 `json:"page" required:"true" `
}

func (s *AlibabaAlihealthDrugtraceTopYljgQueryBillstatusRequest) SetRefEntId(v string) *AlibabaAlihealthDrugtraceTopYljgQueryBillstatusRequest {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryBillstatusRequest) SetBeginDate(v string) *AlibabaAlihealthDrugtraceTopYljgQueryBillstatusRequest {
    s.BeginDate = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryBillstatusRequest) SetEndDate(v string) *AlibabaAlihealthDrugtraceTopYljgQueryBillstatusRequest {
    s.EndDate = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryBillstatusRequest) SetBillType(v string) *AlibabaAlihealthDrugtraceTopYljgQueryBillstatusRequest {
    s.BillType = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryBillstatusRequest) SetBillCode(v string) *AlibabaAlihealthDrugtraceTopYljgQueryBillstatusRequest {
    s.BillCode = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryBillstatusRequest) SetDrugType(v string) *AlibabaAlihealthDrugtraceTopYljgQueryBillstatusRequest {
    s.DrugType = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryBillstatusRequest) SetDealStatus(v string) *AlibabaAlihealthDrugtraceTopYljgQueryBillstatusRequest {
    s.DealStatus = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryBillstatusRequest) SetFromUserId(v string) *AlibabaAlihealthDrugtraceTopYljgQueryBillstatusRequest {
    s.FromUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryBillstatusRequest) SetToUserId(v string) *AlibabaAlihealthDrugtraceTopYljgQueryBillstatusRequest {
    s.ToUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryBillstatusRequest) SetAgentRefUserId(v string) *AlibabaAlihealthDrugtraceTopYljgQueryBillstatusRequest {
    s.AgentRefUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryBillstatusRequest) SetPageSize(v int64) *AlibabaAlihealthDrugtraceTopYljgQueryBillstatusRequest {
    s.PageSize = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryBillstatusRequest) SetPage(v int64) *AlibabaAlihealthDrugtraceTopYljgQueryBillstatusRequest {
    s.Page = &v
    return s
}

func (req *AlibabaAlihealthDrugtraceTopYljgQueryBillstatusRequest) ToMap() map[string]interface{} {
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
    if(req.BillType != nil) {
        paramMap["bill_type"] = *req.BillType
    }
    if(req.BillCode != nil) {
        paramMap["bill_code"] = *req.BillCode
    }
    if(req.DrugType != nil) {
        paramMap["drug_type"] = *req.DrugType
    }
    if(req.DealStatus != nil) {
        paramMap["deal_status"] = *req.DealStatus
    }
    if(req.FromUserId != nil) {
        paramMap["from_user_id"] = *req.FromUserId
    }
    if(req.ToUserId != nil) {
        paramMap["to_user_id"] = *req.ToUserId
    }
    if(req.AgentRefUserId != nil) {
        paramMap["agent_ref_user_id"] = *req.AgentRefUserId
    }
    if(req.PageSize != nil) {
        paramMap["page_size"] = *req.PageSize
    }
    if(req.Page != nil) {
        paramMap["page"] = *req.Page
    }
    return paramMap
}

func (req *AlibabaAlihealthDrugtraceTopYljgQueryBillstatusRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}