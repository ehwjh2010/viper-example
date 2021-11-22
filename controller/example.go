package controller

import (
	"github.com/ehwjh2010/cobra-example/conf"
	"github.com/ehwjh2010/cobra-example/resource"
	"github.com/ehwjh2010/cobra-example/resource/model"
	"github.com/ehwjh2010/cobra/db/rdb"
	"github.com/ehwjh2010/cobra/http/response"
	"github.com/ehwjh2010/cobra/log"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// Helloworld 测试接口
// @Title helloworld
// @Description helloworld
// @Tags test
// @Success 200 {string} helloworld
// @Router	/helloworld [get]
func Helloworld(c *gin.Context)  {
	c.JSON(http.StatusOK,"helloworld")
}

// GetProjectConfig 获取项目配置
// @Summary GetProjectConfig
// @Description 获取项目配置
// @Accept json
// @Tags project,config
// @Router /test [get]
// @Success 200 {object} response.Result{data=conf.Config}
func GetProjectConfig(c *gin.Context) {
	log.Info("你好")
	response.Success(c, conf.Conf)
}


// AddRecord 添加商品
// @Summary AddRecord
// @Description 添加商品
// @Accept json
// @Produce json
// @Tags add,product
// @Success 200 {object} response.Result{data=model.Product} "商品数据"
// @Router /test/add [get]
func AddRecord(c *gin.Context) {

	product := model.Product{
		Name:       "Cake",
		Price:      30,
		TotalCount: 10000,
	}

	err := resource.DBClient.AddRecord(&product)

	if err != nil {
		//util.Fail(c, util.ResultWithCode(10000), util.ResultWithMessage("Insert failed!"))
		return
	}

	response.Success(c, product)
}

//
//func UpdateRecord(c *gin.Context) {
//	product := model2.NewProduct()
//
//	id, err := strconv.Atoi(c.Param("id"))
//	if err != nil {
//		//util.Fail(c, util.ResultWithCode(1000))
//		return
//	}
//
//	product.TotalCount = 99
//	product.Price = 9900
//	err = dao2.DBClient.UpdateById(product.TableName(), int64(id), product)
//
//	if err != nil {
//		//util.Fail(
//		//	c,
//		//	util.ResultWithCode(2000),
//		//	util.ResultWithMessage(fmt.Sprintf("Update failed, %+v\n", err)))
//		return
//	}
//
//	util.Success(c, product)
//}
//

// QueryByIds 通过ID列表查询
// @Summary QueryByIds
// @Description 通过ID列表查询
// @Accept json
// @Produce json
// @Tags product
// @Param id query int true "商品ID"
// @Success 200 {object} response.Result{data=[]model.Product} "商品数据"
// @Router /test/ids [get]
func QueryByIds(c *gin.Context) {

	id := c.Query("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		response.InvalidRequest(c, "Id must be int")
		return
	}

	var product []model.Product

	_, err = resource.DBClient.QueryByIds([]int64{int64(idInt)}, &product)
	if err != nil {
		//util.Fail(c, util.ResultWithCode(1000))
		return
	}

	response.Success(c, product)
}

// QueryById 通过ID查询
// @Summary QueryById
// @Description 通过ID查询
// @Accept json
// @Produce json
// @Tags product
// @Param id path int true "主键"
// @Success 200 {object} response.Result{data=model.Product} "商品数据"
// @Router /test/{id} [get]
func QueryById(c *gin.Context) {

	id := c.Param("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		response.InvalidRequest(c, "Id must be int")
		return
	}

	product := model.NewProduct()

	exist, err := resource.DBClient.QueryById(int64(idInt), &product)
	if err != nil {
		//util.Fail(c, util.ResultWithCode(1000))
		return
	}

	if !exist {
		response.Success(c, nil)
		return
	}

	response.Success(c, product)
}

// QueryByCond 通过条件查询
// @Summary QueryByCond
// @Description 通过条件查询
// @Accept json
// @Produce json
// @Tags product
// @Param name query string false "商品名称"
// @Param page query int true "页数"
// @Param pageSize query int true "每页数量"
// @Success 200 {object} response.Result{data=response.Pageable{rows=[]model.Product}} "商品数据"
// @Router /test/cond [get]
func QueryByCond(c *gin.Context) {
	names := c.QueryArray("name")

	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	cond := rdb.NewQueryCondition()

	cond.SetPage(page).SetPageSize(pageSize).AddSort(rdb.NewDescOrder("price"))

	cond.SetTotalCount(true)

	//cond.AddWhere(rdb.NewNotEqWhere("total_count", 90))

	if len(names) > 0 {
		cond.AddWhere(rdb.NewInWhere("name", names))
	}

	var products []model.Product

	count, _ := resource.DBClient.Query(model.NewProduct().TableName(), cond, &products)

	response.Success(c, response.NewPageable(
		products,
		response.PageableWithPage(page),
		response.PageableWithPageSize(pageSize),
		response.PageableWithTotalCount(count)),
	)
}
//
////QueryCountByCond 查询数量
//func QueryCountByCond(c *gin.Context) {
//	product := model2.NewProduct()
//
//	cond := util.NewQueryCondition()
//
//	cond.AddWhere(util.NewEqWhere("total_count", 10))
//	cond.AddWhere(util.NewEqWhere("price", 30))
//
//	count, err := dao2.DBClient.QueryCount(product.TableName(), cond)
//
//	if err != nil {
//		//util.Fail(c, util.ResultWithCode(991111))
//		return
//	}
//
//	util.Success(c, map[string]int64{"count": count})
//}
//
////QueryByCache 查缓存
//func QueryByCache(c *gin.Context) {
//	name := c.Param("name")
//
//	nameValue, err := dao2.CacheClient.GetString(name)
//
//	if err != nil {
//		//util.Fail(c, util.ResultWithCode(1000))
//		return
//	}
//
//	util.Success(c, map[string]types.NullString{"name": nameValue})
//
//}
//
////SetJob 查缓存
//func SetJob(c *gin.Context) {
//	job := c.Param("job")
//
//	err := dao2.CacheClient.Set("job", job, 300)
//
//	if err != nil {
//		//util.Fail(c, util.ResultWithCode(1000))
//		return
//	}
//
//	util.Success(c, map[string]bool{"ok": true})
//
//}
//
////GetJob 查缓存
//func GetJob(c *gin.Context) {
//	job, err := dao2.CacheClient.GetString("job")
//
//	if err != nil {
//		//util.Fail(c, util.ResultWithCode(1000))
//		return
//	}
//
//	util.Success(c, map[string]types.NullString{"job": job})
//
//}
