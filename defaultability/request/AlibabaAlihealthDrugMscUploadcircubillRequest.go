package request

import (
        "topsdk/util"
    )

type AlibabaAlihealthDrugMscUploadcircubillRequest struct {
    /*
        单据编号（唯一）     */
    BillCode  *string `json:"bill_code" required:"true" `
    /*
        单据时间（一般为药品入出库时间）     */
    BillTime  *util.LocalTime `json:"bill_time" required:"true" `
    /*
        单据类型【102代表采购入库,201代表销售出库，其它单据类型详见文档】     */
    BillType  *int64 `json:"bill_type" required:"true" `
    /*
        上传企业（单据的所有者，上传人）【注意：该入参是ref_ent_id，不是ent_id】,可通过“根据企业名称查询企业唯一标识”接口获取     */
    RefUserId  *string `json:"ref_user_id" required:"true" `
    /*
        发货企业【注意：该入参是ent_id,并不是ref_ent_id】     */
    FromUserId  *string `json:"from_user_id" required:"true" `
    /*
        收货企业【注意：该入参是ent_id,并不是ref_ent_id】     */
    ToUserId  *string `json:"to_user_id" required:"true" `
    /*
        直调企业【注意：该入参是ent_id,并不是ref_ent_id】     */
    DestUserId  *string `json:"dest_user_id,omitempty" required:"false" `
    /*
        操作人标识（写appkey编号）     */
    OperIcCode  *string `json:"oper_ic_code" required:"true" `
    /*
        单据提交者姓名     */
    OperIcName  *string `json:"oper_ic_name,omitempty" required:"false" `
    /*
        单据文件体【bas64字符串】，看对接文档中的代码示例，示例中有相应说明。     */
    FileContent  *string `json:"file_content" required:"true" `
    /*
        单据名称     */
    UploadFileName  *string `json:"upload_file_name" required:"true" `
    /*
        请求端类型【暂定都写2】     */
    ClientType  *int64 `json:"client_type" required:"true" `
    /*
        发货地址     */
    FromAddress  *string `json:"from_address,omitempty" required:"false" `
    /*
        收货地址     */
    ToAddress  *string `json:"to_address,omitempty" required:"false" `
    /*
        发货单编号     */
    FromBillCode  *string `json:"from_bill_code,omitempty" required:"false" `
    /*
        订货单编号     */
    OrderCode  *string `json:"order_code,omitempty" required:"false" `
    /*
        发货人(特药出单据必填)     */
    FromPerson  *string `json:"from_person,omitempty" required:"false" `
    /*
        收货人(特药入单据必填)     */
    ToPerson  *string `json:"to_person,omitempty" required:"false" `
    /*
        药品配送企业refentid【出库单填写，与dis_ent_id入参选其一添加】     */
    DisRefEntId  *string `json:"dis_ref_ent_id,omitempty" required:"false" `
    /*
        药品配送企业entId【出库单填写】     */
    DisEntId  *string `json:"dis_ent_id,omitempty" required:"false" `
    /*
        应收货总数量     */
    QuReceivable  *int64 `json:"qu_receivable,omitempty" required:"false" `
    /*
        未验证通过原因【验证未通过时填写】     */
    XtIsCheck  *string `json:"xt_is_check,omitempty" required:"false" `
    /*
        是否验证，0：未通过验证，1：已验证     */
    XtCheckCode  *string `json:"xt_check_code,omitempty" required:"false" `
    /*
        未验证通过原因描述【验证未通过时填写】     */
    XtCheckCodeDesc  *string `json:"xt_check_code_desc,omitempty" required:"false" `
    /*
        【可不填】药品列表Json："codeCount":         药品数量 "commDrugId":     国家药品唯一标识 "exprieDate":         生产日期 "physicInfo":          药品信息 "pkgSpec":           包状规格 "prepnCount":       制剂数量 "produceBatchNo":生产批次 "produceDate":      生产日期     */
    DrugListJson  *string `json:"drug_list_json,omitempty" required:"false" `
    /*
        单据委托企业refEntId     */
    AssRefEntId  *string `json:"ass_ref_ent_id,omitempty" required:"false" `
    /*
        单据委托企业entId     */
    AssEntId  *string `json:"ass_ent_id,omitempty" required:"false" `
    /*
        忽略校验项，1: 代表忽略码重复上传校验   0：代表需要进行码重复上传校验     */
    IgnoreCheck  *string `json:"ignore_check,omitempty" required:"false" `
}

