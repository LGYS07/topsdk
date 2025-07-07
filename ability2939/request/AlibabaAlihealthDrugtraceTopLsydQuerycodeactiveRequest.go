package request

import (
	"github.com/LGYS07/topsdk/util"
)

type AlibabaAlihealthDrugtraceTopLsydQuerycodeactiveRequest struct {
	/*
	   企业唯一标识     */
	RefEntId *string `json:"ref_ent_id" required:"true" `
	/*
	   码列表【多个码时用逗号拼接的字符串。要求数量在4000个码以下，但一般不要传这么多，如果网络不好很容易传输一半报错】     */
	Codes *[]string `json:"codes" required:"true" `
}

func (s *AlibabaAlihealthDrugtraceTopLsydQuerycodeactiveRequest) SetRefEntId(v string) *AlibabaAlihealthDrugtraceTopLsydQuerycodeactiveRequest {
	s.RefEntId = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQuerycodeactiveRequest) SetCodes(v []string) *AlibabaAlihealthDrugtraceTopLsydQuerycodeactiveRequest {
	s.Codes = &v
	return s
}

func (req *AlibabaAlihealthDrugtraceTopLsydQuerycodeactiveRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.RefEntId != nil {
		paramMap["ref_ent_id"] = *req.RefEntId
	}
	if req.Codes != nil {
		paramMap["codes"] = util.ConvertBasicList(*req.Codes)
	}
	return paramMap
}

func (req *AlibabaAlihealthDrugtraceTopLsydQuerycodeactiveRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
