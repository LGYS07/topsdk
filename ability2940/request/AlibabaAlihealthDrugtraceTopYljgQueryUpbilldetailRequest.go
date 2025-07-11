package request


type AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailRequest struct {
    /*
        单据号码     */
    BillCode  *string `json:"bill_code" required:"true" `
    /*
        本企业refEntId     */
    RefEntId  *string `json:"ref_ent_id" required:"true" `
}

func (s *AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailRequest) SetBillCode(v string) *AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailRequest {
    s.BillCode = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailRequest) SetRefEntId(v string) *AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailRequest {
    s.RefEntId = &v
    return s
}

func (req *AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.BillCode != nil) {
        paramMap["bill_code"] = *req.BillCode
    }
    if(req.RefEntId != nil) {
        paramMap["ref_ent_id"] = *req.RefEntId
    }
    return paramMap
}

func (req *AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}