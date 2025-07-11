package request

import (
	"github.com/LGYS07/topsdk/ability2940/domain"
	"github.com/LGYS07/topsdk/util"
)

type AlibabaAlihealthDrugYljgGetentinfonewRequest struct {
	/*
	   refEntId     */
	RefEntId *string `json:"ref_ent_id" required:"true" `
	/*
	   查询企业信息参数     */
	QueryParam *domain.AlibabaAlihealthDrugYljgGetentinfonewTopEntInfoReqDto `json:"query_param,omitempty" required:"false" `
}

func (s *AlibabaAlihealthDrugYljgGetentinfonewRequest) SetRefEntId(v string) *AlibabaAlihealthDrugYljgGetentinfonewRequest {
	s.RefEntId = &v
	return s
}
func (s *AlibabaAlihealthDrugYljgGetentinfonewRequest) SetQueryParam(v domain.AlibabaAlihealthDrugYljgGetentinfonewTopEntInfoReqDto) *AlibabaAlihealthDrugYljgGetentinfonewRequest {
	s.QueryParam = &v
	return s
}

func (req *AlibabaAlihealthDrugYljgGetentinfonewRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.RefEntId != nil {
		paramMap["ref_ent_id"] = *req.RefEntId
	}
	if req.QueryParam != nil {
		paramMap["query_param"] = util.ConvertStruct(*req.QueryParam)
	}
	return paramMap
}

func (req *AlibabaAlihealthDrugYljgGetentinfonewRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
