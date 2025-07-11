package request


type AlibabaAlihealthDrugtraceTopYljgQueryListpartsRequest struct {
    /*
        企业唯一标识     */
    RefEntId  *string `json:"ref_ent_id" required:"true" `
    /*
        企业名称     */
    EntName  *string `json:"ent_name,omitempty" required:"false" `
    /*
        企业自定义编号     */
    RefPartnerId  *string `json:"ref_partner_id,omitempty" required:"false" `
    /*
        开始时间     */
    BeginDate  *string `json:"begin_date,omitempty" required:"false" `
    /*
        结束时间     */
    EndDate  *string `json:"end_date,omitempty" required:"false" `
    /*
        页大小     */
    PageSize  *int64 `json:"page_size" required:"true" `
    /*
        页码     */
    Page  *int64 `json:"page" required:"true" `
}

func (s *AlibabaAlihealthDrugtraceTopYljgQueryListpartsRequest) SetRefEntId(v string) *AlibabaAlihealthDrugtraceTopYljgQueryListpartsRequest {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryListpartsRequest) SetEntName(v string) *AlibabaAlihealthDrugtraceTopYljgQueryListpartsRequest {
    s.EntName = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryListpartsRequest) SetRefPartnerId(v string) *AlibabaAlihealthDrugtraceTopYljgQueryListpartsRequest {
    s.RefPartnerId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryListpartsRequest) SetBeginDate(v string) *AlibabaAlihealthDrugtraceTopYljgQueryListpartsRequest {
    s.BeginDate = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryListpartsRequest) SetEndDate(v string) *AlibabaAlihealthDrugtraceTopYljgQueryListpartsRequest {
    s.EndDate = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryListpartsRequest) SetPageSize(v int64) *AlibabaAlihealthDrugtraceTopYljgQueryListpartsRequest {
    s.PageSize = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryListpartsRequest) SetPage(v int64) *AlibabaAlihealthDrugtraceTopYljgQueryListpartsRequest {
    s.Page = &v
    return s
}

func (req *AlibabaAlihealthDrugtraceTopYljgQueryListpartsRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.RefEntId != nil) {
        paramMap["ref_ent_id"] = *req.RefEntId
    }
    if(req.EntName != nil) {
        paramMap["ent_name"] = *req.EntName
    }
    if(req.RefPartnerId != nil) {
        paramMap["ref_partner_id"] = *req.RefPartnerId
    }
    if(req.BeginDate != nil) {
        paramMap["begin_date"] = *req.BeginDate
    }
    if(req.EndDate != nil) {
        paramMap["end_date"] = *req.EndDate
    }
    if(req.PageSize != nil) {
        paramMap["page_size"] = *req.PageSize
    }
    if(req.Page != nil) {
        paramMap["page"] = *req.Page
    }
    return paramMap
}

func (req *AlibabaAlihealthDrugtraceTopYljgQueryListpartsRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}