package request

import (
        "topsdk/ability2939/domain"
        "topsdk/util"
    )

type AlibabaAlihealthDrugLsydSaveentRequest struct {
    /*
        添加企业唯一标识     */
    RefEntId  *string `json:"ref_ent_id" required:"true" `
    /*
        新增企业信息     */
    AddEntReq  *domain.AlibabaAlihealthDrugLsydSaveentAddEntReqDto `json:"add_ent_req" required:"true" `
    /*
        图片数据流。图片大小务必控制在2M以内     */
    LicPictureByte  *[]byte `json:"lic_picture_byte" required:"true" `
}

func (s *AlibabaAlihealthDrugLsydSaveentRequest) SetRefEntId(v string) *AlibabaAlihealthDrugLsydSaveentRequest {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydSaveentRequest) SetAddEntReq(v domain.AlibabaAlihealthDrugLsydSaveentAddEntReqDto) *AlibabaAlihealthDrugLsydSaveentRequest {
    s.AddEntReq = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydSaveentRequest) SetLicPictureByte(v []byte) *AlibabaAlihealthDrugLsydSaveentRequest {
    s.LicPictureByte = &v
    return s
}

func (req *AlibabaAlihealthDrugLsydSaveentRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.RefEntId != nil) {
        paramMap["ref_ent_id"] = *req.RefEntId
    }
    if(req.AddEntReq != nil) {
        paramMap["add_ent_req"] = util.ConvertStruct(*req.AddEntReq)
    }
    return paramMap
}

func (req *AlibabaAlihealthDrugLsydSaveentRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    if req.LicPictureByte != nil {
        fileMap["lic_picture_byte"] = *req.LicPictureByte
    }
    return fileMap
}