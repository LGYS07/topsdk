package request


type AlibabaAlihealthDrugMscBillinDetailwithcodeRequest struct {
    /*
        单据号码     */
    BillCode  *string `json:"bill_code" required:"true" `
    /*
        企业refEntId     */
    RefEntId  *string `json:"ref_ent_id" required:"true" `
}

func (s *AlibabaAlihealthDrugMscBillinDetailwithcodeRequest) SetBillCode(v string) *AlibabaAlihealthDrugMscBillinDetailwithcodeRequest {
    s.BillCode = &v
    return s
}
func (s *AlibabaAlihealthDrugMscBillinDetailwithcodeRequest) SetRefEntId(v string) *AlibabaAlihealthDrugMscBillinDetailwithcodeRequest {
    s.RefEntId = &v
    return s
}

func (req *AlibabaAlihealthDrugMscBillinDetailwithcodeRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.BillCode != nil) {
        paramMap["bill_code"] = *req.BillCode
    }
    if(req.RefEntId != nil) {
        paramMap["ref_ent_id"] = *req.RefEntId
    }
    return paramMap
}

func (req *AlibabaAlihealthDrugMscBillinDetailwithcodeRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}