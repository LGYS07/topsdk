package request


type AlibabaAlihealthDrugMscBilloutDetailwithcodesRequest struct {
    /*
        单据号码     */
    BillCode  *string `json:"bill_code" required:"true" `
    /*
        企业refEntId     */
    RefEntId  *string `json:"ref_ent_id" required:"true" `
}

func (s *AlibabaAlihealthDrugMscBilloutDetailwithcodesRequest) SetBillCode(v string) *AlibabaAlihealthDrugMscBilloutDetailwithcodesRequest {
    s.BillCode = &v
    return s
}
func (s *AlibabaAlihealthDrugMscBilloutDetailwithcodesRequest) SetRefEntId(v string) *AlibabaAlihealthDrugMscBilloutDetailwithcodesRequest {
    s.RefEntId = &v
    return s
}

func (req *AlibabaAlihealthDrugMscBilloutDetailwithcodesRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.BillCode != nil) {
        paramMap["bill_code"] = *req.BillCode
    }
    if(req.RefEntId != nil) {
        paramMap["ref_ent_id"] = *req.RefEntId
    }
    return paramMap
}

func (req *AlibabaAlihealthDrugMscBilloutDetailwithcodesRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}