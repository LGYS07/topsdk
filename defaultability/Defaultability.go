package defaultability

import (
    "log"
    "errors"
    "topsdk"
    "topsdk/defaultability/request"
    "topsdk/defaultability/response"
    "topsdk/util"
)

type Defaultability struct {
    Client *topsdk.TopClient
}

func NewDefaultability(client *topsdk.TopClient) *Defaultability{
    return &Defaultability{client}
}

/*
    批发企业上传入出库单据接口
*/
func (ability *Defaultability) AlibabaAlihealthDrugMscUploadcircubill(req *request.AlibabaAlihealthDrugMscUploadcircubillRequest) (*response.AlibabaAlihealthDrugMscUploadcircubillResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("alibaba.alihealth.drug.msc.uploadcircubill",req.ToMap(),req.ToFileMap())
    var respStruct = response.AlibabaAlihealthDrugMscUploadcircubillResponse{}
    if(err != nil){
        log.Println("alibabaAlihealthDrugMscUploadcircubill error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    通过时间段批量查询入出库单信息
*/
func (ability *Defaultability) AlibabaAlihealthDrugMscSearchbill(req *request.AlibabaAlihealthDrugMscSearchbillRequest) (*response.AlibabaAlihealthDrugMscSearchbillResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("alibaba.alihealth.drug.msc.searchbill",req.ToMap(),req.ToFileMap())
    var respStruct = response.AlibabaAlihealthDrugMscSearchbillResponse{}
    if(err != nil){
        log.Println("alibabaAlihealthDrugMscSearchbill error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    查询单据详情
*/
func (ability *Defaultability) AlibabaAlihealthDrugMscSearchbillDetail(req *request.AlibabaAlihealthDrugMscSearchbillDetailRequest) (*response.AlibabaAlihealthDrugMscSearchbillDetailResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("alibaba.alihealth.drug.msc.searchbill.detail",req.ToMap(),req.ToFileMap())
    var respStruct = response.AlibabaAlihealthDrugMscSearchbillDetailResponse{}
    if(err != nil){
        log.Println("alibabaAlihealthDrugMscSearchbillDetail error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    零头出入库单据上传
*/
func (ability *Defaultability) AlibabaAlihealthDrugMscRemnantbillUpload(req *request.AlibabaAlihealthDrugMscRemnantbillUploadRequest) (*response.AlibabaAlihealthDrugMscRemnantbillUploadResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("alibaba.alihealth.drug.msc.remnantbill.upload",req.ToMap(),req.ToFileMap())
    var respStruct = response.AlibabaAlihealthDrugMscRemnantbillUploadResponse{}
    if(err != nil){
        log.Println("alibabaAlihealthDrugMscRemnantbillUpload error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    单据处理状态查询
*/
func (ability *Defaultability) AlibabaAlihealthDrugMscBillSearchstatus(req *request.AlibabaAlihealthDrugMscBillSearchstatusRequest) (*response.AlibabaAlihealthDrugMscBillSearchstatusResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("alibaba.alihealth.drug.msc.bill.searchstatus",req.ToMap(),req.ToFileMap())
    var respStruct = response.AlibabaAlihealthDrugMscBillSearchstatusResponse{}
    if(err != nil){
        log.Println("alibabaAlihealthDrugMscBillSearchstatus error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    获取企业信息
*/
func (ability *Defaultability) AlibabaAlihealthDrugMscGetentinfonew(req *request.AlibabaAlihealthDrugMscGetentinfonewRequest) (*response.AlibabaAlihealthDrugMscGetentinfonewResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("alibaba.alihealth.drug.msc.getentinfonew",req.ToMap(),req.ToFileMap())
    var respStruct = response.AlibabaAlihealthDrugMscGetentinfonewResponse{}
    if(err != nil){
        log.Println("alibabaAlihealthDrugMscGetentinfonew error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    新增往来单位企业
*/
func (ability *Defaultability) AlibabaAlihealthDrugMscSaveent(req *request.AlibabaAlihealthDrugMscSaveentRequest) (*response.AlibabaAlihealthDrugMscSaveentResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("alibaba.alihealth.drug.msc.saveent",req.ToMap(),req.ToFileMap())
    var respStruct = response.AlibabaAlihealthDrugMscSaveentResponse{}
    if(err != nil){
        log.Println("alibabaAlihealthDrugMscSaveent error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    【准备停止使用】该接口无法返回多条同名的企业信息，请对接同权限包下的“alibaba.alihealth.drug.msc.getentinfonew”接口。根据企业名称查询企业唯一标识【ref_ent_id】和企业ID【ent_id】
*/
func (ability *Defaultability) AlibabaAlihealthDrugMscGetentinfo(req *request.AlibabaAlihealthDrugMscGetentinfoRequest) (*response.AlibabaAlihealthDrugMscGetentinfoResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("alibaba.alihealth.drug.msc.getentinfo",req.ToMap(),req.ToFileMap())
    var respStruct = response.AlibabaAlihealthDrugMscGetentinfoResponse{}
    if(err != nil){
        log.Println("alibabaAlihealthDrugMscGetentinfo error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    企业上传出入库信息
*/
func (ability *Defaultability) AlibabaAlihealthDrugMscUploadinoutbill(req *request.AlibabaAlihealthDrugMscUploadinoutbillRequest) (*response.AlibabaAlihealthDrugMscUploadinoutbillResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("alibaba.alihealth.drug.msc.uploadinoutbill",req.ToMap(),req.ToFileMap())
    var respStruct = response.AlibabaAlihealthDrugMscUploadinoutbillResponse{}
    if(err != nil){
        log.Println("alibabaAlihealthDrugMscUploadinoutbill error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    查询往来单位列表
*/
func (ability *Defaultability) AlibabaAlihealthDrugMscListparts(req *request.AlibabaAlihealthDrugMscListpartsRequest) (*response.AlibabaAlihealthDrugMscListpartsResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("alibaba.alihealth.drug.msc.listparts",req.ToMap(),req.ToFileMap())
    var respStruct = response.AlibabaAlihealthDrugMscListpartsResponse{}
    if(err != nil){
        log.Println("alibabaAlihealthDrugMscListparts error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    查询入库单据详情以及码
*/
func (ability *Defaultability) AlibabaAlihealthDrugMscBillinDetailwithcode(req *request.AlibabaAlihealthDrugMscBillinDetailwithcodeRequest) (*response.AlibabaAlihealthDrugMscBillinDetailwithcodeResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("alibaba.alihealth.drug.msc.billin.detailwithcode",req.ToMap(),req.ToFileMap())
    var respStruct = response.AlibabaAlihealthDrugMscBillinDetailwithcodeResponse{}
    if(err != nil){
        log.Println("alibabaAlihealthDrugMscBillinDetailwithcode error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    根据企业唯一标识查看企业详细信息
*/
func (ability *Defaultability) AlibabaAlihealthDrugMscGetbyrefentid(req *request.AlibabaAlihealthDrugMscGetbyrefentidRequest) (*response.AlibabaAlihealthDrugMscGetbyrefentidResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("alibaba.alihealth.drug.msc.getbyrefentid",req.ToMap(),req.ToFileMap())
    var respStruct = response.AlibabaAlihealthDrugMscGetbyrefentidResponse{}
    if(err != nil){
        log.Println("alibabaAlihealthDrugMscGetbyrefentid error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    物流企业查询货主企业信息
*/
func (ability *Defaultability) AlibabaAlihealthDrugMscSynonymauths(req *request.AlibabaAlihealthDrugMscSynonymauthsRequest) (*response.AlibabaAlihealthDrugMscSynonymauthsResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("alibaba.alihealth.drug.msc.synonymauths",req.ToMap(),req.ToFileMap())
    var respStruct = response.AlibabaAlihealthDrugMscSynonymauthsResponse{}
    if(err != nil){
        log.Println("alibabaAlihealthDrugMscSynonymauths error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    根据企业主键查看企业详细信息
*/
func (ability *Defaultability) AlibabaAlihealthDrugMscGetbyentid(req *request.AlibabaAlihealthDrugMscGetbyentidRequest) (*response.AlibabaAlihealthDrugMscGetbyentidResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("alibaba.alihealth.drug.msc.getbyentid",req.ToMap(),req.ToFileMap())
    var respStruct = response.AlibabaAlihealthDrugMscGetbyentidResponse{}
    if(err != nil){
        log.Println("alibabaAlihealthDrugMscGetbyentid error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    查询出库单据详情以及码
*/
func (ability *Defaultability) AlibabaAlihealthDrugMscBilloutDetailwithcodes(req *request.AlibabaAlihealthDrugMscBilloutDetailwithcodesRequest) (*response.AlibabaAlihealthDrugMscBilloutDetailwithcodesResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("alibaba.alihealth.drug.msc.billout.detailwithcodes",req.ToMap(),req.ToFileMap())
    var respStruct = response.AlibabaAlihealthDrugMscBilloutDetailwithcodesResponse{}
    if(err != nil){
        log.Println("alibabaAlihealthDrugMscBilloutDetailwithcodes error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    查询药品目录信息
*/
func (ability *Defaultability) AlibabaAlihealthDrugMscDrugtable(req *request.AlibabaAlihealthDrugMscDrugtableRequest) (*response.AlibabaAlihealthDrugMscDrugtableResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("alibaba.alihealth.drug.msc.drugtable",req.ToMap(),req.ToFileMap())
    var respStruct = response.AlibabaAlihealthDrugMscDrugtableResponse{}
    if(err != nil){
        log.Println("alibabaAlihealthDrugMscDrugtable error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
