package request

import (
	"github.com/LGYS07/topsdk/ability2940/domain"
	"github.com/LGYS07/topsdk/util"
)

type AlibabaAlihealthDrugYljgSaveentRequest struct {
	/*
	   添加企业唯一标识     */
	RefEntId *string `json:"ref_ent_id" required:"true" `
	/*
	   新增企业信息     */
	AddEntReq *domain.AlibabaAlihealthDrugYljgSaveentAddEntReqDto `json:"add_ent_req" required:"true" `
	/*
	   图片数据流。图片大小务必控制在2M以内     */
	LicPictureByte *[]byte `json:"lic_picture_byte" required:"true" `
}

func (s *AlibabaAlihealthDrugYljgSaveentRequest) SetRefEntId(v string) *AlibabaAlihealthDrugYljgSaveentRequest {
	s.RefEntId = &v
	return s
}
func (s *AlibabaAlihealthDrugYljgSaveentRequest) SetAddEntReq(v domain.AlibabaAlihealthDrugYljgSaveentAddEntReqDto) *AlibabaAlihealthDrugYljgSaveentRequest {
	s.AddEntReq = &v
	return s
}
func (s *AlibabaAlihealthDrugYljgSaveentRequest) SetLicPictureByte(v []byte) *AlibabaAlihealthDrugYljgSaveentRequest {
	s.LicPictureByte = &v
	return s
}

func (req *AlibabaAlihealthDrugYljgSaveentRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.RefEntId != nil {
		paramMap["ref_ent_id"] = *req.RefEntId
	}
	if req.AddEntReq != nil {
		paramMap["add_ent_req"] = util.ConvertStruct(*req.AddEntReq)
	}
	return paramMap
}

func (req *AlibabaAlihealthDrugYljgSaveentRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	if req.LicPictureByte != nil {
		fileMap["lic_picture_byte"] = *req.LicPictureByte
	}
	return fileMap
}