func (s *AlibabaAlihealthDrugMscUploadcircubillRequest) SetBillCode(v string) *AlibabaAlihealthDrugMscUploadcircubillRequest {
    s.BillCode = &v
    return s
}
func (s *AlibabaAlihealthDrugMscUploadcircubillRequest) SetBillTime(v util.LocalTime) *AlibabaAlihealthDrugMscUploadcircubillRequest {
    s.BillTime = &v
    return s
}
func (s *AlibabaAlihealthDrugMscUploadcircubillRequest) SetBillType(v int64) *AlibabaAlihealthDrugMscUploadcircubillRequest {
    s.BillType = &v
    return s
}
func (s *AlibabaAlihealthDrugMscUploadcircubillRequest) SetRefUserId(v string) *AlibabaAlihealthDrugMscUploadcircubillRequest {
    s.RefUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugMscUploadcircubillRequest) SetFromUserId(v string) *AlibabaAlihealthDrugMscUploadcircubillRequest {
    s.FromUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugMscUploadcircubillRequest) SetToUserId(v string) *AlibabaAlihealthDrugMscUploadcircubillRequest {
    s.ToUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugMscUploadcircubillRequest) SetDestUserId(v string) *AlibabaAlihealthDrugMscUploadcircubillRequest {
    s.DestUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugMscUploadcircubillRequest) SetOperIcCode(v string) *AlibabaAlihealthDrugMscUploadcircubillRequest {
    s.OperIcCode = &v
    return s
}
func (s *AlibabaAlihealthDrugMscUploadcircubillRequest) SetOperIcName(v string) *AlibabaAlihealthDrugMscUploadcircubillRequest {
    s.OperIcName = &v
    return s
}
func (s *AlibabaAlihealthDrugMscUploadcircubillRequest) SetFileContent(v string) *AlibabaAlihealthDrugMscUploadcircubillRequest {
    s.FileContent = &v
    return s
}
func (s *AlibabaAlihealthDrugMscUploadcircubillRequest) SetUploadFileName(v string) *AlibabaAlihealthDrugMscUploadcircubillRequest {
    s.UploadFileName = &v
    return s
}
func (s *AlibabaAlihealthDrugMscUploadcircubillRequest) SetClientType(v int64) *AlibabaAlihealthDrugMscUploadcircubillRequest {
    s.ClientType = &v
    return s
}
func (s *AlibabaAlihealthDrugMscUploadcircubillRequest) SetFromAddress(v string) *AlibabaAlihealthDrugMscUploadcircubillRequest {
    s.FromAddress = &v
    return s
}
func (s *AlibabaAlihealthDrugMscUploadcircubillRequest) SetToAddress(v string) *AlibabaAlihealthDrugMscUploadcircubillRequest {
    s.ToAddress = &v
    return s
}
func (s *AlibabaAlihealthDrugMscUploadcircubillRequest) SetFromBillCode(v string) *AlibabaAlihealthDrugMscUploadcircubillRequest {
    s.FromBillCode = &v
    return s
}
func (s *AlibabaAlihealthDrugMscUploadcircubillRequest) SetOrderCode(v string) *AlibabaAlihealthDrugMscUploadcircubillRequest {
    s.OrderCode = &v
    return s
}
func (s *AlibabaAlihealthDrugMscUploadcircubillRequest) SetFromPerson(v string) *AlibabaAlihealthDrugMscUploadcircubillRequest {
    s.FromPerson = &v
    return s
}
func (s *AlibabaAlihealthDrugMscUploadcircubillRequest) SetToPerson(v string) *AlibabaAlihealthDrugMscUploadcircubillRequest {
    s.ToPerson = &v
    return s
}
func (s *AlibabaAlihealthDrugMscUploadcircubillRequest) SetDisRefEntId(v string) *AlibabaAlihealthDrugMscUploadcircubillRequest {
    s.DisRefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugMscUploadcircubillRequest) SetDisEntId(v string) *AlibabaAlihealthDrugMscUploadcircubillRequest {
    s.DisEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugMscUploadcircubillRequest) SetQuReceivable(v int64) *AlibabaAlihealthDrugMscUploadcircubillRequest {
    s.QuReceivable = &v
    return s
}
func (s *AlibabaAlihealthDrugMscUploadcircubillRequest) SetXtIsCheck(v string) *AlibabaAlihealthDrugMscUploadcircubillRequest {
    s.XtIsCheck = &v
    return s
}
func (s *AlibabaAlihealthDrugMscUploadcircubillRequest) SetXtCheckCode(v string) *AlibabaAlihealthDrugMscUploadcircubillRequest {
    s.XtCheckCode = &v
    return s
}
func (s *AlibabaAlihealthDrugMscUploadcircubillRequest) SetXtCheckCodeDesc(v string) *AlibabaAlihealthDrugMscUploadcircubillRequest {
    s.XtCheckCodeDesc = &v
    return s
}
func (s *AlibabaAlihealthDrugMscUploadcircubillRequest) SetDrugListJson(v string) *AlibabaAlihealthDrugMscUploadcircubillRequest {
    s.DrugListJson = &v
    return s
}
func (s *AlibabaAlihealthDrugMscUploadcircubillRequest) SetAssRefEntId(v string) *AlibabaAlihealthDrugMscUploadcircubillRequest {
    s.AssRefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugMscUploadcircubillRequest) SetAssEntId(v string) *AlibabaAlihealthDrugMscUploadcircubillRequest {
    s.AssEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugMscUploadcircubillRequest) SetIgnoreCheck(v string) *AlibabaAlihealthDrugMscUploadcircubillRequest {
    s.IgnoreCheck = &v
    return s
}

