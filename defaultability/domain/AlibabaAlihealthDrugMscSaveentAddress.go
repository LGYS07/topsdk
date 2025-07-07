package domain


type AlibabaAlihealthDrugMscSaveentAddress struct {
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

func (s *AlibabaAlihealthDrugMscSaveentAddress) SetAreaName(v string) *AlibabaAlihealthDrugMscSaveentAddress {
    s.AreaName = &v
    return s
}
func (s *AlibabaAlihealthDrugMscSaveentAddress) SetCityName(v string) *AlibabaAlihealthDrugMscSaveentAddress {
    s.CityName = &v
    return s
}
func (s *AlibabaAlihealthDrugMscSaveentAddress) SetProvName(v string) *AlibabaAlihealthDrugMscSaveentAddress {
    s.ProvName = &v
    return s
}
