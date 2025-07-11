package request


type AlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeRequest struct {
    /*
        追溯码     */
    Code  *string `json:"code,omitempty" required:"false" `
    /*
        企业ID （一般为要查询单据的收货企业）     */
    RefEntId  *string `json:"ref_ent_id,omitempty" required:"false" `
}

func (s *AlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeRequest) SetCode(v string) *AlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeRequest {
    s.Code = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeRequest) SetRefEntId(v string) *AlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeRequest {
    s.RefEntId = &v
    return s
}

func (req *AlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Code != nil) {
        paramMap["code"] = *req.Code
    }
    if(req.RefEntId != nil) {
        paramMap["ref_ent_id"] = *req.RefEntId
    }
    return paramMap
}

func (req *AlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}