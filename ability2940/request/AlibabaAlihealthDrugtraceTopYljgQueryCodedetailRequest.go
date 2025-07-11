package request

import (
	"github.com/LGYS07/topsdk/util"
)

type AlibabaAlihealthDrugtraceTopYljgQueryCodedetailRequest struct {
	/*
	   企业唯一标识     */
	RefEntId *string `json:"ref_ent_id" required:"true" `
	/*
	   码列表【多个码用逗号拼接的字符串。要求数量在1000个码以下，但一般不要传这么多，如果网络不好很容易传输一半报错】     */
	Codes *[]string `json:"codes" required:"true" `
}

func (s *AlibabaAlihealthDrugtraceTopYljgQueryCodedetailRequest) SetRefEntId(v string) *AlibabaAlihealthDrugtraceTopYljgQueryCodedetailRequest {
	s.RefEntId = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryCodedetailRequest) SetCodes(v []string) *AlibabaAlihealthDrugtraceTopYljgQueryCodedetailRequest {
	s.Codes = &v
	return s
}

func (req *AlibabaAlihealthDrugtraceTopYljgQueryCodedetailRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.RefEntId != nil {
		paramMap["ref_ent_id"] = *req.RefEntId
	}
	if req.Codes != nil {
		paramMap["codes"] = util.ConvertBasicList(*req.Codes)
	}
	return paramMap
}

func (req *AlibabaAlihealthDrugtraceTopYljgQueryCodedetailRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
