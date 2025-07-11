package domain


type AlibabaAlihealthDrugtraceTopYljgQueryCodedetailCodeStatusTypeDto struct {
    /*
        码状态（A:已激活;I:已核注;O:已核销;C:已注销;E:码不存在）     */
    CodeStatus  *string `json:"code_status,omitempty" `

}

func (s *AlibabaAlihealthDrugtraceTopYljgQueryCodedetailCodeStatusTypeDto) SetCodeStatus(v string) *AlibabaAlihealthDrugtraceTopYljgQueryCodedetailCodeStatusTypeDto {
    s.CodeStatus = &v
    return s
}
