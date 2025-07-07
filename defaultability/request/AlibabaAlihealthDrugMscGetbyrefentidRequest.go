package request


type AlibabaAlihealthDrugMscGetbyrefentidRequest struct {
    /*
        接口调用企业的唯一标识（接口调用者）     */
    RefEntId  *string `json:"ref_ent_id" required:"true" `
    /*
        准备要查询的企业唯一标识（返回该唯一标识企业的详细信息）     */
    DestRefEntId  *string `json:"dest_ref_ent_id" required:"true" `
}

func (s *AlibabaAlihealthDrugMscGetbyrefentidRequest) SetRefEntId(v string) *AlibabaAlihealthDrugMscGetbyrefentidRequest {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugMscGetbyrefentidRequest) SetDestRefEntId(v string) *AlibabaAlihealthDrugMscGetbyrefentidRequest {
    s.DestRefEntId = &v
    return s
}

func (req *AlibabaAlihealthDrugMscGetbyrefentidRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.RefEntId != nil) {
        paramMap["ref_ent_id"] = *req.RefEntId
    }
    if(req.DestRefEntId != nil) {
        paramMap["dest_ref_ent_id"] = *req.DestRefEntId
    }
    return paramMap
}

func (req *AlibabaAlihealthDrugMscGetbyrefentidRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}