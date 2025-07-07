package request


type AlibabaAlihealthDrugtraceTopLsydQueryRelationRequest struct {
    /*
        接口调用企业的唯一标识（接口调用者）     */
    RefEntId  *string `json:"ref_ent_id" required:"true" `
    /*
        追溯码     */
    Code  *string `json:"code" required:"true" `
    /*
        目标企业唯一标识（为哪个企业查询，一般与入参ref_ent_id一样）     */
    DesRefEntId  *string `json:"des_ref_ent_id" required:"true" `
}

func (s *AlibabaAlihealthDrugtraceTopLsydQueryRelationRequest) SetRefEntId(v string) *AlibabaAlihealthDrugtraceTopLsydQueryRelationRequest {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryRelationRequest) SetCode(v string) *AlibabaAlihealthDrugtraceTopLsydQueryRelationRequest {
    s.Code = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryRelationRequest) SetDesRefEntId(v string) *AlibabaAlihealthDrugtraceTopLsydQueryRelationRequest {
    s.DesRefEntId = &v
    return s
}

func (req *AlibabaAlihealthDrugtraceTopLsydQueryRelationRequest) ToMap() map[string]interface{} {
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

func (req *AlibabaAlihealthDrugtraceTopLsydQueryRelationRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}