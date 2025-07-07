package request


type AlibabaAlihealthDrugtraceTopLsydGetkeyflagdruginfoDownloadurlRequest struct {
    /*
        调用接口的企业ID     */
    RefEntId  *string `json:"ref_ent_id" required:"true" `
}

func (s *AlibabaAlihealthDrugtraceTopLsydGetkeyflagdruginfoDownloadurlRequest) SetRefEntId(v string) *AlibabaAlihealthDrugtraceTopLsydGetkeyflagdruginfoDownloadurlRequest {
    s.RefEntId = &v
    return s
}

func (req *AlibabaAlihealthDrugtraceTopLsydGetkeyflagdruginfoDownloadurlRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.RefEntId != nil) {
        paramMap["ref_ent_id"] = *req.RefEntId
    }
    return paramMap
}

func (req *AlibabaAlihealthDrugtraceTopLsydGetkeyflagdruginfoDownloadurlRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}