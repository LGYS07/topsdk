package request

import (
	"github.com/LGYS07/topsdk/util"
)

type AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest struct {
	/*
	   单据编码（每个单据号只能上传一次）     */
	BillCode *string `json:"bill_code" required:"true" `
	/*
	   单据时间（一般为药品入出库时间）     */
	BillTime *util.LocalTime `json:"bill_time" required:"true" `
	/*
	   单据类型【102代表采购入库】     */
	BillType *int64 `json:"bill_type" required:"true" `
	/*
	   药品类型【3普药2特药】可不填     */
	PhysicType *int64 `json:"physic_type" required:"true" `
	/*
	   上传单据企业的单位编码【注意：该入参是ref_ent_id，不是ent_id】可通过“通过企业名得到唯一标识”接口获取     */
	RefUserId *string `json:"ref_user_id" required:"true" `
	/*
	   代理企业【注意：该入参是ref_ent_id，不是ent_id】，该字段兼容之前接口逻辑，后期将不允许使用，不要填值。     */
	AgentRefUserId *string `json:"agent_ref_user_id,omitempty" required:"false" `
	/*
	   发货企业entId，可通过“通过企业名得到唯一标识”接口获取     */
	FromUserId *string `json:"from_user_id" required:"true" `
	/*
	   收货企业entId，可通过“通过企业名得到唯一标识”接口获取     */
	ToUserId *string `json:"to_user_id" required:"true" `
	/*
	   直调企业标识     */
	DestUserId *string `json:"dest_user_id,omitempty" required:"false" `
	/*
	   单据提交者（appkey编号）可为空     */
	OperIcCode *string `json:"oper_ic_code,omitempty" required:"false" `
	/*
	   单据提交者姓名（可为空）     */
	OperIcName *string `json:"oper_ic_name,omitempty" required:"false" `
	/*
	   仓号     */
	WarehouseId *string `json:"warehouse_id,omitempty" required:"false" `
	/*
	   药品ID[企业自已系统的药品ID]     */
	DrugId *string `json:"drug_id,omitempty" required:"false" `
	/*
	   追溯码【多个码时用逗号拼接的字符串。要求数量在3500个码以下，但一般不要传这么多，如果网络不好很容易传输一半报错】注意：在同一张单据里，不能有重复的码；在同一张单据中不能同时上传有关联关系的大、小码     */
	TraceCodes *[]string `json:"trace_codes" required:"true" `
	/*
	   客户端类型[必须填2]     */
	ClientType *string `json:"client_type" required:"true" `
	/*
	   已废弃，无需填写     */
	ReturnReasonCode *string `json:"return_reason_code,omitempty" required:"false" `
	/*
	   已废弃，无需填写     */
	ReturnReasonDes *string `json:"return_reason_des,omitempty" required:"false" `
	/*
	   已废弃，无需填写     */
	CancelReasonCode *string `json:"cancel_reason_code,omitempty" required:"false" `
	/*
	   已废弃，无需填写     */
	CancelReasonDes *string `json:"cancel_reason_des,omitempty" required:"false" `
	/*
	   已废弃，无需填写     */
	ExecuterName *string `json:"executer_name,omitempty" required:"false" `
	/*
	   已废弃，无需填写     */
	ExecuterCode *string `json:"executer_code,omitempty" required:"false" `
	/*
	   已废弃，无需填写     */
	SuperviserName *string `json:"superviser_name,omitempty" required:"false" `
	/*
	   已废弃，无需填写     */
	SuperviserCode *string `json:"superviser_code,omitempty" required:"false" `
	/*
	   （协同平台数据合规）发货地址     */
	FromAddress *string `json:"from_address,omitempty" required:"false" `
	/*
	   （协同平台数据合规）收货地址     */
	ToAddress *string `json:"to_address,omitempty" required:"false" `
	/*
	   （协同平台数据合规）发货单编号     */
	FromBillCode *string `json:"from_bill_code,omitempty" required:"false" `
	/*
	   （协同平台数据合规）订货单编号     */
	OrderCode *string `json:"order_code,omitempty" required:"false" `
	/*
	   （协同平台数据合规）发货人     */
	FromPerson *string `json:"from_person,omitempty" required:"false" `
	/*
	   （协同平台数据合规）收货人     */
	ToPerson *string `json:"to_person,omitempty" required:"false" `
	/*
	   （协同平台数据合规）药品配送企业【出库单，收货方为医疗机构时填写】     */
	DisRefEntId *string `json:"dis_ref_ent_id,omitempty" required:"false" `
	/*
	   （协同平台数据合规）药品配送企业entId【出库单，收货方为医疗机构时填写】     */
	DisEntId *string `json:"dis_ent_id,omitempty" required:"false" `
	/*
	   （协同平台数据合规）应收货总数量【可为空】     */
	QuReceivable *int64 `json:"qu_receivable,omitempty" required:"false" `
	/*
	   （协同平台数据合规）是否验证，0：未通过验证，1：已验证     */
	XtIsCheck *string `json:"xt_is_check,omitempty" required:"false" `
	/*
	   （协同平台数据合规）未验证通过原因【验证未通过时填写】     */
	XtCheckCode *string `json:"xt_check_code,omitempty" required:"false" `
	/*
	   （协同平台数据合规）未验证通过原因描述【验证未通过时填写】     */
	XtCheckCodeDesc *string `json:"xt_check_code_desc,omitempty" required:"false" `
	/*
	   （协同平台数据合规）药品列表Json："codeCount": 药品数量 "commDrugId": 国家药品唯一标识 "exprieDate": 生产日期 "physicInfo": 药品信息 "pkgSpec": 包状规格 "prepnCount": 制剂数量 "produceBatchNo":生产批次 "produceDate": 生产日期     */
	DrugListJson *string `json:"drug_list_json,omitempty" required:"false" `
	/*
	   （协同平台数据合规）单据委托企业refEntId【疫苗药品出库单填写】     */
	AssRefEntId *string `json:"ass_ref_ent_id,omitempty" required:"false" `
	/*
	   （协同平台数据合规）单据委托企业entId【疫苗药品出库单填写】     */
	AssEntId *string `json:"ass_ent_id,omitempty" required:"false" `
}

