package domain


type AlibabaAlihealthDrugtraceTopYljgQueryListpartsPEntParDto struct {
    /*
        往来单位自定义编码     */
    PartnerId  *string `json:"partner_id,omitempty" `

    /*
        往来单位名称     */
    PartnerName  *string `json:"partner_name,omitempty" `

    /*
        企业id（查询企业自已的）     */
    EntId  *string `json:"ent_id,omitempty" `

    /*
        查询企业的唯一标识（查询企业自已的）     */
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
        类型     */
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

func (s *AlibabaAlihealthDrugtraceTopYljgQueryListpartsPEntParDto) SetPartnerId(v string) *AlibabaAlihealthDrugtraceTopYljgQueryListpartsPEntParDto {
    s.PartnerId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryListpartsPEntParDto) SetPartnerName(v string) *AlibabaAlihealthDrugtraceTopYljgQueryListpartsPEntParDto {
    s.PartnerName = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryListpartsPEntParDto) SetEntId(v string) *AlibabaAlihealthDrugtraceTopYljgQueryListpartsPEntParDto {
    s.EntId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryListpartsPEntParDto) SetRefEntId(v string) *AlibabaAlihealthDrugtraceTopYljgQueryListpartsPEntParDto {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryListpartsPEntParDto) SetEntProvCode(v string) *AlibabaAlihealthDrugtraceTopYljgQueryListpartsPEntParDto {
    s.EntProvCode = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryListpartsPEntParDto) SetProvName(v string) *AlibabaAlihealthDrugtraceTopYljgQueryListpartsPEntParDto {
    s.ProvName = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryListpartsPEntParDto) SetAreaName(v string) *AlibabaAlihealthDrugtraceTopYljgQueryListpartsPEntParDto {
    s.AreaName = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryListpartsPEntParDto) SetCityName(v string) *AlibabaAlihealthDrugtraceTopYljgQueryListpartsPEntParDto {
    s.CityName = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryListpartsPEntParDto) SetIsNetwork(v string) *AlibabaAlihealthDrugtraceTopYljgQueryListpartsPEntParDto {
    s.IsNetwork = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryListpartsPEntParDto) SetPartnerCapitalName(v string) *AlibabaAlihealthDrugtraceTopYljgQueryListpartsPEntParDto {
    s.PartnerCapitalName = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryListpartsPEntParDto) SetPartnerType(v string) *AlibabaAlihealthDrugtraceTopYljgQueryListpartsPEntParDto {
    s.PartnerType = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryListpartsPEntParDto) SetPartnerEntId(v string) *AlibabaAlihealthDrugtraceTopYljgQueryListpartsPEntParDto {
    s.PartnerEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryListpartsPEntParDto) SetLastModDate(v string) *AlibabaAlihealthDrugtraceTopYljgQueryListpartsPEntParDto {
    s.LastModDate = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryListpartsPEntParDto) SetCrtDate(v string) *AlibabaAlihealthDrugtraceTopYljgQueryListpartsPEntParDto {
    s.CrtDate = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryListpartsPEntParDto) SetCrtIcName(v string) *AlibabaAlihealthDrugtraceTopYljgQueryListpartsPEntParDto {
    s.CrtIcName = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryListpartsPEntParDto) SetStatus(v string) *AlibabaAlihealthDrugtraceTopYljgQueryListpartsPEntParDto {
    s.Status = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryListpartsPEntParDto) SetModIcName(v string) *AlibabaAlihealthDrugtraceTopYljgQueryListpartsPEntParDto {
    s.ModIcName = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryListpartsPEntParDto) SetPartnerLevel(v string) *AlibabaAlihealthDrugtraceTopYljgQueryListpartsPEntParDto {
    s.PartnerLevel = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryListpartsPEntParDto) SetModIcCode(v string) *AlibabaAlihealthDrugtraceTopYljgQueryListpartsPEntParDto {
    s.ModIcCode = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryListpartsPEntParDto) SetPEntParId(v string) *AlibabaAlihealthDrugtraceTopYljgQueryListpartsPEntParDto {
    s.PEntParId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryListpartsPEntParDto) SetCrtIcCode(v string) *AlibabaAlihealthDrugtraceTopYljgQueryListpartsPEntParDto {
    s.CrtIcCode = &v
    return s
}
