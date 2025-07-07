package request


type AlibabaAlihealthDrugtraceTopLsydQueryBillstatusRequest struct {
    /*
        企业ref_ent_id     */
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
        发货企业【注：是ent_id,不是ref_ent_id】     */
    FromUserId  *string `json:"from_user_id,omitempty" required:"false" `
    /*
        收货企业【注：是ent_id,不是ref_ent_id】     */
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

func (s *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusRequest) SetRefEntId(v string) *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusRequest {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusRequest) SetBeginDate(v string) *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusRequest {
    s.BeginDate = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusRequest) SetEndDate(v string) *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusRequest {
    s.EndDate = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusRequest) SetBillType(v string) *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusRequest {
    s.BillType = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusRequest) SetBillCode(v string) *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusRequest {
    s.BillCode = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusRequest) SetDrugType(v string) *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusRequest {
    s.DrugType = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusRequest) SetDealStatus(v string) *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusRequest {
    s.DealStatus = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusRequest) SetFromUserId(v string) *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusRequest {
    s.FromUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusRequest) SetToUserId(v string) *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusRequest {
    s.ToUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusRequest) SetAgentRefUserId(v string) *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusRequest {
    s.AgentRefUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusRequest) SetPageSize(v int64) *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusRequest {
    s.PageSize = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusRequest) SetPage(v int64) *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusRequest {
    s.Page = &v
    return s
}

func (req *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusRequest) ToMap() map[string]interface{} {
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

func (req *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}