func (s *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) SetBillCode(v string) *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest {
	s.BillCode = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) SetBillTime(v util.LocalTime) *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest {
	s.BillTime = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) SetBillType(v int64) *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest {
	s.BillType = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) SetPhysicType(v int64) *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest {
	s.PhysicType = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) SetRefUserId(v string) *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest {
	s.RefUserId = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) SetAgentRefUserId(v string) *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest {
	s.AgentRefUserId = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) SetFromUserId(v string) *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest {
	s.FromUserId = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) SetToUserId(v string) *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest {
	s.ToUserId = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) SetDestUserId(v string) *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest {
	s.DestUserId = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) SetOperIcCode(v string) *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest {
	s.OperIcCode = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) SetOperIcName(v string) *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest {
	s.OperIcName = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) SetWarehouseId(v string) *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest {
	s.WarehouseId = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) SetDrugId(v string) *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest {
	s.DrugId = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) SetTraceCodes(v []string) *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest {
	s.TraceCodes = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) SetClientType(v string) *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest {
	s.ClientType = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) SetReturnReasonCode(v string) *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest {
	s.ReturnReasonCode = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) SetReturnReasonDes(v string) *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest {
	s.ReturnReasonDes = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) SetCancelReasonCode(v string) *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest {
	s.CancelReasonCode = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) SetCancelReasonDes(v string) *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest {
	s.CancelReasonDes = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) SetExecuterName(v string) *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest {
	s.ExecuterName = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) SetExecuterCode(v string) *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest {
	s.ExecuterCode = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) SetSuperviserName(v string) *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest {
	s.SuperviserName = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) SetSuperviserCode(v string) *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest {
	s.SuperviserCode = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) SetFromAddress(v string) *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest {
	s.FromAddress = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) SetToAddress(v string) *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest {
	s.ToAddress = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) SetFromBillCode(v string) *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest {
	s.FromBillCode = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) SetOrderCode(v string) *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest {
	s.OrderCode = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) SetFromPerson(v string) *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest {
	s.FromPerson = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) SetToPerson(v string) *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest {
	s.ToPerson = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) SetDisRefEntId(v string) *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest {
	s.DisRefEntId = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) SetDisEntId(v string) *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest {
	s.DisEntId = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) SetQuReceivable(v int64) *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest {
	s.QuReceivable = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) SetXtIsCheck(v string) *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest {
	s.XtIsCheck = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) SetXtCheckCode(v string) *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest {
	s.XtCheckCode = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) SetXtCheckCodeDesc(v string) *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest {
	s.XtCheckCodeDesc = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) SetDrugListJson(v string) *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest {
	s.DrugListJson = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) SetAssRefEntId(v string) *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest {
	s.AssRefEntId = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) SetAssEntId(v string) *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest {
	s.AssEntId = &v
	return s
}

