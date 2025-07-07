package request


type AlibabaAlihealthDrugMscDrugtableRequest struct {
    /*
        调用企业id     */
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
        生产企业名字     */
    ProduceEntName  *string `json:"produce_ent_name,omitempty" required:"false" `
    /*
        页码     */
    PageSize  *int64 `json:"page_size" required:"true" `
    /*
        页面大小     */
    Page  *int64 `json:"page" required:"true" `
}

func (s *AlibabaAlihealthDrugMscDrugtableRequest) SetRefEntId(v string) *AlibabaAlihealthDrugMscDrugtableRequest {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugMscDrugtableRequest) SetPhysicName(v string) *AlibabaAlihealthDrugMscDrugtableRequest {
    s.PhysicName = &v
    return s
}
func (s *AlibabaAlihealthDrugMscDrugtableRequest) SetApprovalLicenceNo(v string) *AlibabaAlihealthDrugMscDrugtableRequest {
    s.ApprovalLicenceNo = &v
    return s
}
func (s *AlibabaAlihealthDrugMscDrugtableRequest) SetPackageSpec(v string) *AlibabaAlihealthDrugMscDrugtableRequest {
    s.PackageSpec = &v
    return s
}
func (s *AlibabaAlihealthDrugMscDrugtableRequest) SetPrepnSpec(v string) *AlibabaAlihealthDrugMscDrugtableRequest {
    s.PrepnSpec = &v
    return s
}
func (s *AlibabaAlihealthDrugMscDrugtableRequest) SetProduceEntName(v string) *AlibabaAlihealthDrugMscDrugtableRequest {
    s.ProduceEntName = &v
    return s
}
func (s *AlibabaAlihealthDrugMscDrugtableRequest) SetPageSize(v int64) *AlibabaAlihealthDrugMscDrugtableRequest {
    s.PageSize = &v
    return s
}
func (s *AlibabaAlihealthDrugMscDrugtableRequest) SetPage(v int64) *AlibabaAlihealthDrugMscDrugtableRequest {
    s.Page = &v
    return s
}

func (req *AlibabaAlihealthDrugMscDrugtableRequest) ToMap() map[string]interface{} {
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
    if(req.ProduceEntName != nil) {
        paramMap["produce_ent_name"] = *req.ProduceEntName
    }
    if(req.PageSize != nil) {
        paramMap["page_size"] = *req.PageSize
    }
    if(req.Page != nil) {
        paramMap["page"] = *req.Page
    }
    return paramMap
}

func (req *AlibabaAlihealthDrugMscDrugtableRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}