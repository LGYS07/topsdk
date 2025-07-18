package domain


type AlibabaAlihealthDrugYljgSaveentAddEntReqDto struct {
    /*
        企业详细注册地址     */
    DictRegionDetail  *string `json:"dict_region_detail,omitempty" `

    /*
        用类型code值，并且需要与userRoleType值相映射：比如userRoleType传的3，只能取枚举值中第三个参数是3的类型；如果没有映射的细分类型，则传0. 企业/机构细分类型枚举结构：(类型code，类型名称，对应企业/机构类型code)；枚举值：PRODUCE_CHINA(13, "境内生产企业", 1), PRODUCE_FOREIGN(14, "境外生产企业", 1), DISEASE_CENTER(31, "疾控中心", 3), INOCULATION_ENT(32, "接种单位", 3), OTHER_MEDICAL_ENT(33, "其他", 3), MEDICAL_34(34, "综合医院", 3), MEDICAL_35(35, "中医医院", 3), MEDICAL_36(36, "中西医结合医院", 3), MEDICAL_37(37, "民族医院", 3), MEDICAL_38(38, "专科医院，不含中医专科医院", 3), MEDICAL_39(39, "疗养院", 3), MEDICAL_310(310, "护理院(站)", 3), MEDICAL_311(311, "社区卫生服务中心", 3), MEDICAL_312(312, "社区卫生服务站", 3), MEDICAL_313(313, "街道卫生院", 3), MEDICAL_314(314, "乡镇卫生院", 3), MEDICAL_315(315, "门诊部", 3), MEDICAL_316(316, "诊所", 3), MEDICAL_317(317, "卫生所(室)", 3), MEDICAL_318(318, "医务室", 3), MEDICAL_319(319, "中小学卫生保健所", 3), MEDICAL_320(320, "村卫生室", 3), MEDICAL_321(321, "急救中心", 3), MEDICAL_322(322, "急救中心站", 3), MEDICAL_323(323, "急救站", 3), MEDICAL_324(324, "血站", 3), MEDICAL_325(325, "单采血浆站", 3), MEDICAL_326(326, "妇幼保健院", 3), MEDICAL_327(327, "妇幼保健所包括妇女、儿童保健所", 3), MEDICAL_328(328, "妇幼保健站包括妇幼保健中心", 3), MEDICAL_329(329, "生殖保健中心", 3), MEDICAL_330(330, "专科疾病防治院", 3), MEDICAL_331(331, "专科疾病防治所(站、中心)", 3), MEDICAL_333(333, "卫生防疫站", 3), MEDICAL_334(334, "卫生防病中心", 3), MEDICAL_335(335, "预防保健中心", 3), MEDICAL_336(336, "卫生监督所(局)", 3), MEDICAL_337(337, "卫生(综合)监督检验(监测、检测)所(站)", 3), MEDICAL_338(338, "环境卫生监督检验(监测、检测)所(站)", 3), MEDICAL_339(339, "放射卫生监督检验(监测、检测)所(站)", 3), MEDICAL_340(340, "劳动(职业、工业)卫生监督检验(监测、检测)所(站)", 3), MEDICAL_341(341, "食品卫生监督检验(监测、检测)所(站)", 3), MEDICAL_342(342, "学校卫生监督检验(监测、检测)所(站)", 3), MEDICAL_343(343, "其他卫生监督检验(监测、检测)所(站)", 3), MEDICAL_344(344, "医学科学(研究)院（所）", 3), MEDICAL_345(345, "预防医学研究院（所）", 3), MEDICAL_346(346, "中医(药)研究院(所)", 3), MEDICAL_347(347, "中西医结合研究所", 3), MEDICAL_348(348, "民族医(药)学研究所", 3), MEDICAL_349(349, "医学专科研究所", 3), MEDICAL_350(350, "药学研究所（包括药用植物研究所）", 3), MEDICAL_351(351, "医学普通高中等学校", 3), MEDICAL_352(352, "医学成人学校", 3), MEDICAL_353(353, "医学在职培训机构包括各类卫生技术人员、管理人员培训中心等", 3), MEDICAL_354(354, "健康教育所包括健康教育研究所", 3), MEDICAL_355(355, "健康教育站(中心)包括卫生宣教馆", 3), MEDICAL_356(356, "临床检验中心（所、站）", 3), MEDICAL_357(357, "卫生新闻出版社", 3), MEDICAL_358(358, "其他卫生事业机构", 3), MEDICAL_359(359, "红十字会", 3), MEDICAL_360(360, "医学会包括预防医学会、护理学会、医院管理学会等各类卫生专业学会", 3), MEDICAL_361(361, "协会", 3), MEDICAL_362(362, "其他卫生社会团体", 3), CHAIN_DRUG_STORE(41, "连锁药店总部", 4), DIRECT_CHAIN_DRUGS_STORE(42, "直营连锁药店", 4), JOIN_CHAIN_DRUG_STORE(43, "加盟连锁药店", 4), SINGLE_DRUG_STORE(44, "单体药店", 4), OTHER_DRUG_STORE(45, "其他", 4), MARKET_AUTH_HOLD_CHAINA(71, "境内持证商", 7), MARKET_AUTH_HOLD_FOREIGN(72, "境外持证商", 7) ;     */
    EntOrgType  *int64 `json:"ent_org_type,omitempty" `

    /*
        企业其他地址省市区信息     */
    StoreAddrs  *[]AlibabaAlihealthDrugYljgSaveentAddress `json:"store_addrs,omitempty" `

    /*
        新增企业名称     */
    EntName  *string `json:"ent_name,omitempty" `

    /*
        用类型code值：比如是医疗机构则传3.企业/机构类型枚举结构：(类型code，类型名称)； 枚举值：PROD_ENT(1,"生产企业"), WHOLESALE_ENT(2,"批发企业"), MEDICAL_ENT(3,"医疗机构"), RESALE_ENT(4,"零售企业"), LOGISTICS_ENT(5,"物流企业"), MAH_ENT(7,"上市许可持有人"), CENTRAL_RANDOMIZATION_SYSTEM_PROVIDER(8,"中央随机化系统提供商");     */
    UserRoleType  *int64 `json:"user_role_type,omitempty" `

    /*
        企业注册地址省市区信息     */
    RegAddr  *AlibabaAlihealthDrugYljgSaveentAddress `json:"reg_addr,omitempty" `

    /*
        用类型code值：比如营业执照资质则传8.资质证照类型枚举结构：(类型code，类型值)；枚举值：DRUG_BUSINESS_LICENSE(1, "药品经营许可证"),PRACTICE_LICENSE_OF_MEDICAL_INSTITUTION(3, "医疗机构执业许可证"),PRACTICE_LICENSE_OF_MEDICAL(24, "诊所备案凭证"),PHARMACEUTICAL_PRODUCTION_LICENSE(5, "药品生产许可证"),BUSINESS_LICENSE(8, "企业法人营业执照"),PUBLIC_INSTITUTIONS_LICENSE(11, "事业单位法人证书"),IMPORT_DRUG_REGISTRATION_LICENSE(23, "进口药品注册证")     */
    LicenseType  *int64 `json:"license_type,omitempty" `

    /*
        追溯业务负责人联系电话     */
    MobilePhone  *string `json:"mobile_phone,omitempty" `

    /*
        唯一认证代码：唯一认证代码是证件的编号，是国家对您的唯一认证编码。 如：证件是营业执照，唯一认证编码就是营业执照上的社会信用代码或组织机构代码。 证件是身份证，唯一认证编码就是身份证号。 证件是护照，唯一认证编码就是护照编号。 证件是医疗职业许可证，唯一认证编码就是许可证上的卫生机构组织代码或登记号。 以药品注册证为证件入驻的境外企业，唯一认证代码可填“无”。     */
    OrgCode  *string `json:"org_code,omitempty" `

    /*
        固定电话     */
    Tel  *string `json:"tel,omitempty" `

    /*
        追溯业务负责人名称     */
    ContactPsnNm  *string `json:"contact_psn_nm,omitempty" `

    /*
        追溯负责人联系邮箱     */
    Email  *string `json:"email,omitempty" `

}

