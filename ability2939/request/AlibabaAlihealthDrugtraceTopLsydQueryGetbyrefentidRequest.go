package request


type AlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidRequest struct {
    /*
        接口调用企业的唯一标识（接口调用者）     */
    RefEntId  *string `json:"ref_ent_id" required:"true" `
    /*
        准备要查询的企业唯一标识（返回该唯一标识企业的详细信息）     */
    DestRefEntId  *string `json:"dest_ref_ent_id" required:"true" `
}

func (s *AlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidRequest) SetRefEntId(v string) *AlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidRequest {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidRequest) SetDestRefEntId(v string) *AlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidRequest {
    s.DestRefEntId = &v
    return s
}

func (req *AlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.RefEntId != nil) {
        paramMap["ref_ent_id"] = *req.RefEntId
    }
    if(req.DestRefEntId != nil) {
        paramMap["dest_ref_ent_id"] = *req.DestRefEntId
    }
    return paramMap
}

func (req *AlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}