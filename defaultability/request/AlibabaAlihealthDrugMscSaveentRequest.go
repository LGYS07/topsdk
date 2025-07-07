package request

import (
	"github.com/LGYS07/topsdk/defaultability/domain"
	"github.com/LGYS07/topsdk/util"
)

type AlibabaAlihealthDrugMscSaveentRequest struct {
	/*
	   添加企业唯一标识     */
	RefEntId *string `json:"ref_ent_id" required:"true" `
	/*
	   新增企业信息     */
	AddEntReq *domain.AlibabaAlihealthDrugMscSaveentAddEntReqDto `json:"add_ent_req" required:"true" `
	/*
	   图片数据流。图片大小务必控制在2M以内     */
	LicPictureByte *[]byte `json:"lic_picture_byte" required:"true" `
}

func (s *AlibabaAlihealthDrugMscSaveentRequest) SetRefEntId(v string) *AlibabaAlihealthDrugMscSaveentRequest {
	s.RefEntId = &v
	return s
}
func (s *AlibabaAlihealthDrugMscSaveentRequest) SetAddEntReq(v domain.AlibabaAlihealthDrugMscSaveentAddEntReqDto) *AlibabaAlihealthDrugMscSaveentRequest {
	s.AddEntReq = &v
	return s
}
func (s *AlibabaAlihealthDrugMscSaveentRequest) SetLicPictureByte(v []byte) *AlibabaAlihealthDrugMscSaveentRequest {
	s.LicPictureByte = &v
	return s
}

func (req *AlibabaAlihealthDrugMscSaveentRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.RefEntId != nil {
		paramMap["ref_ent_id"] = *req.RefEntId
	}
	if req.AddEntReq != nil {
		paramMap["add_ent_req"] = util.ConvertStruct(*req.AddEntReq)
	}
	return paramMap
}

func (req *AlibabaAlihealthDrugMscSaveentRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	if req.LicPictureByte != nil {
		fileMap["lic_picture_byte"] = *req.LicPictureByte
	}
	return fileMap
}
