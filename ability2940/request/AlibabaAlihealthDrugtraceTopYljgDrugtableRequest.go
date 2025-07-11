package request


type AlibabaAlihealthDrugtraceTopYljgDrugtableRequest struct {
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
        开始日期     */
    StartDate  *string `json:"start_date,omitempty" required:"false" `
    /*
        结束日期     */
    EndDate  *string `json:"end_date,omitempty" required:"false" `
    /*
        企业名称     */
    EntName  *string `json:"ent_name,omitempty" required:"false" `
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

func (s *AlibabaAlihealthDrugtraceTopYljgDrugtableRequest) SetRefEntId(v string) *AlibabaAlihealthDrugtraceTopYljgDrugtableRequest {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgDrugtableRequest) SetPhysicName(v string) *AlibabaAlihealthDrugtraceTopYljgDrugtableRequest {
    s.PhysicName = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgDrugtableRequest) SetApprovalLicenceNo(v string) *AlibabaAlihealthDrugtraceTopYljgDrugtableRequest {
    s.ApprovalLicenceNo = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgDrugtableRequest) SetStartDate(v string) *AlibabaAlihealthDrugtraceTopYljgDrugtableRequest {
    s.StartDate = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgDrugtableRequest) SetEndDate(v string) *AlibabaAlihealthDrugtraceTopYljgDrugtableRequest {
    s.EndDate = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgDrugtableRequest) SetEntName(v string) *AlibabaAlihealthDrugtraceTopYljgDrugtableRequest {
    s.EntName = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgDrugtableRequest) SetPackageSpec(v string) *AlibabaAlihealthDrugtraceTopYljgDrugtableRequest {
    s.PackageSpec = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgDrugtableRequest) SetPrepnSpec(v string) *AlibabaAlihealthDrugtraceTopYljgDrugtableRequest {
    s.PrepnSpec = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgDrugtableRequest) SetPageSize(v int64) *AlibabaAlihealthDrugtraceTopYljgDrugtableRequest {
    s.PageSize = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgDrugtableRequest) SetPage(v int64) *AlibabaAlihealthDrugtraceTopYljgDrugtableRequest {
    s.Page = &v
    return s
}

func (req *AlibabaAlihealthDrugtraceTopYljgDrugtableRequest) ToMap() map[string]interface{} {
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
    if(req.StartDate != nil) {
        paramMap["start_date"] = *req.StartDate
    }
    if(req.EndDate != nil) {
        paramMap["end_date"] = *req.EndDate
    }
    if(req.EntName != nil) {
        paramMap["ent_name"] = *req.EntName
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

func (req *AlibabaAlihealthDrugtraceTopYljgDrugtableRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}