package domain


type AlibabaAlihealthDrugtraceTopLsydQueryCodedetailCodeStatusTypeDto struct {
    /*
        码状态（A:已激活;I:已核注;O:已核销;C:已注销;E:码不存在）     */
    CodeStatus  *string `json:"code_status,omitempty" `

}

func (s *AlibabaAlihealthDrugtraceTopLsydQueryCodedetailCodeStatusTypeDto) SetCodeStatus(v string) *AlibabaAlihealthDrugtraceTopLsydQueryCodedetailCodeStatusTypeDto {
    s.CodeStatus = &v
    return s
}
