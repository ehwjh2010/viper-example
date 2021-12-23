package controller

import (
	"fmt"
	"github.com/ehwjh2010/viper-example/config"
	"github.com/ehwjh2010/viper-example/internal/dao"
	"github.com/ehwjh2010/viper-example/internal/model"
	"github.com/ehwjh2010/viper-example/internal/proxy/cache"
	"github.com/ehwjh2010/viper/db/rdb"
	"github.com/ehwjh2010/viper/extend/ginext"
	"github.com/ehwjh2010/viper/extend/ginext/response"
	"github.com/ehwjh2010/viper/log"
	"github.com/ehwjh2010/viper/types"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// Helloworld 测试接口
// @Title helloworld
// @Description helloworld
// @Tags helloworld
// @Success 200 {string} helloworld
// @Router	/helloworld [get]
func Helloworld(c *gin.Context) {
	c.JSON(http.StatusOK, "helloworld")
}

// GetProjectConfig 获取项目配置
// @Summary GetProjectConfig
// @Description 获取项目配置
// @Accept json
// @Tags project
// @Router /test [get]
// @Success 200 {object} types.Result{data=config.Config}
func GetProjectConfig(c *gin.Context) {
	log.Info("你好")
	response.Success(c, config.Conf)
}

// AddRecord 添加商品
// @Summary AddRecord
// @Description 添加商品
// @Accept json
// @Produce json
// @Tags add,product
// @Success 200 {object} types.Result{data=model.Product} "商品数据"
// @Router /test/add [get]
func AddRecord(c *gin.Context) {

	product := model.Product{
		Name:       "Cake",
		Price:      1000,
		TotalCount: 5000,
		Brand:      types.NewStr("华为"),
	}

	err := dao.DBClient.AddRecord(&product)

	if err != nil {
		log.Error(err.Error())
		response.Fail(c, 1000, "插入失败")
		return
	}

	response.Success(c, product)
}

// UpdateRecord 更新商品
// @Summary UpdateRecord
// @Description 更新商品
// @Accept json
// @Produce json
// @Tags update,product
// @Param id query int true "商品ID"
// @Success 200 {object} types.Result{data=map[string]bool} "商品数据"
// @Router /test/update [get]
func UpdateRecord(c *gin.Context) {
	product := model.NewProduct()

	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		response.InvalidRequest(c, "ID必须为整数")
		return
	}

	product.TotalCount = 99
	product.Price = 9900
	err = dao.DBClient.UpdateById(product.TableName(), int64(id), product)

	if err != nil {
		response.Success(c, map[string]bool{"ok": false})
	} else {
		response.Success(c, map[string]bool{"ok": true})
	}
}

