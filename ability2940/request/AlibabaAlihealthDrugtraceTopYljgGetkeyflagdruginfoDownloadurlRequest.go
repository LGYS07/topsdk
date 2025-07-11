package request


type AlibabaAlihealthDrugtraceTopYljgGetkeyflagdruginfoDownloadurlRequest struct {
    /*
        调用接口的企业ID     */
    RefEntId  *string `json:"ref_ent_id" required:"true" `
}

func (s *AlibabaAlihealthDrugtraceTopYljgGetkeyflagdruginfoDownloadurlRequest) SetRefEntId(v string) *AlibabaAlihealthDrugtraceTopYljgGetkeyflagdruginfoDownloadurlRequest {
    s.RefEntId = &v
    return s
}

func (req *AlibabaAlihealthDrugtraceTopYljgGetkeyflagdruginfoDownloadurlRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.RefEntId != nil) {
        paramMap["ref_ent_id"] = *req.RefEntId
    }
    return paramMap
}

func (req *AlibabaAlihealthDrugtraceTopYljgGetkeyflagdruginfoDownloadurlRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}