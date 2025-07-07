package request


type AlibabaAlihealthDrugMscGetentinfoRequest struct {
    /*
        调用者企业唯一标识     */
    RefEntId  *string `json:"ref_ent_id" required:"true" `
    /*
        企业名称     */
    EntName  *string `json:"ent_name" required:"true" `
}

func (s *AlibabaAlihealthDrugMscGetentinfoRequest) SetRefEntId(v string) *AlibabaAlihealthDrugMscGetentinfoRequest {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugMscGetentinfoRequest) SetEntName(v string) *AlibabaAlihealthDrugMscGetentinfoRequest {
    s.EntName = &v
    return s
}

func (req *AlibabaAlihealthDrugMscGetentinfoRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.RefEntId != nil) {
        paramMap["ref_ent_id"] = *req.RefEntId
    }
    if(req.EntName != nil) {
        paramMap["ent_name"] = *req.EntName
    }
    return paramMap
}

func (req *AlibabaAlihealthDrugMscGetentinfoRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}