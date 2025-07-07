package domain


type AlibabaAlihealthDrugMscListpartsPEntParDTO struct {
    /*
        往来单位自定义编码     */
    PartnerId  *string `json:"partner_id,omitempty" `

    /*
        往来单位名称     */
    PartnerName  *string `json:"partner_name,omitempty" `

    /*
        企业id     */
    EntId  *string `json:"ent_id,omitempty" `

    /*
        查询企业的唯一标识     */
    RefEntId  *string `json:"ref_ent_id,omitempty" `

    /*
        往来单位企业所在省编码     */
    EntProvCode  *string `json:"ent_prov_code,omitempty" `

    /*
        所在省     */
    ProvName  *string `json:"prov_name,omitempty" `

    /*
        所在市     */
    AreaName  *string `json:"area_name,omitempty" `

    /*
        所在县     */
    CityName  *string `json:"city_name,omitempty" `

    /*
        是不是入网企业[1代表入网企业，其它为非入网]     */
    IsNetwork  *string `json:"is_network,omitempty" `

    /*
        拼音缩写     */
    PartnerCapitalName  *string `json:"partner_capital_name,omitempty" `

    /*
        企业类型（1: 生产企业  2: 批发企业  3: 医疗机构  4: 零售企业  5: 物流企业  7: 上市许可持有人  8: 中央随机化系统提供商  110: 政府系统)     */
    PartnerType  *string `json:"partner_type,omitempty" `

    /*
        往来单位企业id【单据上传时的收发货企业填的就这个字段】     */
    PartnerEntId  *string `json:"partner_ent_id,omitempty" `

    /*
        最近修改日期     */
    LastModDate  *string `json:"last_mod_date,omitempty" `

    /*
        创建日期     */
    CrtDate  *string `json:"crt_date,omitempty" `

    /*
        创建IC名称     */
    CrtIcName  *string `json:"crt_ic_name,omitempty" `

    /*
        状态     */
    Status  *string `json:"status,omitempty" `

    /*
        修改IC名称     */
    ModIcName  *string `json:"mod_ic_name,omitempty" `

    /*
        级别     */
    PartnerLevel  *string `json:"partner_level,omitempty" `

    /*
        修改IC码     */
    ModIcCode  *string `json:"mod_ic_code,omitempty" `

    /*
        合作ID     */
    PEntParId  *string `json:"p_ent_par_id,omitempty" `

    /*
        创建IC码     */
    CrtIcCode  *string `json:"crt_ic_code,omitempty" `

}

func (s *AlibabaAlihealthDrugMscListpartsPEntParDTO) SetPartnerId(v string) *AlibabaAlihealthDrugMscListpartsPEntParDTO {
    s.PartnerId = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListpartsPEntParDTO) SetPartnerName(v string) *AlibabaAlihealthDrugMscListpartsPEntParDTO {
    s.PartnerName = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListpartsPEntParDTO) SetEntId(v string) *AlibabaAlihealthDrugMscListpartsPEntParDTO {
    s.EntId = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListpartsPEntParDTO) SetRefEntId(v string) *AlibabaAlihealthDrugMscListpartsPEntParDTO {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListpartsPEntParDTO) SetEntProvCode(v string) *AlibabaAlihealthDrugMscListpartsPEntParDTO {
    s.EntProvCode = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListpartsPEntParDTO) SetProvName(v string) *AlibabaAlihealthDrugMscListpartsPEntParDTO {
    s.ProvName = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListpartsPEntParDTO) SetAreaName(v string) *AlibabaAlihealthDrugMscListpartsPEntParDTO {
    s.AreaName = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListpartsPEntParDTO) SetCityName(v string) *AlibabaAlihealthDrugMscListpartsPEntParDTO {
    s.CityName = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListpartsPEntParDTO) SetIsNetwork(v string) *AlibabaAlihealthDrugMscListpartsPEntParDTO {
    s.IsNetwork = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListpartsPEntParDTO) SetPartnerCapitalName(v string) *AlibabaAlihealthDrugMscListpartsPEntParDTO {
    s.PartnerCapitalName = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListpartsPEntParDTO) SetPartnerType(v string) *AlibabaAlihealthDrugMscListpartsPEntParDTO {
    s.PartnerType = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListpartsPEntParDTO) SetPartnerEntId(v string) *AlibabaAlihealthDrugMscListpartsPEntParDTO {
    s.PartnerEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListpartsPEntParDTO) SetLastModDate(v string) *AlibabaAlihealthDrugMscListpartsPEntParDTO {
    s.LastModDate = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListpartsPEntParDTO) SetCrtDate(v string) *AlibabaAlihealthDrugMscListpartsPEntParDTO {
    s.CrtDate = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListpartsPEntParDTO) SetCrtIcName(v string) *AlibabaAlihealthDrugMscListpartsPEntParDTO {
    s.CrtIcName = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListpartsPEntParDTO) SetStatus(v string) *AlibabaAlihealthDrugMscListpartsPEntParDTO {
    s.Status = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListpartsPEntParDTO) SetModIcName(v string) *AlibabaAlihealthDrugMscListpartsPEntParDTO {
    s.ModIcName = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListpartsPEntParDTO) SetPartnerLevel(v string) *AlibabaAlihealthDrugMscListpartsPEntParDTO {
    s.PartnerLevel = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListpartsPEntParDTO) SetModIcCode(v string) *AlibabaAlihealthDrugMscListpartsPEntParDTO {
    s.ModIcCode = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListpartsPEntParDTO) SetPEntParId(v string) *AlibabaAlihealthDrugMscListpartsPEntParDTO {
    s.PEntParId = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListpartsPEntParDTO) SetCrtIcCode(v string) *AlibabaAlihealthDrugMscListpartsPEntParDTO {
    s.CrtIcCode = &v
    return s
}