func (s *AlibabaAlihealthDrugYljgSaveentAddEntReqDto) SetDictRegionDetail(v string) *AlibabaAlihealthDrugYljgSaveentAddEntReqDto {
    s.DictRegionDetail = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgSaveentAddEntReqDto) SetEntOrgType(v int64) *AlibabaAlihealthDrugYljgSaveentAddEntReqDto {
    s.EntOrgType = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgSaveentAddEntReqDto) SetStoreAddrs(v []AlibabaAlihealthDrugYljgSaveentAddress) *AlibabaAlihealthDrugYljgSaveentAddEntReqDto {
    s.StoreAddrs = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgSaveentAddEntReqDto) SetEntName(v string) *AlibabaAlihealthDrugYljgSaveentAddEntReqDto {
    s.EntName = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgSaveentAddEntReqDto) SetUserRoleType(v int64) *AlibabaAlihealthDrugYljgSaveentAddEntReqDto {
    s.UserRoleType = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgSaveentAddEntReqDto) SetRegAddr(v AlibabaAlihealthDrugYljgSaveentAddress) *AlibabaAlihealthDrugYljgSaveentAddEntReqDto {
    s.RegAddr = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgSaveentAddEntReqDto) SetLicenseType(v int64) *AlibabaAlihealthDrugYljgSaveentAddEntReqDto {
    s.LicenseType = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgSaveentAddEntReqDto) SetMobilePhone(v string) *AlibabaAlihealthDrugYljgSaveentAddEntReqDto {
    s.MobilePhone = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgSaveentAddEntReqDto) SetOrgCode(v string) *AlibabaAlihealthDrugYljgSaveentAddEntReqDto {
    s.OrgCode = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgSaveentAddEntReqDto) SetTel(v string) *AlibabaAlihealthDrugYljgSaveentAddEntReqDto {
    s.Tel = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgSaveentAddEntReqDto) SetContactPsnNm(v string) *AlibabaAlihealthDrugYljgSaveentAddEntReqDto {
    s.ContactPsnNm = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgSaveentAddEntReqDto) SetEmail(v string) *AlibabaAlihealthDrugYljgSaveentAddEntReqDto {
    s.Email = &v
    return s
}
