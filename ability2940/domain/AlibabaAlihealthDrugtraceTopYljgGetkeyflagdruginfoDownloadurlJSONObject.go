package domain


type AlibabaAlihealthDrugtraceTopYljgGetkeyflagdruginfoDownloadurlJSONObject struct {
    /*
        文件下载地址     */
    Url  *string `json:"url,omitempty" `

}

func (s *AlibabaAlihealthDrugtraceTopYljgGetkeyflagdruginfoDownloadurlJSONObject) SetUrl(v string) *AlibabaAlihealthDrugtraceTopYljgGetkeyflagdruginfoDownloadurlJSONObject {
    s.Url = &v
    return s
}
