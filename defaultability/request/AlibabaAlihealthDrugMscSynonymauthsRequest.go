package request


type AlibabaAlihealthDrugMscSynonymauthsRequest struct {
    /*
        调用企业id     */
    RefEntId  *string `json:"ref_ent_id" required:"true" `
    /*
        企业名称     */
    EntName  *string `json:"ent_name,omitempty" required:"false" `
    /*
        货主自定义编号     */
    SynOwnEntId  *string `json:"syn_own_ent_id,omitempty" required:"false" `
    /*
        页码     */
    PageSize  *int64 `json:"page_size" required:"true" `
    /*
        页面大小     */
    Page  *int64 `json:"page" required:"true" `
}

func (s *AlibabaAlihealthDrugMscSynonymauthsRequest) SetRefEntId(v string) *AlibabaAlihealthDrugMscSynonymauthsRequest {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugMscSynonymauthsRequest) SetEntName(v string) *AlibabaAlihealthDrugMscSynonymauthsRequest {
    s.EntName = &v
    return s
}
func (s *AlibabaAlihealthDrugMscSynonymauthsRequest) SetSynOwnEntId(v string) *AlibabaAlihealthDrugMscSynonymauthsRequest {
    s.SynOwnEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugMscSynonymauthsRequest) SetPageSize(v int64) *AlibabaAlihealthDrugMscSynonymauthsRequest {
    s.PageSize = &v
    return s
}
func (s *AlibabaAlihealthDrugMscSynonymauthsRequest) SetPage(v int64) *AlibabaAlihealthDrugMscSynonymauthsRequest {
    s.Page = &v
    return s
}

func (req *AlibabaAlihealthDrugMscSynonymauthsRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.RefEntId != nil) {
        paramMap["ref_ent_id"] = *req.RefEntId
    }
    if(req.EntName != nil) {
        paramMap["ent_name"] = *req.EntName
    }
    if(req.SynOwnEntId != nil) {
        paramMap["syn_own_ent_id"] = *req.SynOwnEntId
    }
    if(req.PageSize != nil) {
        paramMap["page_size"] = *req.PageSize
    }
    if(req.Page != nil) {
        paramMap["page"] = *req.Page
    }
    return paramMap
}

func (req *AlibabaAlihealthDrugMscSynonymauthsRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}