func (req *AlibabaAlihealthDrugMscUploadcircubillRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.BillCode != nil) {
        paramMap["bill_code"] = *req.BillCode
    }
    if(req.BillTime != nil) {
        paramMap["bill_time"] = *req.BillTime
    }
    if(req.BillType != nil) {
        paramMap["bill_type"] = *req.BillType
    }
    if(req.RefUserId != nil) {
        paramMap["ref_user_id"] = *req.RefUserId
    }
    if(req.FromUserId != nil) {
        paramMap["from_user_id"] = *req.FromUserId
    }
    if(req.ToUserId != nil) {
        paramMap["to_user_id"] = *req.ToUserId
    }
    if(req.DestUserId != nil) {
        paramMap["dest_user_id"] = *req.DestUserId
    }
    if(req.OperIcCode != nil) {
        paramMap["oper_ic_code"] = *req.OperIcCode
    }
    if(req.OperIcName != nil) {
        paramMap["oper_ic_name"] = *req.OperIcName
    }
    if(req.FileContent != nil) {
        paramMap["file_content"] = *req.FileContent
    }
    if(req.UploadFileName != nil) {
        paramMap["upload_file_name"] = *req.UploadFileName
    }
    if(req.ClientType != nil) {
        paramMap["client_type"] = *req.ClientType
    }
    if(req.FromAddress != nil) {
        paramMap["from_address"] = *req.FromAddress
    }
    if(req.ToAddress != nil) {
        paramMap["to_address"] = *req.ToAddress
    }
    if(req.FromBillCode != nil) {
        paramMap["from_bill_code"] = *req.FromBillCode
    }
    if(req.OrderCode != nil) {
        paramMap["order_code"] = *req.OrderCode
    }
    if(req.FromPerson != nil) {
        paramMap["from_person"] = *req.FromPerson
    }
    if(req.ToPerson != nil) {
        paramMap["to_person"] = *req.ToPerson
    }
    if(req.DisRefEntId != nil) {
        paramMap["dis_ref_ent_id"] = *req.DisRefEntId
    }
    if(req.DisEntId != nil) {
        paramMap["dis_ent_id"] = *req.DisEntId
    }
    if(req.QuReceivable != nil) {
        paramMap["qu_receivable"] = *req.QuReceivable
    }
    if(req.XtIsCheck != nil) {
        paramMap["xt_is_check"] = *req.XtIsCheck
    }
    if(req.XtCheckCode != nil) {
        paramMap["xt_check_code"] = *req.XtCheckCode
    }
    if(req.XtCheckCodeDesc != nil) {
        paramMap["xt_check_code_desc"] = *req.XtCheckCodeDesc
    }
    if(req.DrugListJson != nil) {
        paramMap["drug_list_json"] = *req.DrugListJson
    }
    if(req.AssRefEntId != nil) {
        paramMap["ass_ref_ent_id"] = *req.AssRefEntId
    }
    if(req.AssEntId != nil) {
        paramMap["ass_ent_id"] = *req.AssEntId
    }
    if(req.IgnoreCheck != nil) {
        paramMap["ignore_check"] = *req.IgnoreCheck
    }
    return paramMap
}

func (req *AlibabaAlihealthDrugMscUploadcircubillRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}