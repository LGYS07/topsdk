package request

import (
        "topsdk/util"
    )

type AlibabaAlihealthDrugMscUploadinoutbillRequest struct {
    /*
        单据编号【同一个企业不能上传相同单据号】     */
    BillCode  *string `json:"bill_code" required:"true" `
    /*
        单据时间（扫码时间）     */
    BillTime  *util.LocalTime `json:"bill_time" required:"true" `
    /*
        单据类型【102代表采购入库,201代表销售出库，其它单据类型详见文档】     */
    BillType  *int64 `json:"bill_type" required:"true" `
    /*
        药品类型【3普药2特药】89开头的码定义为特药，其它码定义成普药【可以随便填写，单据上传后会以实际为准】     */
    PhysicType  *int64 `json:"physic_type" required:"true" `
    /*
        货主（单据的所有者，上传人），上传企业的单位维一编码【注意：该入参是ref_ent_id，不是ent_id】     */
    RefUserId  *string `json:"ref_user_id" required:"true" `
    /*
        不要填值。该字段是为了兼容之前接口逻辑。     */
    AgentRefUserId  *string `json:"agent_ref_user_id,omitempty" required:"false" `
    /*
        发货企业entId【注意：该入参是ent_id,并不是ref_ent_id】     */
    FromUserId  *string `json:"from_user_id" required:"true" `
    /*
        收货企业entId【注意：该入参是ent_id,并不是ref_ent_id】     */
    ToUserId  *string `json:"to_user_id" required:"true" `
    /*
        直调企业标识【注意：该入参是ent_id,并不是ref_ent_id】     */
    DestUserId  *string `json:"dest_user_id,omitempty" required:"false" `
    /*
        单据提交者（调用接口时的appkey编号）     */
    OperIcCode  *string `json:"oper_ic_code" required:"true" `
    /*
        单据提交者姓名（出入库单上传人的名子）     */
    OperIcName  *string `json:"oper_ic_name" required:"true" `
    /*
        仓号     */
    WarehouseId  *string `json:"warehouse_id,omitempty" required:"false" `
    /*
        药品ID[企业自已系统的药品ID]     */
    DrugId  *string `json:"drug_id,omitempty" required:"false" `
    /*
        追溯码【多个码时用逗号拼接的字符串。要求数量在3500个码以下，但一般不要传这么多，如果网络不好很容易传输一半报错】     */
    TraceCodes  *[]string `json:"trace_codes" required:"true" `
    /*
        调用方类型[必须填2]     */
    ClientType  *string `json:"client_type" required:"true" `
    /*
        退货原因代码[退货入出库时填写] （1:破损 2:召回 3:滞销 4:过期失效 5:近效期 6:其他）     */
    ReturnReasonCode  *string `json:"return_reason_code,omitempty" required:"false" `
    /*
        退货原因描述[退货入出库时填写]     */
    ReturnReasonDes  *string `json:"return_reason_des,omitempty" required:"false" `
    /*
        注销原因代码【销毁出库时填写】（1:破损 2:霉变 3:过期失效 4:其他）     */
    CancelReasonCode  *string `json:"cancel_reason_code,omitempty" required:"false" `
    /*
        注销原因描述【销毁出库时填写】     */
    CancelReasonDes  *string `json:"cancel_reason_des,omitempty" required:"false" `
    /*
        执行人姓名【销毁出库时填写】     */
    ExecuterName  *string `json:"executer_name,omitempty" required:"false" `
    /*
        执行人证件号【销毁出库时填写】     */
    ExecuterCode  *string `json:"executer_code,omitempty" required:"false" `
    /*
        监督人姓名【销毁出库时填写】     */
    SuperviserName  *string `json:"superviser_name,omitempty" required:"false" `
    /*
        监督人证件号【销毁出库时填写】     */
    SuperviserCode  *string `json:"superviser_code,omitempty" required:"false" `
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
        订货单编号【可选】     */
    OrderCode  *string `json:"order_code,omitempty" required:"false" `
    /*
        发货人     */
    FromPerson  *string `json:"from_person,omitempty" required:"false" `
    /*
        收货人     */
    ToPerson  *string `json:"to_person,omitempty" required:"false" `
    /*
        药品配送ref_enti_d企业     */
    DisRefEntId  *string `json:"dis_ref_ent_id,omitempty" required:"false" `
    /*
        药品配送企业ent_Id出库单     */
    DisEntId  *string `json:"dis_ent_id,omitempty" required:"false" `
    /*
        应收货总数量     */
    QuReceivable  *string `json:"qu_receivable,omitempty" required:"false" `
    /*
        是否验证，0：未通过验证，1：已验证     */
    XtIsCheck  *string `json:"xt_is_check,omitempty" required:"false" `
    /*
        未验证通过原因【验证未通过时填写】     */
    XtCheckCode  *string `json:"xt_check_code,omitempty" required:"false" `
    /*
        未验证通过原因描述【验证未通过时填写】     */
    XtCheckCodeDesc  *string `json:"xt_check_code_desc,omitempty" required:"false" `
    /*
        药品列表Json："codeCount": 药品数量 "commDrugId": 国家药品唯一标识 "exprieDate": 生产日期 "physicInfo": 药品信息 "pkgSpec": 包状规格 "prepnCount": 制剂数量 "produceBatchNo":生产批次 "produceDate": 生产日期     */
    DrugListJson  *string `json:"drug_list_json,omitempty" required:"false" `
    /*
        单据委托企业refEntId     */
    AssRefEntId  *string `json:"ass_ref_ent_id,omitempty" required:"false" `
    /*
        单据委托企业entId     */
    AssEntId  *string `json:"ass_ent_id,omitempty" required:"false" `
}

