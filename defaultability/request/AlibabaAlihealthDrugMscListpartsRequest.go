package request


type AlibabaAlihealthDrugMscListpartsRequest struct {
    /*
        调用者企业唯一标识     */
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

func (s *AlibabaAlihealthDrugMscListpartsRequest) SetRefEntId(v string) *AlibabaAlihealthDrugMscListpartsRequest {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListpartsRequest) SetEntName(v string) *AlibabaAlihealthDrugMscListpartsRequest {
    s.EntName = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListpartsRequest) SetRefPartnerId(v string) *AlibabaAlihealthDrugMscListpartsRequest {
    s.RefPartnerId = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListpartsRequest) SetBeginDate(v string) *AlibabaAlihealthDrugMscListpartsRequest {
    s.BeginDate = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListpartsRequest) SetEndDate(v string) *AlibabaAlihealthDrugMscListpartsRequest {
    s.EndDate = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListpartsRequest) SetPageSize(v int64) *AlibabaAlihealthDrugMscListpartsRequest {
    s.PageSize = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListpartsRequest) SetPage(v int64) *AlibabaAlihealthDrugMscListpartsRequest {
    s.Page = &v
    return s
}

func (req *AlibabaAlihealthDrugMscListpartsRequest) ToMap() map[string]interface{} {
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

func (req *AlibabaAlihealthDrugMscListpartsRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}