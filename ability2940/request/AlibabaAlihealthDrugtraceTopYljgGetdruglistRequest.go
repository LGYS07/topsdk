package request


type AlibabaAlihealthDrugtraceTopYljgGetdruglistRequest struct {
    /*
        企业ID     */
    RefEntId  *string `json:"ref_ent_id" required:"true" `
    /*
        药品通用名     */
    PhysicName  *string `json:"physic_name" required:"true" `
    /*
        批准文号     */
    ApprovalLicenceNo  *string `json:"approval_licence_no,omitempty" required:"false" `
    /*
        包装规格     */
    PackageSpec  *string `json:"package_spec,omitempty" required:"false" `
    /*
        制剂规格     */
    PrepnSpec  *string `json:"prepn_spec,omitempty" required:"false" `
    /*
        页大小     */
    PageSize  *int64 `json:"page_size" required:"true" `
    /*
        页码     */
    Page  *int64 `json:"page" required:"true" `
}

func (s *AlibabaAlihealthDrugtraceTopYljgGetdruglistRequest) SetRefEntId(v string) *AlibabaAlihealthDrugtraceTopYljgGetdruglistRequest {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgGetdruglistRequest) SetPhysicName(v string) *AlibabaAlihealthDrugtraceTopYljgGetdruglistRequest {
    s.PhysicName = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgGetdruglistRequest) SetApprovalLicenceNo(v string) *AlibabaAlihealthDrugtraceTopYljgGetdruglistRequest {
    s.ApprovalLicenceNo = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgGetdruglistRequest) SetPackageSpec(v string) *AlibabaAlihealthDrugtraceTopYljgGetdruglistRequest {
    s.PackageSpec = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgGetdruglistRequest) SetPrepnSpec(v string) *AlibabaAlihealthDrugtraceTopYljgGetdruglistRequest {
    s.PrepnSpec = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgGetdruglistRequest) SetPageSize(v int64) *AlibabaAlihealthDrugtraceTopYljgGetdruglistRequest {
    s.PageSize = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgGetdruglistRequest) SetPage(v int64) *AlibabaAlihealthDrugtraceTopYljgGetdruglistRequest {
    s.Page = &v
    return s
}

func (req *AlibabaAlihealthDrugtraceTopYljgGetdruglistRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.RefEntId != nil) {
        paramMap["ref_ent_id"] = *req.RefEntId
    }
    if(req.PhysicName != nil) {
        paramMap["physic_name"] = *req.PhysicName
    }
    if(req.ApprovalLicenceNo != nil) {
        paramMap["approval_licence_no"] = *req.ApprovalLicenceNo
    }
    if(req.PackageSpec != nil) {
        paramMap["package_spec"] = *req.PackageSpec
    }
    if(req.PrepnSpec != nil) {
        paramMap["prepn_spec"] = *req.PrepnSpec
    }
    if(req.PageSize != nil) {
        paramMap["page_size"] = *req.PageSize
    }
    if(req.Page != nil) {
        paramMap["page"] = *req.Page
    }
    return paramMap
}

func (req *AlibabaAlihealthDrugtraceTopYljgGetdruglistRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}