package request


type AlibabaAlihealthDrugMscSearchbillRequest struct {
    /*
        企业标识     */
    RefEntId  *string `json:"ref_ent_id" required:"true" `
    /*
        货主（物流企业代货主查询时填写）     */
    AuthRefUserId  *string `json:"auth_ref_user_id,omitempty" required:"false" `
    /*
        开始日期     */
    BeginDate  *string `json:"begin_date" required:"true" `
    /*
        结束日期     */
    EndDate  *string `json:"end_date" required:"true" `
    /*
        发货企业     */
    PartnerIdSend  *string `json:"partner_id_send,omitempty" required:"false" `
    /*
        收货企业     */
    PartnerIdRecv  *string `json:"partner_id_recv,omitempty" required:"false" `
    /*
        单据号码     */
    BillCode  *string `json:"bill_code,omitempty" required:"false" `
    /*
        单据类型  A : 所有  AI :入库    AO:出库     */
    BillType  *string `json:"bill_type" required:"true" `
    /*
        当前页     */
    CurPage  *int64 `json:"cur_page" required:"true" `
    /*
        页大小     */
    PageSize  *int64 `json:"page_size" required:"true" `
}

func (s *AlibabaAlihealthDrugMscSearchbillRequest) SetRefEntId(v string) *AlibabaAlihealthDrugMscSearchbillRequest {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugMscSearchbillRequest) SetAuthRefUserId(v string) *AlibabaAlihealthDrugMscSearchbillRequest {
    s.AuthRefUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugMscSearchbillRequest) SetBeginDate(v string) *AlibabaAlihealthDrugMscSearchbillRequest {
    s.BeginDate = &v
    return s
}
func (s *AlibabaAlihealthDrugMscSearchbillRequest) SetEndDate(v string) *AlibabaAlihealthDrugMscSearchbillRequest {
    s.EndDate = &v
    return s
}
func (s *AlibabaAlihealthDrugMscSearchbillRequest) SetPartnerIdSend(v string) *AlibabaAlihealthDrugMscSearchbillRequest {
    s.PartnerIdSend = &v
    return s
}
func (s *AlibabaAlihealthDrugMscSearchbillRequest) SetPartnerIdRecv(v string) *AlibabaAlihealthDrugMscSearchbillRequest {
    s.PartnerIdRecv = &v
    return s
}
func (s *AlibabaAlihealthDrugMscSearchbillRequest) SetBillCode(v string) *AlibabaAlihealthDrugMscSearchbillRequest {
    s.BillCode = &v
    return s
}
func (s *AlibabaAlihealthDrugMscSearchbillRequest) SetBillType(v string) *AlibabaAlihealthDrugMscSearchbillRequest {
    s.BillType = &v
    return s
}
func (s *AlibabaAlihealthDrugMscSearchbillRequest) SetCurPage(v int64) *AlibabaAlihealthDrugMscSearchbillRequest {
    s.CurPage = &v
    return s
}
func (s *AlibabaAlihealthDrugMscSearchbillRequest) SetPageSize(v int64) *AlibabaAlihealthDrugMscSearchbillRequest {
    s.PageSize = &v
    return s
}

func (req *AlibabaAlihealthDrugMscSearchbillRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.RefEntId != nil) {
        paramMap["ref_ent_id"] = *req.RefEntId
    }
    if(req.AuthRefUserId != nil) {
        paramMap["auth_ref_user_id"] = *req.AuthRefUserId
    }
    if(req.BeginDate != nil) {
        paramMap["begin_date"] = *req.BeginDate
    }
    if(req.EndDate != nil) {
        paramMap["end_date"] = *req.EndDate
    }
    if(req.PartnerIdSend != nil) {
        paramMap["partner_id_send"] = *req.PartnerIdSend
    }
    if(req.PartnerIdRecv != nil) {
        paramMap["partner_id_recv"] = *req.PartnerIdRecv
    }
    if(req.BillCode != nil) {
        paramMap["bill_code"] = *req.BillCode
    }
    if(req.BillType != nil) {
        paramMap["bill_type"] = *req.BillType
    }
    if(req.CurPage != nil) {
        paramMap["cur_page"] = *req.CurPage
    }
    if(req.PageSize != nil) {
        paramMap["page_size"] = *req.PageSize
    }
    return paramMap
}

func (req *AlibabaAlihealthDrugMscSearchbillRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}