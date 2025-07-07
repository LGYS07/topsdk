package request

import (
	"github.com/LGYS07/topsdk/util"
)

type AlibabaAlihealthDrugtraceTopLsydQueryCodedetailRequest struct {
	/*
	   企业ref_ent_id      */
	RefEntId *string `json:"ref_ent_id" required:"true" `
	/*
	   码列表【多个码用逗号拼接的字符串。要求数量在1000个码以下，但一般不要传这么多，如果网络不好很容易传输一半报错】     */
	Codes *[]string `json:"codes" required:"true" `
}

func (s *AlibabaAlihealthDrugtraceTopLsydQueryCodedetailRequest) SetRefEntId(v string) *AlibabaAlihealthDrugtraceTopLsydQueryCodedetailRequest {
	s.RefEntId = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryCodedetailRequest) SetCodes(v []string) *AlibabaAlihealthDrugtraceTopLsydQueryCodedetailRequest {
	s.Codes = &v
	return s
}

func (req *AlibabaAlihealthDrugtraceTopLsydQueryCodedetailRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.RefEntId != nil {
		paramMap["ref_ent_id"] = *req.RefEntId
	}
	if req.Codes != nil {
		paramMap["codes"] = util.ConvertBasicList(*req.Codes)
	}
	return paramMap
}

func (req *AlibabaAlihealthDrugtraceTopLsydQueryCodedetailRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
