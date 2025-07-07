package request


type AlibabaAlihealthDrugMscGetbyentidRequest struct {
    /*
        接口调用企业的唯一标识（接口调用者）     */
    RefEntId  *string `json:"ref_ent_id" required:"true" `
    /*
        准备要查询的企业ID（返回该企业ID的详细信息）     */
    EntId  *string `json:"ent_id" required:"true" `
}

func (s *AlibabaAlihealthDrugMscGetbyentidRequest) SetRefEntId(v string) *AlibabaAlihealthDrugMscGetbyentidRequest {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugMscGetbyentidRequest) SetEntId(v string) *AlibabaAlihealthDrugMscGetbyentidRequest {
    s.EntId = &v
    return s
}

func (req *AlibabaAlihealthDrugMscGetbyentidRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.RefEntId != nil) {
        paramMap["ref_ent_id"] = *req.RefEntId
    }
    if(req.EntId != nil) {
        paramMap["ent_id"] = *req.EntId
    }
    return paramMap
}

func (req *AlibabaAlihealthDrugMscGetbyentidRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}