package domain


type AlibabaAlihealthDrugtraceTopLsydQueryListpartsPEntParDto struct {
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

func (s *AlibabaAlihealthDrugtraceTopLsydQueryListpartsPEntParDto) SetPartnerId(v string) *AlibabaAlihealthDrugtraceTopLsydQueryListpartsPEntParDto {
    s.PartnerId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryListpartsPEntParDto) SetPartnerName(v string) *AlibabaAlihealthDrugtraceTopLsydQueryListpartsPEntParDto {
    s.PartnerName = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryListpartsPEntParDto) SetEntId(v string) *AlibabaAlihealthDrugtraceTopLsydQueryListpartsPEntParDto {
    s.EntId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryListpartsPEntParDto) SetRefEntId(v string) *AlibabaAlihealthDrugtraceTopLsydQueryListpartsPEntParDto {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryListpartsPEntParDto) SetEntProvCode(v string) *AlibabaAlihealthDrugtraceTopLsydQueryListpartsPEntParDto {
    s.EntProvCode = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryListpartsPEntParDto) SetProvName(v string) *AlibabaAlihealthDrugtraceTopLsydQueryListpartsPEntParDto {
    s.ProvName = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryListpartsPEntParDto) SetAreaName(v string) *AlibabaAlihealthDrugtraceTopLsydQueryListpartsPEntParDto {
    s.AreaName = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryListpartsPEntParDto) SetCityName(v string) *AlibabaAlihealthDrugtraceTopLsydQueryListpartsPEntParDto {
    s.CityName = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryListpartsPEntParDto) SetIsNetwork(v string) *AlibabaAlihealthDrugtraceTopLsydQueryListpartsPEntParDto {
    s.IsNetwork = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryListpartsPEntParDto) SetPartnerCapitalName(v string) *AlibabaAlihealthDrugtraceTopLsydQueryListpartsPEntParDto {
    s.PartnerCapitalName = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryListpartsPEntParDto) SetPartnerType(v string) *AlibabaAlihealthDrugtraceTopLsydQueryListpartsPEntParDto {
    s.PartnerType = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryListpartsPEntParDto) SetPartnerEntId(v string) *AlibabaAlihealthDrugtraceTopLsydQueryListpartsPEntParDto {
    s.PartnerEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryListpartsPEntParDto) SetLastModDate(v string) *AlibabaAlihealthDrugtraceTopLsydQueryListpartsPEntParDto {
    s.LastModDate = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryListpartsPEntParDto) SetCrtDate(v string) *AlibabaAlihealthDrugtraceTopLsydQueryListpartsPEntParDto {
    s.CrtDate = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryListpartsPEntParDto) SetCrtIcName(v string) *AlibabaAlihealthDrugtraceTopLsydQueryListpartsPEntParDto {
    s.CrtIcName = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryListpartsPEntParDto) SetStatus(v string) *AlibabaAlihealthDrugtraceTopLsydQueryListpartsPEntParDto {
    s.Status = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryListpartsPEntParDto) SetModIcName(v string) *AlibabaAlihealthDrugtraceTopLsydQueryListpartsPEntParDto {
    s.ModIcName = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryListpartsPEntParDto) SetPartnerLevel(v string) *AlibabaAlihealthDrugtraceTopLsydQueryListpartsPEntParDto {
    s.PartnerLevel = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryListpartsPEntParDto) SetModIcCode(v string) *AlibabaAlihealthDrugtraceTopLsydQueryListpartsPEntParDto {
    s.ModIcCode = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryListpartsPEntParDto) SetPEntParId(v string) *AlibabaAlihealthDrugtraceTopLsydQueryListpartsPEntParDto {
    s.PEntParId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryListpartsPEntParDto) SetCrtIcCode(v string) *AlibabaAlihealthDrugtraceTopLsydQueryListpartsPEntParDto {
    s.CrtIcCode = &v
    return s
}