// QueryByIds 通过ID列表查询
// @Summary QueryByIds
// @Description 通过ID列表查询
// @Accept json
// @Produce json
// @Tags product
// @Param id query int true "商品ID"
// @Success 200 {object} types.Result{data=[]model.Product} "商品数据"
// @Router /test/ids [get]
func QueryByIds(c *gin.Context) {

	id := c.Query("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		response.InvalidRequest(c, "Id must be int")
		return
	}

	var product []model.Product

	_, err = dao.DBClient.QueryByIds([]int64{int64(idInt)}, &product)
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
// @Success 200 {object} types.Result{data=model.Product} "商品数据"
// @Router /test/{id} [get]
func QueryById(c *gin.Context) {

	id := c.Param("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		response.InvalidRequest(c, "Id must be int")
		return
	}

	product := model.NewProduct()

	exist, err := dao.DBClient.QueryById(int64(idInt), &product)
	if err != nil {
		response.Fail(c, 2000, "系统错误")
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
// @Success 200 {object} types.Result{data=types.Pageable{rows=[]model.Product}} "商品数据"
// @Router /test/cond [get]
func QueryByCond(c *gin.Context) {
	name := c.Query("name")

	//page, _ := strconv.Atoi(c.Query("page"))
	//pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	cond := rdb.NewQueryCondition()

	//cond.SetPage(page).SetPageSize(pageSize).AddSort(rdb.NewDescOrder("price"))

	//cond.SetTotalCount(true)

	//w := rdb.NewEqWhere("name", name).Or(rdb.NewEqWhere("name", "banana")).Or(rdb.NewEqWhere("price", 30))
	//
	//cond.AddWhere(w)

	cond.AddWhere(rdb.NewLeftLikeWhere("name", name))

	var products []model.Product

	if _, err := dao.DBClient.Query(model.NewProduct().TableName(), cond, &products); err == nil {
		response.Success(c, products)
	} else {
		response.Success(c, nil)
	}
}

// QueryCountByCond 查询数量
// @Summary QueryCountByCond
// @Description 查询数量
// @Accept json
// @Produce json
// @Tags product,count
// @Success 200 {object} types.Result{data=map[string]int} "商品数量"
// @Router /test/count [get]
func QueryCountByCond(c *gin.Context) {
	product := model.NewProduct()

	cond := rdb.NewQueryCondition()

	//cond.AddWhere(rdb.NewEqWhere("total_count", 10))
	//cond.AddWhere(rdb.NewEqWhere("price", 30))

	count, err := dao.DBClient.QueryCount(product.TableName(), cond)

	if err != nil {
		response.Fail(c, 90000, "查询结果失败")
		return
	} else {
		response.Success(c, map[string]int64{"count": count})
	}
}

//GetCache 查缓存
//@Summary GetCache
//@Description 查缓存
//@Accept json
//@Produce json
//@Tags cache
//@Param name query string true "缓存Key"
//@Success 200 {object} types.Result{data=map[string]string} "商品数量"
//@Router /test/cache/get [get]
func GetCache(c *gin.Context) {
	name := c.Query("name")
	start, _ := strconv.Atoi(c.Query("start"))
	end, _ := strconv.Atoi(c.Query("end"))

	v, err := cache.RedisClient.ZRangeWithScore(name, start, end, true)

	if err != nil {
		log.Error(err.Error())
		response.Fail(c, 30000, "操作redis失败")
		return
	}

	response.Success(c, v)

}

//SetCache 设置缓存
//@Summary SetCache
//@Description 设置缓存
//@Accept json
//@Produce json
//@Tags cache
//@Param name query string true "缓存Key"
//@Param value query bool true "缓存值"
//@Success 200 {object} types.Result{data=map[string]bool} "商品数量"
//@Router /test/cache/set [get]
func SetCache(c *gin.Context) {
	name := c.Query("name")
	score, _ := strconv.ParseFloat(c.Query("score"), 10)
	value := c.Query("value")

	err := cache.RedisClient.ZSet(name, score, value)

	if err != nil {
		log.Error(err.Error())
		response.Fail(c, 30000, "操作redis失败")
		return
	}

	response.Success(c, true)
}

type User struct {
	Name     string    `json:"name" binding:"required,gte=2"`
	Age      int       `json:"age" binding:"required,gt=0,lte=200"`
	Addr     []Address `json:"addr" binding:"dive"`
	Pwd      string    `json:"pwd" binding:"gte=5"`
	CheckPwd string    `json:"checkPwd" binding:"required_with=Pwd,eqfield=Pwd"`
}

type Address struct {
	Province string `json:"province" binding:"required"`
	City     string `json:"city" binding:"city"`
	Street   string `json:"street" binding:"required"`
}

// ValidateUser	测试校验器
// @Summary 	ValidateUser
// @Description 测试校验器
// @Accept 		json
// @Produce 	json
// @Tags 		validate
// @Param 		user body User true "用户姓名"
// @Success 	200 {object} types.Result{data=map[string]bool} "校验是否成功"
// @Router 		/validate [post]
func ValidateUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		response.InvalidRequest(c, ginext.Translate(err))
		return
	}

	log.Info(fmt.Sprintf("%+v", user))

	response.Success(c, map[string]bool{"ok": true})
}
