package domain


type AlibabaAlihealthDrugtraceTopYljgDrugtableCodeResList struct {
    /*
        码前缀     */
    CodePrefix  *string `json:"code_prefix,omitempty" `

    /*
        资源码     */
    ResCode  *string `json:"res_code,omitempty" `

    /*
        层级     */
    CodeLevel  *string `json:"code_level,omitempty" `

    /*
        包装比例     */
    PkgRatio  *string `json:"pkg_ratio,omitempty" `

}

func (s *AlibabaAlihealthDrugtraceTopYljgDrugtableCodeResList) SetCodePrefix(v string) *AlibabaAlihealthDrugtraceTopYljgDrugtableCodeResList {
    s.CodePrefix = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgDrugtableCodeResList) SetResCode(v string) *AlibabaAlihealthDrugtraceTopYljgDrugtableCodeResList {
    s.ResCode = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgDrugtableCodeResList) SetCodeLevel(v string) *AlibabaAlihealthDrugtraceTopYljgDrugtableCodeResList {
    s.CodeLevel = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgDrugtableCodeResList) SetPkgRatio(v string) *AlibabaAlihealthDrugtraceTopYljgDrugtableCodeResList {
    s.PkgRatio = &v
    return s
}
