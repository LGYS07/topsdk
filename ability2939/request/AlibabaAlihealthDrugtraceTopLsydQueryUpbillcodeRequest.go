package request


type AlibabaAlihealthDrugtraceTopLsydQueryUpbillcodeRequest struct {
    /*
        追溯码     */
    Code  *string `json:"code,omitempty" required:"false" `
    /*
        企业REF_ENT_ID （当前企业的唯一标识）     */
    RefEntId  *string `json:"ref_ent_id,omitempty" required:"false" `
}

func (s *AlibabaAlihealthDrugtraceTopLsydQueryUpbillcodeRequest) SetCode(v string) *AlibabaAlihealthDrugtraceTopLsydQueryUpbillcodeRequest {
    s.Code = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryUpbillcodeRequest) SetRefEntId(v string) *AlibabaAlihealthDrugtraceTopLsydQueryUpbillcodeRequest {
    s.RefEntId = &v
    return s
}

func (req *AlibabaAlihealthDrugtraceTopLsydQueryUpbillcodeRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Code != nil) {
        paramMap["code"] = *req.Code
    }
    if(req.RefEntId != nil) {
        paramMap["ref_ent_id"] = *req.RefEntId
    }
    return paramMap
}

func (req *AlibabaAlihealthDrugtraceTopLsydQueryUpbillcodeRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}