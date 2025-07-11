package ability2940

import (
	"errors"
	"log"

	"github.com/LGYS07/topsdk"
	"github.com/LGYS07/topsdk/ability2940/request"
	"github.com/LGYS07/topsdk/ability2940/response"
	"github.com/LGYS07/topsdk/util"
)

type Ability2940 struct {
	Client *topsdk.TopClient
}

func NewAbility2940(client *topsdk.TopClient) *Ability2940 {
	return &Ability2940{client}
}

/*
查询药品信息
*/
func (ability *Ability2940) AlibabaAlihealthDrugtraceTopYljgGetdruglist(req *request.AlibabaAlihealthDrugtraceTopYljgGetdruglistRequest) (*response.AlibabaAlihealthDrugtraceTopYljgGetdruglistResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability2940 topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.drugtrace.top.yljg.getdruglist", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthDrugtraceTopYljgGetdruglistResponse{}
	if err != nil {
		log.Println("alibabaAlihealthDrugtraceTopYljgGetdruglist error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
医疗机构查询本企业上游企业出库单据信息
*/
func (ability *Ability2940) AlibabaAlihealthDrugtraceTopYljgListupout(req *request.AlibabaAlihealthDrugtraceTopYljgListupoutRequest) (*response.AlibabaAlihealthDrugtraceTopYljgListupoutResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability2940 topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.drugtrace.top.yljg.listupout", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthDrugtraceTopYljgListupoutResponse{}
	if err != nil {
		log.Println("alibabaAlihealthDrugtraceTopYljgListupout error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
新增往来单位企业
*/
func (ability *Ability2940) AlibabaAlihealthDrugYljgSaveent(req *request.AlibabaAlihealthDrugYljgSaveentRequest) (*response.AlibabaAlihealthDrugYljgSaveentResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability2940 topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.drug.yljg.saveent", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthDrugYljgSaveentResponse{}
	if err != nil {
		log.Println("alibabaAlihealthDrugYljgSaveent error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
获取企业信息
*/
func (ability *Ability2940) AlibabaAlihealthDrugYljgGetentinfonew(req *request.AlibabaAlihealthDrugYljgGetentinfonewRequest) (*response.AlibabaAlihealthDrugYljgGetentinfonewResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability2940 topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.drug.yljg.getentinfonew", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthDrugYljgGetentinfonewResponse{}
	if err != nil {
		log.Println("alibabaAlihealthDrugYljgGetentinfonew error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
获取服务截止日期
*/
func (ability *Ability2940) AlibabaAlihealthDrugtraceTopYljgServiceGetenddate(req *request.AlibabaAlihealthDrugtraceTopYljgServiceGetenddateRequest) (*response.AlibabaAlihealthDrugtraceTopYljgServiceGetenddateResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability2940 topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.drugtrace.top.yljg.service.getenddate", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthDrugtraceTopYljgServiceGetenddateResponse{}
	if err != nil {
		log.Println("alibabaAlihealthDrugtraceTopYljgServiceGetenddate error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
获取重点追溯品种明细下载URL
*/
func (ability *Ability2940) AlibabaAlihealthDrugtraceTopYljgGetkeyflagdruginfoDownloadurl(req *request.AlibabaAlihealthDrugtraceTopYljgGetkeyflagdruginfoDownloadurlRequest) (*response.AlibabaAlihealthDrugtraceTopYljgGetkeyflagdruginfoDownloadurlResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability2940 topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.drugtrace.top.yljg.getkeyflagdruginfo.downloadurl", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthDrugtraceTopYljgGetkeyflagdruginfoDownloadurlResponse{}
	if err != nil {
		log.Println("alibabaAlihealthDrugtraceTopYljgGetkeyflagdruginfoDownloadurl error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
查询药品目录信息
*/
func (ability *Ability2940) AlibabaAlihealthDrugtraceTopYljgDrugtable(req *request.AlibabaAlihealthDrugtraceTopYljgDrugtableRequest) (*response.AlibabaAlihealthDrugtraceTopYljgDrugtableResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability2940 topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.drugtrace.top.yljg.drugtable", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthDrugtraceTopYljgDrugtableResponse{}
	if err != nil {
		log.Println("alibabaAlihealthDrugtraceTopYljgDrugtable error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
零售单据上传接口
*/
func (ability *Ability2940) AlibabaAlihealthDrugtraceTopYljgUploadretail(req *request.AlibabaAlihealthDrugtraceTopYljgUploadretailRequest) (*response.AlibabaAlihealthDrugtraceTopYljgUploadretailResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability2940 topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.drugtrace.top.yljg.uploadretail", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthDrugtraceTopYljgUploadretailResponse{}
	if err != nil {
		log.Println("alibabaAlihealthDrugtraceTopYljgUploadretail error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
出入库单据上传
*/
func (ability *Ability2940) AlibabaAlihealthDrugtraceTopYljgUploadinoutbill(req *request.AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) (*response.AlibabaAlihealthDrugtraceTopYljgUploadinoutbillResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability2940 topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.drugtrace.top.yljg.uploadinoutbill", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthDrugtraceTopYljgUploadinoutbillResponse{}
	if err != nil {
		log.Println("alibabaAlihealthDrugtraceTopYljgUploadinoutbill error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
上传单据后处理状态查询
*/
func (ability *Ability2940) AlibabaAlihealthDrugtraceTopYljgQueryBillstatus(req *request.AlibabaAlihealthDrugtraceTopYljgQueryBillstatusRequest) (*response.AlibabaAlihealthDrugtraceTopYljgQueryBillstatusResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability2940 topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.drugtrace.top.yljg.query.billstatus", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthDrugtraceTopYljgQueryBillstatusResponse{}
	if err != nil {
		log.Println("alibabaAlihealthDrugtraceTopYljgQueryBillstatus error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
通过一个码，查询这个码对应的上游企业出库单的单据号
*/
func (ability *Ability2940) AlibabaAlihealthDrugtraceTopYljgQueryUpbillcode(req *request.AlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeRequest) (*response.AlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability2940 topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.drugtrace.top.yljg.query.upbillcode", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeResponse{}
	if err != nil {
		log.Println("alibabaAlihealthDrugtraceTopYljgQueryUpbillcode error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
根据单据号查询单据的详情信息【注意：查询的是本企业的单据】
*/
func (ability *Ability2940) AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetail(req *request.AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailRequest) (*response.AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability2940 topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.drugtrace.top.yljg.query.upbilldetail", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailResponse{}
	if err != nil {
		log.Println("alibabaAlihealthDrugtraceTopYljgQueryUpbilldetail error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
单码关联关系查询
*/
func (ability *Ability2940) AlibabaAlihealthDrugtraceTopYljgQueryRelation(req *request.AlibabaAlihealthDrugtraceTopYljgQueryRelationRequest) (*response.AlibabaAlihealthDrugtraceTopYljgQueryRelationResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability2940 topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.drugtrace.top.yljg.query.relation", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthDrugtraceTopYljgQueryRelationResponse{}
	if err != nil {
		log.Println("alibabaAlihealthDrugtraceTopYljgQueryRelation error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
【准备停止使用】该接口无法返回多条同名的企业信息，请对接同权限包下的“alibaba.alihealth.drug.yljg.getentinfonew”接口。通过企业名得到唯一标识ref_ent_id及企业ent_id
*/
func (ability *Ability2940) AlibabaAlihealthDrugtraceTopYljgQueryGetentinfo(req *request.AlibabaAlihealthDrugtraceTopYljgQueryGetentinfoRequest) (*response.AlibabaAlihealthDrugtraceTopYljgQueryGetentinfoResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability2940 topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.drugtrace.top.yljg.query.getentinfo", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthDrugtraceTopYljgQueryGetentinfoResponse{}
	if err != nil {
		log.Println("alibabaAlihealthDrugtraceTopYljgQueryGetentinfo error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
根据码查询码信息
*/
func (ability *Ability2940) AlibabaAlihealthDrugtraceTopYljgQueryCodedetail(req *request.AlibabaAlihealthDrugtraceTopYljgQueryCodedetailRequest) (*response.AlibabaAlihealthDrugtraceTopYljgQueryCodedetailResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability2940 topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.drugtrace.top.yljg.query.codedetail", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthDrugtraceTopYljgQueryCodedetailResponse{}
	if err != nil {
		log.Println("alibabaAlihealthDrugtraceTopYljgQueryCodedetail error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
上游出库单单据明细查询
*/
func (ability *Ability2940) AlibabaAlihealthDrugtraceTopYljgListupoutDetail(req *request.AlibabaAlihealthDrugtraceTopYljgListupoutDetailRequest) (*response.AlibabaAlihealthDrugtraceTopYljgListupoutDetailResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability2940 topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.drugtrace.top.yljg.listupout.detail", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthDrugtraceTopYljgListupoutDetailResponse{}
	if err != nil {
		log.Println("alibabaAlihealthDrugtraceTopYljgListupoutDetail error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
往来单位查询
*/
func (ability *Ability2940) AlibabaAlihealthDrugtraceTopYljgQueryListparts(req *request.AlibabaAlihealthDrugtraceTopYljgQueryListpartsRequest) (*response.AlibabaAlihealthDrugtraceTopYljgQueryListpartsResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability2940 topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.drugtrace.top.yljg.query.listparts", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthDrugtraceTopYljgQueryListpartsResponse{}
	if err != nil {
		log.Println("alibabaAlihealthDrugtraceTopYljgQueryListparts error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
根据企业唯一标识查看企业详细信息
*/
func (ability *Ability2940) AlibabaAlihealthDrugtraceTopYljgQueryGetbyrefentid(req *request.AlibabaAlihealthDrugtraceTopYljgQueryGetbyrefentidRequest) (*response.AlibabaAlihealthDrugtraceTopYljgQueryGetbyrefentidResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability2940 topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.drugtrace.top.yljg.query.getbyrefentid", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthDrugtraceTopYljgQueryGetbyrefentidResponse{}
	if err != nil {
		log.Println("alibabaAlihealthDrugtraceTopYljgQueryGetbyrefentid error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}
