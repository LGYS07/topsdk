package request


type AlibabaAlihealthDrugMscBillSearchstatusRequest struct {
    /*
        企业ref_ent_id（货主企业的ref_ent_id）     */
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
        单据号（精确值，不支持模糊查询）     */
    BillCode  *string `json:"bill_code,omitempty" required:"false" `
    /*
        药品类型     */
    DrugType  *string `json:"drug_type,omitempty" required:"false" `
    /*
        状态  0, 处理中     3, 处理成功     4, 处理失败     */
    DealStatus  *string `json:"deal_status,omitempty" required:"false" `
    /*
        发货商     */
    FromUserId  *string `json:"from_user_id,omitempty" required:"false" `
    /*
        收货商     */
    ToUserId  *string `json:"to_user_id,omitempty" required:"false" `
    /*
        代理商（第三方物流企业）     */
    AgentRefUserId  *string `json:"agent_ref_user_id,omitempty" required:"false" `
    /*
        页大小     */
    PageSize  *int64 `json:"page_size" required:"true" `
    /*
        页码     */
    Page  *int64 `json:"page" required:"true" `
}

func (s *AlibabaAlihealthDrugMscBillSearchstatusRequest) SetRefEntId(v string) *AlibabaAlihealthDrugMscBillSearchstatusRequest {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugMscBillSearchstatusRequest) SetBeginDate(v string) *AlibabaAlihealthDrugMscBillSearchstatusRequest {
    s.BeginDate = &v
    return s
}
func (s *AlibabaAlihealthDrugMscBillSearchstatusRequest) SetEndDate(v string) *AlibabaAlihealthDrugMscBillSearchstatusRequest {
    s.EndDate = &v
    return s
}
func (s *AlibabaAlihealthDrugMscBillSearchstatusRequest) SetBillType(v string) *AlibabaAlihealthDrugMscBillSearchstatusRequest {
    s.BillType = &v
    return s
}
func (s *AlibabaAlihealthDrugMscBillSearchstatusRequest) SetBillCode(v string) *AlibabaAlihealthDrugMscBillSearchstatusRequest {
    s.BillCode = &v
    return s
}
func (s *AlibabaAlihealthDrugMscBillSearchstatusRequest) SetDrugType(v string) *AlibabaAlihealthDrugMscBillSearchstatusRequest {
    s.DrugType = &v
    return s
}
func (s *AlibabaAlihealthDrugMscBillSearchstatusRequest) SetDealStatus(v string) *AlibabaAlihealthDrugMscBillSearchstatusRequest {
    s.DealStatus = &v
    return s
}
func (s *AlibabaAlihealthDrugMscBillSearchstatusRequest) SetFromUserId(v string) *AlibabaAlihealthDrugMscBillSearchstatusRequest {
    s.FromUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugMscBillSearchstatusRequest) SetToUserId(v string) *AlibabaAlihealthDrugMscBillSearchstatusRequest {
    s.ToUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugMscBillSearchstatusRequest) SetAgentRefUserId(v string) *AlibabaAlihealthDrugMscBillSearchstatusRequest {
    s.AgentRefUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugMscBillSearchstatusRequest) SetPageSize(v int64) *AlibabaAlihealthDrugMscBillSearchstatusRequest {
    s.PageSize = &v
    return s
}
func (s *AlibabaAlihealthDrugMscBillSearchstatusRequest) SetPage(v int64) *AlibabaAlihealthDrugMscBillSearchstatusRequest {
    s.Page = &v
    return s
}

func (req *AlibabaAlihealthDrugMscBillSearchstatusRequest) ToMap() map[string]interface{} {
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

func (req *AlibabaAlihealthDrugMscBillSearchstatusRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}