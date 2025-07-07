package domain


type AlibabaAlihealthDrugtraceTopLsydGetkeyflagdruginfoDownloadurlResultModel struct {
    /*
        true：接口调用成功     */
    Success  *bool `json:"success,omitempty" `

    /*
        返回值     */
    Model  *AlibabaAlihealthDrugtraceTopLsydGetkeyflagdruginfoDownloadurlJSONObject `json:"model,omitempty" `

    /*
        接口调用失败具体信息     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        接口调用失败具体code     */
    MsgCode  *string `json:"msg_code,omitempty" `

}

func (s *AlibabaAlihealthDrugtraceTopLsydGetkeyflagdruginfoDownloadurlResultModel) SetSuccess(v bool) *AlibabaAlihealthDrugtraceTopLsydGetkeyflagdruginfoDownloadurlResultModel {
    s.Success = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydGetkeyflagdruginfoDownloadurlResultModel) SetModel(v AlibabaAlihealthDrugtraceTopLsydGetkeyflagdruginfoDownloadurlJSONObject) *AlibabaAlihealthDrugtraceTopLsydGetkeyflagdruginfoDownloadurlResultModel {
    s.Model = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydGetkeyflagdruginfoDownloadurlResultModel) SetMsgInfo(v string) *AlibabaAlihealthDrugtraceTopLsydGetkeyflagdruginfoDownloadurlResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydGetkeyflagdruginfoDownloadurlResultModel) SetMsgCode(v string) *AlibabaAlihealthDrugtraceTopLsydGetkeyflagdruginfoDownloadurlResultModel {
    s.MsgCode = &v
    return s
}
