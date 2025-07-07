package request


type AlibabaAlihealthDrugMscSearchbillDetailRequest struct {
    /*
        单据号码     */
    BillCode  *string `json:"bill_code" required:"true" `
    /*
        企业refEntId     */
    RefEntId  *string `json:"ref_ent_id" required:"true" `
    /*
        货主/配送     */
    AuthRefUserId  *string `json:"auth_ref_user_id,omitempty" required:"false" `
}

func (s *AlibabaAlihealthDrugMscSearchbillDetailRequest) SetBillCode(v string) *AlibabaAlihealthDrugMscSearchbillDetailRequest {
    s.BillCode = &v
    return s
}
func (s *AlibabaAlihealthDrugMscSearchbillDetailRequest) SetRefEntId(v string) *AlibabaAlihealthDrugMscSearchbillDetailRequest {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugMscSearchbillDetailRequest) SetAuthRefUserId(v string) *AlibabaAlihealthDrugMscSearchbillDetailRequest {
    s.AuthRefUserId = &v
    return s
}

func (req *AlibabaAlihealthDrugMscSearchbillDetailRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.BillCode != nil) {
        paramMap["bill_code"] = *req.BillCode
    }
    if(req.RefEntId != nil) {
        paramMap["ref_ent_id"] = *req.RefEntId
    }
    if(req.AuthRefUserId != nil) {
        paramMap["auth_ref_user_id"] = *req.AuthRefUserId
    }
    return paramMap
}

func (req *AlibabaAlihealthDrugMscSearchbillDetailRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}