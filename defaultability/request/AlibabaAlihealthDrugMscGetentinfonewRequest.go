package request

import (
	"github.com/LGYS07/topsdk/defaultability/domain"
	"github.com/LGYS07/topsdk/util"
)

type AlibabaAlihealthDrugMscGetentinfonewRequest struct {
	/*
	   refEntId     */
	RefEntId *string `json:"ref_ent_id" required:"true" `
	/*
	   查询企业信息参数     */
	QueryParam *domain.AlibabaAlihealthDrugMscGetentinfonewTopEntInfoReqDto `json:"query_param,omitempty" required:"false" `
}

func (s *AlibabaAlihealthDrugMscGetentinfonewRequest) SetRefEntId(v string) *AlibabaAlihealthDrugMscGetentinfonewRequest {
	s.RefEntId = &v
	return s
}
func (s *AlibabaAlihealthDrugMscGetentinfonewRequest) SetQueryParam(v domain.AlibabaAlihealthDrugMscGetentinfonewTopEntInfoReqDto) *AlibabaAlihealthDrugMscGetentinfonewRequest {
	s.QueryParam = &v
	return s
}

func (req *AlibabaAlihealthDrugMscGetentinfonewRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.RefEntId != nil {
		paramMap["ref_ent_id"] = *req.RefEntId
	}
	if req.QueryParam != nil {
		paramMap["query_param"] = util.ConvertStruct(*req.QueryParam)
	}
	return paramMap
}

func (req *AlibabaAlihealthDrugMscGetentinfonewRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