func (s *AlibabaAlihealthDrugMscUploadinoutbillRequest) SetBillCode(v string) *AlibabaAlihealthDrugMscUploadinoutbillRequest {
    s.BillCode = &v
    return s
}
func (s *AlibabaAlihealthDrugMscUploadinoutbillRequest) SetBillTime(v util.LocalTime) *AlibabaAlihealthDrugMscUploadinoutbillRequest {
    s.BillTime = &v
    return s
}
func (s *AlibabaAlihealthDrugMscUploadinoutbillRequest) SetBillType(v int64) *AlibabaAlihealthDrugMscUploadinoutbillRequest {
    s.BillType = &v
    return s
}
func (s *AlibabaAlihealthDrugMscUploadinoutbillRequest) SetPhysicType(v int64) *AlibabaAlihealthDrugMscUploadinoutbillRequest {
    s.PhysicType = &v
    return s
}
func (s *AlibabaAlihealthDrugMscUploadinoutbillRequest) SetRefUserId(v string) *AlibabaAlihealthDrugMscUploadinoutbillRequest {
    s.RefUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugMscUploadinoutbillRequest) SetAgentRefUserId(v string) *AlibabaAlihealthDrugMscUploadinoutbillRequest {
    s.AgentRefUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugMscUploadinoutbillRequest) SetFromUserId(v string) *AlibabaAlihealthDrugMscUploadinoutbillRequest {
    s.FromUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugMscUploadinoutbillRequest) SetToUserId(v string) *AlibabaAlihealthDrugMscUploadinoutbillRequest {
    s.ToUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugMscUploadinoutbillRequest) SetDestUserId(v string) *AlibabaAlihealthDrugMscUploadinoutbillRequest {
    s.DestUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugMscUploadinoutbillRequest) SetOperIcCode(v string) *AlibabaAlihealthDrugMscUploadinoutbillRequest {
    s.OperIcCode = &v
    return s
}
func (s *AlibabaAlihealthDrugMscUploadinoutbillRequest) SetOperIcName(v string) *AlibabaAlihealthDrugMscUploadinoutbillRequest {
    s.OperIcName = &v
    return s
}
func (s *AlibabaAlihealthDrugMscUploadinoutbillRequest) SetWarehouseId(v string) *AlibabaAlihealthDrugMscUploadinoutbillRequest {
    s.WarehouseId = &v
    return s
}
func (s *AlibabaAlihealthDrugMscUploadinoutbillRequest) SetDrugId(v string) *AlibabaAlihealthDrugMscUploadinoutbillRequest {
    s.DrugId = &v
    return s
}
func (s *AlibabaAlihealthDrugMscUploadinoutbillRequest) SetTraceCodes(v []string) *AlibabaAlihealthDrugMscUploadinoutbillRequest {
    s.TraceCodes = &v
    return s
}
func (s *AlibabaAlihealthDrugMscUploadinoutbillRequest) SetClientType(v string) *AlibabaAlihealthDrugMscUploadinoutbillRequest {
    s.ClientType = &v
    return s
}
func (s *AlibabaAlihealthDrugMscUploadinoutbillRequest) SetReturnReasonCode(v string) *AlibabaAlihealthDrugMscUploadinoutbillRequest {
    s.ReturnReasonCode = &v
    return s
}
func (s *AlibabaAlihealthDrugMscUploadinoutbillRequest) SetReturnReasonDes(v string) *AlibabaAlihealthDrugMscUploadinoutbillRequest {
    s.ReturnReasonDes = &v
    return s
}
func (s *AlibabaAlihealthDrugMscUploadinoutbillRequest) SetCancelReasonCode(v string) *AlibabaAlihealthDrugMscUploadinoutbillRequest {
    s.CancelReasonCode = &v
    return s
}
func (s *AlibabaAlihealthDrugMscUploadinoutbillRequest) SetCancelReasonDes(v string) *AlibabaAlihealthDrugMscUploadinoutbillRequest {
    s.CancelReasonDes = &v
    return s
}
func (s *AlibabaAlihealthDrugMscUploadinoutbillRequest) SetExecuterName(v string) *AlibabaAlihealthDrugMscUploadinoutbillRequest {
    s.ExecuterName = &v
    return s
}
func (s *AlibabaAlihealthDrugMscUploadinoutbillRequest) SetExecuterCode(v string) *AlibabaAlihealthDrugMscUploadinoutbillRequest {
    s.ExecuterCode = &v
    return s
}
func (s *AlibabaAlihealthDrugMscUploadinoutbillRequest) SetSuperviserName(v string) *AlibabaAlihealthDrugMscUploadinoutbillRequest {
    s.SuperviserName = &v
    return s
}
func (s *AlibabaAlihealthDrugMscUploadinoutbillRequest) SetSuperviserCode(v string) *AlibabaAlihealthDrugMscUploadinoutbillRequest {
    s.SuperviserCode = &v
    return s
}
func (s *AlibabaAlihealthDrugMscUploadinoutbillRequest) SetFromAddress(v string) *AlibabaAlihealthDrugMscUploadinoutbillRequest {
    s.FromAddress = &v
    return s
}
func (s *AlibabaAlihealthDrugMscUploadinoutbillRequest) SetToAddress(v string) *AlibabaAlihealthDrugMscUploadinoutbillRequest {
    s.ToAddress = &v
    return s
}
func (s *AlibabaAlihealthDrugMscUploadinoutbillRequest) SetFromBillCode(v string) *AlibabaAlihealthDrugMscUploadinoutbillRequest {
    s.FromBillCode = &v
    return s
}
func (s *AlibabaAlihealthDrugMscUploadinoutbillRequest) SetOrderCode(v string) *AlibabaAlihealthDrugMscUploadinoutbillRequest {
    s.OrderCode = &v
    return s
}
func (s *AlibabaAlihealthDrugMscUploadinoutbillRequest) SetFromPerson(v string) *AlibabaAlihealthDrugMscUploadinoutbillRequest {
    s.FromPerson = &v
    return s
}
func (s *AlibabaAlihealthDrugMscUploadinoutbillRequest) SetToPerson(v string) *AlibabaAlihealthDrugMscUploadinoutbillRequest {
    s.ToPerson = &v
    return s
}
func (s *AlibabaAlihealthDrugMscUploadinoutbillRequest) SetDisRefEntId(v string) *AlibabaAlihealthDrugMscUploadinoutbillRequest {
    s.DisRefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugMscUploadinoutbillRequest) SetDisEntId(v string) *AlibabaAlihealthDrugMscUploadinoutbillRequest {
    s.DisEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugMscUploadinoutbillRequest) SetQuReceivable(v string) *AlibabaAlihealthDrugMscUploadinoutbillRequest {
    s.QuReceivable = &v
    return s
}
func (s *AlibabaAlihealthDrugMscUploadinoutbillRequest) SetXtIsCheck(v string) *AlibabaAlihealthDrugMscUploadinoutbillRequest {
    s.XtIsCheck = &v
    return s
}
func (s *AlibabaAlihealthDrugMscUploadinoutbillRequest) SetXtCheckCode(v string) *AlibabaAlihealthDrugMscUploadinoutbillRequest {
    s.XtCheckCode = &v
    return s
}
func (s *AlibabaAlihealthDrugMscUploadinoutbillRequest) SetXtCheckCodeDesc(v string) *AlibabaAlihealthDrugMscUploadinoutbillRequest {
    s.XtCheckCodeDesc = &v
    return s
}
func (s *AlibabaAlihealthDrugMscUploadinoutbillRequest) SetDrugListJson(v string) *AlibabaAlihealthDrugMscUploadinoutbillRequest {
    s.DrugListJson = &v
    return s
}
func (s *AlibabaAlihealthDrugMscUploadinoutbillRequest) SetAssRefEntId(v string) *AlibabaAlihealthDrugMscUploadinoutbillRequest {
    s.AssRefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugMscUploadinoutbillRequest) SetAssEntId(v string) *AlibabaAlihealthDrugMscUploadinoutbillRequest {
    s.AssEntId = &v
    return s
}

