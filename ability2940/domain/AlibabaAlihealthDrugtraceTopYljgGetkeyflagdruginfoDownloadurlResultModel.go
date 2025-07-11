package domain


type AlibabaAlihealthDrugtraceTopYljgGetkeyflagdruginfoDownloadurlResultModel struct {
    /*
        true：接口调用成功     */
    Success  *bool `json:"success,omitempty" `

    /*
        返回值     */
    Model  *AlibabaAlihealthDrugtraceTopYljgGetkeyflagdruginfoDownloadurlJSONObject `json:"model,omitempty" `

    /*
        接口调用失败具体信息     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        接口调用失败具体code     */
    MsgCode  *string `json:"msg_code,omitempty" `

}

func (s *AlibabaAlihealthDrugtraceTopYljgGetkeyflagdruginfoDownloadurlResultModel) SetSuccess(v bool) *AlibabaAlihealthDrugtraceTopYljgGetkeyflagdruginfoDownloadurlResultModel {
    s.Success = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgGetkeyflagdruginfoDownloadurlResultModel) SetModel(v AlibabaAlihealthDrugtraceTopYljgGetkeyflagdruginfoDownloadurlJSONObject) *AlibabaAlihealthDrugtraceTopYljgGetkeyflagdruginfoDownloadurlResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgGetkeyflagdruginfoDownloadurlResultModel) SetMsgInfo(v string) *AlibabaAlihealthDrugtraceTopYljgGetkeyflagdruginfoDownloadurlResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgGetkeyflagdruginfoDownloadurlResultModel) SetMsgCode(v string) *AlibabaAlihealthDrugtraceTopYljgGetkeyflagdruginfoDownloadurlResultModel {
    s.MsgCode = &v
    return s
}
