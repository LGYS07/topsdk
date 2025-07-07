package domain


type AlibabaAlihealthDrugLsydSaveentAddress struct {
    /*
        境内填写区县名称/境外则填写境外国家中文名称     */
    AreaName  *string `json:"area_name,omitempty" `

    /*
        城市名称/境外不用填，境内必填     */
    CityName  *string `json:"city_name,omitempty" `

    /*
        省份名称/境外不用填，境内必填     */
    ProvName  *string `json:"prov_name,omitempty" `

}

func (s *AlibabaAlihealthDrugLsydSaveentAddress) SetAreaName(v string) *AlibabaAlihealthDrugLsydSaveentAddress {
    s.AreaName = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydSaveentAddress) SetCityName(v string) *AlibabaAlihealthDrugLsydSaveentAddress {
    s.CityName = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydSaveentAddress) SetProvName(v string) *AlibabaAlihealthDrugLsydSaveentAddress {
    s.ProvName = &v
    return s
}
