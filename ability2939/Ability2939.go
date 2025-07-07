package ability2939

import (
	"errors"
	"log"

	"github.com/LGYS07/topsdk"
	"github.com/LGYS07/topsdk/ability2939/request"
	"github.com/LGYS07/topsdk/ability2939/response"
	"github.com/LGYS07/topsdk/util"
)

type Ability2939 struct {
	Client *topsdk.TopClient
}

func NewAbility2939(client *topsdk.TopClient) *Ability2939 {
	return &Ability2939{client}
}

/*
零售单据上传接口
*/
func (ability *Ability2939) AlibabaAlihealthDrugtraceTopLsydUploadretail(req *request.AlibabaAlihealthDrugtraceTopLsydUploadretailRequest) (*response.AlibabaAlihealthDrugtraceTopLsydUploadretailResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability2939 topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.drugtrace.top.lsyd.uploadretail", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthDrugtraceTopLsydUploadretailResponse{}
	if err != nil {
		log.Println("alibabaAlihealthDrugtraceTopLsydUploadretail error", err)
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
func (ability *Ability2939) AlibabaAlihealthDrugtraceTopLsydUploadinoutbill(req *request.AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) (*response.AlibabaAlihealthDrugtraceTopLsydUploadinoutbillResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability2939 topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.drugtrace.top.lsyd.uploadinoutbill", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthDrugtraceTopLsydUploadinoutbillResponse{}
	if err != nil {
		log.Println("alibabaAlihealthDrugtraceTopLsydUploadinoutbill error", err)
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
func (ability *Ability2939) AlibabaAlihealthDrugtraceTopLsydQueryBillstatus(req *request.AlibabaAlihealthDrugtraceTopLsydQueryBillstatusRequest) (*response.AlibabaAlihealthDrugtraceTopLsydQueryBillstatusResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability2939 topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.drugtrace.top.lsyd.query.billstatus", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthDrugtraceTopLsydQueryBillstatusResponse{}
	if err != nil {
		log.Println("alibabaAlihealthDrugtraceTopLsydQueryBillstatus error", err)
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
func (ability *Ability2939) AlibabaAlihealthDrugtraceTopLsydQueryUpbillcode(req *request.AlibabaAlihealthDrugtraceTopLsydQueryUpbillcodeRequest) (*response.AlibabaAlihealthDrugtraceTopLsydQueryUpbillcodeResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability2939 topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.drugtrace.top.lsyd.query.upbillcode", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthDrugtraceTopLsydQueryUpbillcodeResponse{}
	if err != nil {
		log.Println("alibabaAlihealthDrugtraceTopLsydQueryUpbillcode error", err)
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
func (ability *Ability2939) AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetail(req *request.AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailRequest) (*response.AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability2939 topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.drugtrace.top.lsyd.query.upbilldetail", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailResponse{}
	if err != nil {
		log.Println("alibabaAlihealthDrugtraceTopLsydQueryUpbilldetail error", err)
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
func (ability *Ability2939) AlibabaAlihealthDrugtraceTopLsydQueryRelation(req *request.AlibabaAlihealthDrugtraceTopLsydQueryRelationRequest) (*response.AlibabaAlihealthDrugtraceTopLsydQueryRelationResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability2939 topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.drugtrace.top.lsyd.query.relation", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthDrugtraceTopLsydQueryRelationResponse{}
	if err != nil {
		log.Println("alibabaAlihealthDrugtraceTopLsydQueryRelation error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
零售药店往来单位新增
*/
func (ability *Ability2939) AlibabaAlihealthDrugLsydSaveent(req *request.AlibabaAlihealthDrugLsydSaveentRequest) (*response.AlibabaAlihealthDrugLsydSaveentResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability2939 topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.drug.lsyd.saveent", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthDrugLsydSaveentResponse{}
	if err != nil {
		log.Println("alibabaAlihealthDrugLsydSaveent error", err)
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
func (ability *Ability2939) AlibabaAlihealthDrugtraceTopLsydQueryCodedetail(req *request.AlibabaAlihealthDrugtraceTopLsydQueryCodedetailRequest) (*response.AlibabaAlihealthDrugtraceTopLsydQueryCodedetailResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability2939 topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.drugtrace.top.lsyd.query.codedetail", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthDrugtraceTopLsydQueryCodedetailResponse{}
	if err != nil {
		log.Println("alibabaAlihealthDrugtraceTopLsydQueryCodedetail error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
【准备停止使用】该接口无法返回多条同名的企业信息，请对接同权限包下的“alibaba.alihealth.drug.lsyd.getentinfonew”接口。通过企业名得到唯一标识ref_ent_id及企业ent_id
*/
func (ability *Ability2939) AlibabaAlihealthDrugtraceTopLsydQueryGetentinfo(req *request.AlibabaAlihealthDrugtraceTopLsydQueryGetentinfoRequest) (*response.AlibabaAlihealthDrugtraceTopLsydQueryGetentinfoResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability2939 topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.drugtrace.top.lsyd.query.getentinfo", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthDrugtraceTopLsydQueryGetentinfoResponse{}
	if err != nil {
		log.Println("alibabaAlihealthDrugtraceTopLsydQueryGetentinfo error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
零售药店查询本企业上游企业出库单据信息
*/
func (ability *Ability2939) AlibabaAlihealthDrugtraceTopLsydListupout(req *request.AlibabaAlihealthDrugtraceTopLsydListupoutRequest) (*response.AlibabaAlihealthDrugtraceTopLsydListupoutResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability2939 topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.drugtrace.top.lsyd.listupout", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthDrugtraceTopLsydListupoutResponse{}
	if err != nil {
		log.Println("alibabaAlihealthDrugtraceTopLsydListupout error", err)
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
func (ability *Ability2939) AlibabaAlihealthDrugtraceTopLsydListupoutDetail(req *request.AlibabaAlihealthDrugtraceTopLsydListupoutDetailRequest) (*response.AlibabaAlihealthDrugtraceTopLsydListupoutDetailResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability2939 topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.drugtrace.top.lsyd.listupout.detail", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthDrugtraceTopLsydListupoutDetailResponse{}
	if err != nil {
		log.Println("alibabaAlihealthDrugtraceTopLsydListupoutDetail error", err)
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
func (ability *Ability2939) AlibabaAlihealthDrugLsydGetentinfonew(req *request.AlibabaAlihealthDrugLsydGetentinfonewRequest) (*response.AlibabaAlihealthDrugLsydGetentinfonewResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability2939 topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.drug.lsyd.getentinfonew", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthDrugLsydGetentinfonewResponse{}
	if err != nil {
		log.Println("alibabaAlihealthDrugLsydGetentinfonew error", err)
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
func (ability *Ability2939) AlibabaAlihealthDrugtraceTopLsydGetkeyflagdruginfoDownloadurl(req *request.AlibabaAlihealthDrugtraceTopLsydGetkeyflagdruginfoDownloadurlRequest) (*response.AlibabaAlihealthDrugtraceTopLsydGetkeyflagdruginfoDownloadurlResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability2939 topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.drugtrace.top.lsyd.getkeyflagdruginfo.downloadurl", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthDrugtraceTopLsydGetkeyflagdruginfoDownloadurlResponse{}
	if err != nil {
		log.Println("alibabaAlihealthDrugtraceTopLsydGetkeyflagdruginfoDownloadurl error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
获取企业服务截止时间
*/
func (ability *Ability2939) AlibabaAlihealthDrugtraceTopLsydServiceGetenddate(req *request.AlibabaAlihealthDrugtraceTopLsydServiceGetenddateRequest) (*response.AlibabaAlihealthDrugtraceTopLsydServiceGetenddateResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability2939 topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.drugtrace.top.lsyd.service.getenddate", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthDrugtraceTopLsydServiceGetenddateResponse{}
	if err != nil {
		log.Println("alibabaAlihealthDrugtraceTopLsydServiceGetenddate error", err)
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
func (ability *Ability2939) AlibabaAlihealthDrugtraceTopLsydQueryListparts(req *request.AlibabaAlihealthDrugtraceTopLsydQueryListpartsRequest) (*response.AlibabaAlihealthDrugtraceTopLsydQueryListpartsResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability2939 topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.drugtrace.top.lsyd.query.listparts", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthDrugtraceTopLsydQueryListpartsResponse{}
	if err != nil {
		log.Println("alibabaAlihealthDrugtraceTopLsydQueryListparts error", err)
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
func (ability *Ability2939) AlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentid(req *request.AlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidRequest) (*response.AlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability2939 topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.drugtrace.top.lsyd.query.getbyrefentid", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidResponse{}
	if err != nil {
		log.Println("alibabaAlihealthDrugtraceTopLsydQueryGetbyrefentid error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
查询码是否激活
*/
func (ability *Ability2939) AlibabaAlihealthDrugtraceTopLsydQuerycodeactive(req *request.AlibabaAlihealthDrugtraceTopLsydQuerycodeactiveRequest) (*response.AlibabaAlihealthDrugtraceTopLsydQuerycodeactiveResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability2939 topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.drugtrace.top.lsyd.querycodeactive", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthDrugtraceTopLsydQuerycodeactiveResponse{}
	if err != nil {
		log.Println("alibabaAlihealthDrugtraceTopLsydQuerycodeactive error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}
