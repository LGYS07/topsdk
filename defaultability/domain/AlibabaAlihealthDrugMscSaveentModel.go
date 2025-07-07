package domain


type AlibabaAlihealthDrugMscSaveentModel struct {
    /*
        新增失败的时候错误原因     */
    CheckMsg  *string `json:"check_msg,omitempty" `

    /*
        新增成功后分配的往来单位refEntId     */
    ParRefEntId  *string `json:"par_ref_ent_id,omitempty" `

    /*
        新增成功还是失败，true：新增成功     */
    AddSucess  *bool `json:"add_sucess,omitempty" `

}

func (s *AlibabaAlihealthDrugMscSaveentModel) SetCheckMsg(v string) *AlibabaAlihealthDrugMscSaveentModel {
    s.CheckMsg = &v
    return s
}
func (s *AlibabaAlihealthDrugMscSaveentModel) SetParRefEntId(v string) *AlibabaAlihealthDrugMscSaveentModel {
    s.ParRefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugMscSaveentModel) SetAddSucess(v bool) *AlibabaAlihealthDrugMscSaveentModel {
    s.AddSucess = &v
    return s
}
