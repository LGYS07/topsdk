package request


type AlibabaAlihealthDrugtraceTopLsydServiceGetenddateRequest struct {
    /*
        调用接口的企业ID     */
    RefEntId  *string `json:"ref_ent_id" required:"true" `
    /*
        药 行业线：传 1      */
    Business  *int64 `json:"business" required:"true" `
    /*
        基础版：传 11     */
    Service  *int64 `json:"service" required:"true" `
}

func (s *AlibabaAlihealthDrugtraceTopLsydServiceGetenddateRequest) SetRefEntId(v string) *AlibabaAlihealthDrugtraceTopLsydServiceGetenddateRequest {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydServiceGetenddateRequest) SetBusiness(v int64) *AlibabaAlihealthDrugtraceTopLsydServiceGetenddateRequest {
    s.Business = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydServiceGetenddateRequest) SetService(v int64) *AlibabaAlihealthDrugtraceTopLsydServiceGetenddateRequest {
    s.Service = &v
    return s
}

func (req *AlibabaAlihealthDrugtraceTopLsydServiceGetenddateRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.RefEntId != nil) {
        paramMap["ref_ent_id"] = *req.RefEntId
    }
    if(req.Business != nil) {
        paramMap["business"] = *req.Business
    }
    if(req.Service != nil) {
        paramMap["service"] = *req.Service
    }
    return paramMap
}

func (req *AlibabaAlihealthDrugtraceTopLsydServiceGetenddateRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}