func (req *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.BillCode != nil {
		paramMap["bill_code"] = *req.BillCode
	}
	if req.BillTime != nil {
		paramMap["bill_time"] = *req.BillTime
	}
	if req.BillType != nil {
		paramMap["bill_type"] = *req.BillType
	}
	if req.PhysicType != nil {
		paramMap["physic_type"] = *req.PhysicType
	}
	if req.RefUserId != nil {
		paramMap["ref_user_id"] = *req.RefUserId
	}
	if req.AgentRefUserId != nil {
		paramMap["agent_ref_user_id"] = *req.AgentRefUserId
	}
	if req.FromUserId != nil {
		paramMap["from_user_id"] = *req.FromUserId
	}
	if req.ToUserId != nil {
		paramMap["to_user_id"] = *req.ToUserId
	}
	if req.DestUserId != nil {
		paramMap["dest_user_id"] = *req.DestUserId
	}
	if req.OperIcCode != nil {
		paramMap["oper_ic_code"] = *req.OperIcCode
	}
	if req.OperIcName != nil {
		paramMap["oper_ic_name"] = *req.OperIcName
	}
	if req.WarehouseId != nil {
		paramMap["warehouse_id"] = *req.WarehouseId
	}
	if req.DrugId != nil {
		paramMap["drug_id"] = *req.DrugId
	}
	if req.TraceCodes != nil {
		paramMap["trace_codes"] = util.ConvertBasicList(*req.TraceCodes)
	}
	if req.ClientType != nil {
		paramMap["client_type"] = *req.ClientType
	}
	if req.ReturnReasonCode != nil {
		paramMap["return_reason_code"] = *req.ReturnReasonCode
	}
	if req.ReturnReasonDes != nil {
		paramMap["return_reason_des"] = *req.ReturnReasonDes
	}
	if req.CancelReasonCode != nil {
		paramMap["cancel_reason_code"] = *req.CancelReasonCode
	}
	if req.CancelReasonDes != nil {
		paramMap["cancel_reason_des"] = *req.CancelReasonDes
	}
	if req.ExecuterName != nil {
		paramMap["executer_name"] = *req.ExecuterName
	}
	if req.ExecuterCode != nil {
		paramMap["executer_code"] = *req.ExecuterCode
	}
	if req.SuperviserName != nil {
		paramMap["superviser_name"] = *req.SuperviserName
	}
	if req.SuperviserCode != nil {
		paramMap["superviser_code"] = *req.SuperviserCode
	}
	if req.FromAddress != nil {
		paramMap["from_address"] = *req.FromAddress
	}
	if req.ToAddress != nil {
		paramMap["to_address"] = *req.ToAddress
	}
	if req.FromBillCode != nil {
		paramMap["from_bill_code"] = *req.FromBillCode
	}
	if req.OrderCode != nil {
		paramMap["order_code"] = *req.OrderCode
	}
	if req.FromPerson != nil {
		paramMap["from_person"] = *req.FromPerson
	}
	if req.ToPerson != nil {
		paramMap["to_person"] = *req.ToPerson
	}
	if req.DisRefEntId != nil {
		paramMap["dis_ref_ent_id"] = *req.DisRefEntId
	}
	if req.DisEntId != nil {
		paramMap["dis_ent_id"] = *req.DisEntId
	}
	if req.QuReceivable != nil {
		paramMap["qu_receivable"] = *req.QuReceivable
	}
	if req.XtIsCheck != nil {
		paramMap["xt_is_check"] = *req.XtIsCheck
	}
	if req.XtCheckCode != nil {
		paramMap["xt_check_code"] = *req.XtCheckCode
	}
	if req.XtCheckCodeDesc != nil {
		paramMap["xt_check_code_desc"] = *req.XtCheckCodeDesc
	}
	if req.DrugListJson != nil {
		paramMap["drug_list_json"] = *req.DrugListJson
	}
	if req.AssRefEntId != nil {
		paramMap["ass_ref_ent_id"] = *req.AssRefEntId
	}
	if req.AssEntId != nil {
		paramMap["ass_ent_id"] = *req.AssEntId
	}
	return paramMap
}

func (req *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
