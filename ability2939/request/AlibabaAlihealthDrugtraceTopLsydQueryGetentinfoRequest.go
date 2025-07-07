package request


type AlibabaAlihealthDrugtraceTopLsydQueryGetentinfoRequest struct {
    /*
        公司名称(全称)     */
    EntName  *string `json:"ent_name" required:"true" `
}

func (s *AlibabaAlihealthDrugtraceTopLsydQueryGetentinfoRequest) SetEntName(v string) *AlibabaAlihealthDrugtraceTopLsydQueryGetentinfoRequest {
    s.EntName = &v
    return s
}

func (req *AlibabaAlihealthDrugtraceTopLsydQueryGetentinfoRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.EntName != nil) {
        paramMap["ent_name"] = *req.EntName
    }
    return paramMap
}

func (req *AlibabaAlihealthDrugtraceTopLsydQueryGetentinfoRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}