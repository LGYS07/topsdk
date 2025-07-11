package request


type AlibabaAlihealthDrugtraceTopYljgQueryGetentinfoRequest struct {
    /*
        公司名称(全称)     */
    EntName  *string `json:"ent_name" required:"true" `
}

func (s *AlibabaAlihealthDrugtraceTopYljgQueryGetentinfoRequest) SetEntName(v string) *AlibabaAlihealthDrugtraceTopYljgQueryGetentinfoRequest {
    s.EntName = &v
    return s
}

func (req *AlibabaAlihealthDrugtraceTopYljgQueryGetentinfoRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.EntName != nil) {
        paramMap["ent_name"] = *req.EntName
    }
    return paramMap
}

func (req *AlibabaAlihealthDrugtraceTopYljgQueryGetentinfoRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}