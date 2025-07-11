package domain


type AlibabaAlihealthDrugYljgSaveentAddress struct {
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

func (s *AlibabaAlihealthDrugYljgSaveentAddress) SetAreaName(v string) *AlibabaAlihealthDrugYljgSaveentAddress {
    s.AreaName = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgSaveentAddress) SetCityName(v string) *AlibabaAlihealthDrugYljgSaveentAddress {
    s.CityName = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgSaveentAddress) SetProvName(v string) *AlibabaAlihealthDrugYljgSaveentAddress {
    s.ProvName = &v
    return s
}
