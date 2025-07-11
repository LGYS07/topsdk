package request

import (
	"github.com/LGYS07/topsdk/util"
)

type AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest struct {
	/*
	   单据编号（唯一）     */
	BillCode *string `json:"bill_code" required:"true" `
	/*
	   单据时间（一般为药品入出库时间）     */
	BillTime *util.LocalTime `json:"bill_time" required:"true" `
	/*
	   单据类型：102代表采购入库、202代表退货出库、205代表销毁出库     */
	BillType *int64 `json:"bill_type" required:"true" `
	/*
	   药品类型[2,特药，3,普药]【可以随便填写，单据上传后会以实际为准】     */
	PhysicType *int64 `json:"physic_type" required:"true" `
	/*
	   上传单据的医疗机构在码上放心平台的ref_ent_id，可通过“通过企业名得到唯一标识”接口获取     */
	RefUserId *string `json:"ref_user_id" required:"true" `
	/*
	   代理企业REF标识     */
	AgentRefUserId *string `json:"agent_ref_user_id,omitempty" required:"false" `
	/*
	   发货企业ent_id，可通过“通过企业名得到唯一标识”接口获取；（102采购入库填药品供应商id、202退货出库填医院id、205销毁出库填医院id）     */
	FromUserId *string `json:"from_user_id" required:"true" `
	/*
	   收货企业ent_id，可通过“通过企业名得到唯一标识”接口获取；（102采购入库填医院id、202退货出库填药品供应商id、205销毁出库填医院id）     */
	ToUserId *string `json:"to_user_id" required:"true" `
	/*
	   直调企业标识     */
	DestUserId *string `json:"dest_user_id,omitempty" required:"false" `
	/*
	   单据提交者(appkey编号、可为空)       */
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
	   （协同平台数据合规）发货地址（可为空）     */
	FromAddress *string `json:"from_address,omitempty" required:"false" `
	/*
	   （协同平台数据合规）收货地址（可为空）     */
	ToAddress *string `json:"to_address,omitempty" required:"false" `
	/*
	   （协同平台数据合规）发货单编号（可为空）     */
	FromBillCode *string `json:"from_bill_code,omitempty" required:"false" `
	/*
	   （协同平台数据合规）订货单编号（可为空）     */
	OrderCode *string `json:"order_code,omitempty" required:"false" `
	/*
	   （协同平台数据合规）发货人（可为空）     */
	FromPerson *string `json:"from_person,omitempty" required:"false" `
	/*
	   （协同平台数据合规）收货人（可为空）     */
	ToPerson *string `json:"to_person,omitempty" required:"false" `
	/*
	   （协同平台数据合规）药品配送企业【添写ref_ent_id】     */
	DisRefEntId *string `json:"dis_ref_ent_id,omitempty" required:"false" `
	/*
	   （协同平台数据合规）药品配送企业entId【添写ent_id】     */
	DisEntId *string `json:"dis_ent_id,omitempty" required:"false" `
	/*
	   （协同平台数据合规）应收货总数量（可为空）     */
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
	   （协同平台数据合规）药品列表Json[可不填写]     */
	DrugListJson *string `json:"drug_list_json,omitempty" required:"false" `
	/*
	   （协同平台数据合规）单据委托企业refEntId【疫苗药品出库单填写】     */
	AssRefEntId *string `json:"ass_ref_ent_id,omitempty" required:"false" `
	/*
	   （协同平台数据合规）单据委托企业entId【疫苗药品出库单填写】     */
	AssEntId *string `json:"ass_ent_id,omitempty" required:"false" `
}

func (s *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) SetBillCode(v string) *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest {
	s.BillCode = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) SetBillTime(v util.LocalTime) *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest {
	s.BillTime = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) SetBillType(v int64) *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest {
	s.BillType = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) SetPhysicType(v int64) *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest {
	s.PhysicType = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) SetRefUserId(v string) *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest {
	s.RefUserId = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) SetAgentRefUserId(v string) *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest {
	s.AgentRefUserId = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) SetFromUserId(v string) *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest {
	s.FromUserId = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) SetToUserId(v string) *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest {
	s.ToUserId = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) SetDestUserId(v string) *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest {
	s.DestUserId = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) SetOperIcCode(v string) *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest {
	s.OperIcCode = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) SetOperIcName(v string) *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest {
	s.OperIcName = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) SetWarehouseId(v string) *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest {
	s.WarehouseId = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) SetDrugId(v string) *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest {
	s.DrugId = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) SetTraceCodes(v []string) *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest {
	s.TraceCodes = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) SetClientType(v string) *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest {
	s.ClientType = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) SetReturnReasonCode(v string) *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest {
	s.ReturnReasonCode = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) SetReturnReasonDes(v string) *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest {
	s.ReturnReasonDes = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) SetCancelReasonCode(v string) *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest {
	s.CancelReasonCode = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) SetCancelReasonDes(v string) *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest {
	s.CancelReasonDes = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) SetExecuterName(v string) *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest {
	s.ExecuterName = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) SetExecuterCode(v string) *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest {
	s.ExecuterCode = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) SetSuperviserName(v string) *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest {
	s.SuperviserName = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) SetSuperviserCode(v string) *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest {
	s.SuperviserCode = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) SetFromAddress(v string) *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest {
	s.FromAddress = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) SetToAddress(v string) *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest {
	s.ToAddress = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) SetFromBillCode(v string) *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest {
	s.FromBillCode = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) SetOrderCode(v string) *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest {
	s.OrderCode = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) SetFromPerson(v string) *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest {
	s.FromPerson = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) SetToPerson(v string) *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest {
	s.ToPerson = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) SetDisRefEntId(v string) *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest {
	s.DisRefEntId = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) SetDisEntId(v string) *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest {
	s.DisEntId = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) SetQuReceivable(v int64) *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest {
	s.QuReceivable = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) SetXtIsCheck(v string) *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest {
	s.XtIsCheck = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) SetXtCheckCode(v string) *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest {
	s.XtCheckCode = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) SetXtCheckCodeDesc(v string) *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest {
	s.XtCheckCodeDesc = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) SetDrugListJson(v string) *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest {
	s.DrugListJson = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) SetAssRefEntId(v string) *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest {
	s.AssRefEntId = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) SetAssEntId(v string) *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest {
	s.AssEntId = &v
	return s
}

func (req *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) ToMap() map[string]interface{} {
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

func (req *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
