package request


type AlibabaAlihealthDrugtraceTopYljgQueryRelationRequest struct {
    /*
        接口调用企业的唯一标识（接口调用者）     */
    RefEntId  *string `json:"ref_ent_id" required:"true" `
    /*
        追溯码,多个码需要逗号拼接，最大10个码     */
    Code  *string `json:"code" required:"true" `
    /*
        目标企业唯一标识（为哪个企业查询，一般与入参ref_ent_id一样）     */
    DesRefEntId  *string `json:"des_ref_ent_id" required:"true" `
}

func (s *AlibabaAlihealthDrugtraceTopYljgQueryRelationRequest) SetRefEntId(v string) *AlibabaAlihealthDrugtraceTopYljgQueryRelationRequest {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryRelationRequest) SetCode(v string) *AlibabaAlihealthDrugtraceTopYljgQueryRelationRequest {
    s.Code = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryRelationRequest) SetDesRefEntId(v string) *AlibabaAlihealthDrugtraceTopYljgQueryRelationRequest {
    s.DesRefEntId = &v
    return s
}

func (req *AlibabaAlihealthDrugtraceTopYljgQueryRelationRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.RefEntId != nil) {
        paramMap["ref_ent_id"] = *req.RefEntId
    }
    if(req.Code != nil) {
        paramMap["code"] = *req.Code
    }
    if(req.DesRefEntId != nil) {
        paramMap["des_ref_ent_id"] = *req.DesRefEntId
    }
    return paramMap
}

func (req *AlibabaAlihealthDrugtraceTopYljgQueryRelationRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}