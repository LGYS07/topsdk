package request

import (
        "topsdk/ability2939/domain"
        "topsdk/util"
    )

type AlibabaAlihealthDrugLsydGetentinfonewRequest struct {
    /*
        refEntId     */
    RefEntId  *string `json:"ref_ent_id" required:"true" `
    /*
        查询企业信息参数     */
    QueryParam  *domain.AlibabaAlihealthDrugLsydGetentinfonewTopEntInfoReqDto `json:"query_param,omitempty" required:"false" `
}

func (s *AlibabaAlihealthDrugLsydGetentinfonewRequest) SetRefEntId(v string) *AlibabaAlihealthDrugLsydGetentinfonewRequest {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydGetentinfonewRequest) SetQueryParam(v domain.AlibabaAlihealthDrugLsydGetentinfonewTopEntInfoReqDto) *AlibabaAlihealthDrugLsydGetentinfonewRequest {
    s.QueryParam = &v
    return s
}

func (req *AlibabaAlihealthDrugLsydGetentinfonewRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.RefEntId != nil) {
        paramMap["ref_ent_id"] = *req.RefEntId
    }
    if(req.QueryParam != nil) {
        paramMap["query_param"] = util.ConvertStruct(*req.QueryParam)
    }
    return paramMap
}

func (req *AlibabaAlihealthDrugLsydGetentinfonewRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}