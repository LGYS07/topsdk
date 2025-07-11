package request

import (
	"github.com/LGYS07/topsdk/util"
)

type AlibabaAlihealthDrugtraceTopYljgUploadretailRequest struct {
	/*
	   单据编号（唯一）     */
	BillCode *string `json:"bill_code" required:"true" `
	/*
	   单据时间（一般为药品入出库时间）     */
	BillTime *util.LocalTime `json:"bill_time" required:"true" `
	/*
	   单据类型[323,零售出库][322,疫苗接种][116,消费者退货入库]     */
	BillType *int64 `json:"bill_type" required:"true" `
	/*
	   药品类型[2,特药，3,普药]【可以随便填写，单据上传后会以实际为准】 defalutValue��3    */
	PhysicType *int64 `json:"physic_type,omitempty" required:"false" `
	/*
	   上传单据的医疗机构在码上放心平台的ref_ent_id，可通过“通过企业名得到唯一标识”接口获取     */
	RefUserId *string `json:"ref_user_id" required:"true" `
	/*
	   发货企业(可为空)     */
	FromUserId *string `json:"from_user_id,omitempty" required:"false" `
	/*
	   单据提交者(appkey编号、可为空)     */
	OperIcCode *string `json:"oper_ic_code,omitempty" required:"false" `
	/*
	   单据提交者姓名（可为空）     */
	OperIcName *string `json:"oper_ic_name,omitempty" required:"false" `
	/*
	   追溯码【多个码时用逗号拼接的字符串。要求数量在3500个码以下，但一般不要传这么多，如果网络不好很容易传输一半报错】；注意：在同一张单据里，不能有重复的码；在同一张单据中不能同时上传有关联关系的大、小码；     */
	TraceCodes *[]string `json:"trace_codes" required:"true" `
	/*
	   购买人证件类型【1身份证2护照3 军官证4 医保卡5接种卡6学生证9其它】     */
	CustomerIdType *string `json:"customer_id_type,omitempty" required:"false" `
	/*
	   购买人证件编号     */
	CustomerId *string `json:"customer_id,omitempty" required:"false" `
	/*
	   用药人电话     */
	UserTel *string `json:"user_tel,omitempty" required:"false" `
	/*
	   互联标识 1是  0否     */
	NetworkBillFlag *string `json:"network_bill_flag,omitempty" required:"false" `
	/*
	   处方医师     */
	MedicDoctor *string `json:"medic_doctor,omitempty" required:"false" `
	/*
	   发药人     */
	MedicDispenser *string `json:"medic_dispenser,omitempty" required:"false" `
	/*
	   患者（姓名、院内患者ID均可）     */
	UserName *string `json:"user_name,omitempty" required:"false" `
	/*
	   代理领药人     */
	UserAgent *string `json:"user_agent,omitempty" required:"false" `
	/*
	   备注     */
	Remarks *string `json:"remarks,omitempty" required:"false" `
}

func (s *AlibabaAlihealthDrugtraceTopYljgUploadretailRequest) SetBillCode(v string) *AlibabaAlihealthDrugtraceTopYljgUploadretailRequest {
	s.BillCode = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgUploadretailRequest) SetBillTime(v util.LocalTime) *AlibabaAlihealthDrugtraceTopYljgUploadretailRequest {
	s.BillTime = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgUploadretailRequest) SetBillType(v int64) *AlibabaAlihealthDrugtraceTopYljgUploadretailRequest {
	s.BillType = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgUploadretailRequest) SetPhysicType(v int64) *AlibabaAlihealthDrugtraceTopYljgUploadretailRequest {
	s.PhysicType = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgUploadretailRequest) SetRefUserId(v string) *AlibabaAlihealthDrugtraceTopYljgUploadretailRequest {
	s.RefUserId = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgUploadretailRequest) SetFromUserId(v string) *AlibabaAlihealthDrugtraceTopYljgUploadretailRequest {
	s.FromUserId = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgUploadretailRequest) SetOperIcCode(v string) *AlibabaAlihealthDrugtraceTopYljgUploadretailRequest {
	s.OperIcCode = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgUploadretailRequest) SetOperIcName(v string) *AlibabaAlihealthDrugtraceTopYljgUploadretailRequest {
	s.OperIcName = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgUploadretailRequest) SetTraceCodes(v []string) *AlibabaAlihealthDrugtraceTopYljgUploadretailRequest {
	s.TraceCodes = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgUploadretailRequest) SetCustomerIdType(v string) *AlibabaAlihealthDrugtraceTopYljgUploadretailRequest {
	s.CustomerIdType = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgUploadretailRequest) SetCustomerId(v string) *AlibabaAlihealthDrugtraceTopYljgUploadretailRequest {
	s.CustomerId = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgUploadretailRequest) SetUserTel(v string) *AlibabaAlihealthDrugtraceTopYljgUploadretailRequest {
	s.UserTel = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgUploadretailRequest) SetNetworkBillFlag(v string) *AlibabaAlihealthDrugtraceTopYljgUploadretailRequest {
	s.NetworkBillFlag = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgUploadretailRequest) SetMedicDoctor(v string) *AlibabaAlihealthDrugtraceTopYljgUploadretailRequest {
	s.MedicDoctor = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgUploadretailRequest) SetMedicDispenser(v string) *AlibabaAlihealthDrugtraceTopYljgUploadretailRequest {
	s.MedicDispenser = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgUploadretailRequest) SetUserName(v string) *AlibabaAlihealthDrugtraceTopYljgUploadretailRequest {
	s.UserName = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgUploadretailRequest) SetUserAgent(v string) *AlibabaAlihealthDrugtraceTopYljgUploadretailRequest {
	s.UserAgent = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgUploadretailRequest) SetRemarks(v string) *AlibabaAlihealthDrugtraceTopYljgUploadretailRequest {
	s.Remarks = &v
	return s
}

func (req *AlibabaAlihealthDrugtraceTopYljgUploadretailRequest) ToMap() map[string]interface{} {
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
	if req.FromUserId != nil {
		paramMap["from_user_id"] = *req.FromUserId
	}
	if req.OperIcCode != nil {
		paramMap["oper_ic_code"] = *req.OperIcCode
	}
	if req.OperIcName != nil {
		paramMap["oper_ic_name"] = *req.OperIcName
	}
	if req.TraceCodes != nil {
		paramMap["trace_codes"] = util.ConvertBasicList(*req.TraceCodes)
	}
	if req.CustomerIdType != nil {
		paramMap["customer_id_type"] = *req.CustomerIdType
	}
	if req.CustomerId != nil {
		paramMap["customer_id"] = *req.CustomerId
	}
	if req.UserTel != nil {
		paramMap["user_tel"] = *req.UserTel
	}
	if req.NetworkBillFlag != nil {
		paramMap["network_bill_flag"] = *req.NetworkBillFlag
	}
	if req.MedicDoctor != nil {
		paramMap["medic_doctor"] = *req.MedicDoctor
	}
	if req.MedicDispenser != nil {
		paramMap["medic_dispenser"] = *req.MedicDispenser
	}
	if req.UserName != nil {
		paramMap["user_name"] = *req.UserName
	}
	if req.UserAgent != nil {
		paramMap["user_agent"] = *req.UserAgent
	}
	if req.Remarks != nil {
		paramMap["remarks"] = *req.Remarks
	}
	return paramMap
}

func (req *AlibabaAlihealthDrugtraceTopYljgUploadretailRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
