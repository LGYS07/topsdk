package domain


type AlibabaAlihealthDrugtraceTopLsydGetkeyflagdruginfoDownloadurlJSONObject struct {
    /*
        文件下载地址     */
    Url  *string `json:"url,omitempty" `

}

func (s *AlibabaAlihealthDrugtraceTopLsydGetkeyflagdruginfoDownloadurlJSONObject) SetUrl(v string) *AlibabaAlihealthDrugtraceTopLsydGetkeyflagdruginfoDownloadurlJSONObject {
    s.Url = &v
    return s
}