func (req *AlibabaAlihealthDrugMscUploadinoutbillRequest) ToMap() map[string]interface{} {
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
    if(req.PhysicType != nil) {
        paramMap["physic_type"] = *req.PhysicType
    }
    if(req.RefUserId != nil) {
        paramMap["ref_user_id"] = *req.RefUserId
    }
    if(req.AgentRefUserId != nil) {
        paramMap["agent_ref_user_id"] = *req.AgentRefUserId
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
    if(req.WarehouseId != nil) {
        paramMap["warehouse_id"] = *req.WarehouseId
    }
    if(req.DrugId != nil) {
        paramMap["drug_id"] = *req.DrugId
    }
    if(req.TraceCodes != nil) {
        paramMap["trace_codes"] = util.ConvertBasicList(*req.TraceCodes)
    }
    if(req.ClientType != nil) {
        paramMap["client_type"] = *req.ClientType
    }
    if(req.ReturnReasonCode != nil) {
        paramMap["return_reason_code"] = *req.ReturnReasonCode
    }
    if(req.ReturnReasonDes != nil) {
        paramMap["return_reason_des"] = *req.ReturnReasonDes
    }
    if(req.CancelReasonCode != nil) {
        paramMap["cancel_reason_code"] = *req.CancelReasonCode
    }
    if(req.CancelReasonDes != nil) {
        paramMap["cancel_reason_des"] = *req.CancelReasonDes
    }
    if(req.ExecuterName != nil) {
        paramMap["executer_name"] = *req.ExecuterName
    }
    if(req.ExecuterCode != nil) {
        paramMap["executer_code"] = *req.ExecuterCode
    }
    if(req.SuperviserName != nil) {
        paramMap["superviser_name"] = *req.SuperviserName
    }
    if(req.SuperviserCode != nil) {
        paramMap["superviser_code"] = *req.SuperviserCode
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
    return paramMap
}

func (req *AlibabaAlihealthDrugMscUploadinoutbillRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}