package request


type AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailRequest struct {
    /*
        单据号码     */
    BillCode  *string `json:"bill_code" required:"true" `
    /*
        本企业refEntId     */
    RefEntId  *string `json:"ref_ent_id" required:"true" `
}

func (s *AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailRequest) SetBillCode(v string) *AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailRequest {
    s.BillCode = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailRequest) SetRefEntId(v string) *AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailRequest {
    s.RefEntId = &v
    return s
}

func (req *AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.BillCode != nil) {
        paramMap["bill_code"] = *req.BillCode
    }
    if(req.RefEntId != nil) {
        paramMap["ref_ent_id"] = *req.RefEntId
    }
    return paramMap
}

func (req *AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}