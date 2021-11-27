// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "Swagger API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/helloworld": {
            "get": {
                "description": "helloworld",
                "tags": [
                    "helloworld"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/test": {
            "get": {
                "description": "获取项目配置",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "project"
                ],
                "summary": "GetProjectConfig",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Result"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/conf.Config"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/test/add": {
            "get": {
                "description": "添加商品",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "add",
                    "product"
                ],
                "summary": "AddRecord",
                "responses": {
                    "200": {
                        "description": "商品数据",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Result"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/model.Product"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/test/cache/get": {
            "get": {
                "description": "查缓存",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "cache"
                ],
                "summary": "GetCache",
                "parameters": [
                    {
                        "type": "string",
                        "description": "缓存Key",
                        "name": "name",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "商品数量",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Result"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "object",
                                            "additionalProperties": {
                                                "type": "string"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/test/cache/set": {
            "get": {
                "description": "设置缓存",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "cache"
                ],
                "summary": "SetCache",
                "parameters": [
                    {
                        "type": "string",
                        "description": "缓存Key",
                        "name": "name",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "boolean",
                        "description": "缓存值",
                        "name": "value",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "商品数量",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Result"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "object",
                                            "additionalProperties": {
                                                "type": "boolean"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/test/cond": {
            "get": {
                "description": "通过条件查询",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "product"
                ],
                "summary": "QueryByCond",
                "parameters": [
                    {
                        "type": "string",
                        "description": "商品名称",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "页数",
                        "name": "page",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "每页数量",
                        "name": "pageSize",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "商品数据",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Result"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "allOf": [
                                                {
                                                    "$ref": "#/definitions/response.Pageable"
                                                },
                                                {
                                                    "type": "object",
                                                    "properties": {
                                                        "rows": {
                                                            "type": "array",
                                                            "items": {
                                                                "$ref": "#/definitions/model.Product"
                                                            }
                                                        }
                                                    }
                                                }
                                            ]
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/test/count": {
            "get": {
                "description": "查询数量",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "product",
                    "count"
                ],
                "summary": "QueryCountByCond",
                "responses": {
                    "200": {
                        "description": "商品数量",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Result"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "object",
                                            "additionalProperties": {
                                                "type": "integer"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/test/ids": {
            "get": {
                "description": "通过ID列表查询",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "product"
                ],
                "summary": "QueryByIds",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "商品ID",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "商品数据",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Result"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/model.Product"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/test/update": {
            "get": {
                "description": "更新商品",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "update",
                    "product"
                ],
                "summary": "UpdateRecord",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "商品ID",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "商品数据",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Result"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "object",
                                            "additionalProperties": {
                                                "type": "boolean"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/test/{id}": {
            "get": {
                "description": "通过ID查询",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "product"
                ],
                "summary": "QueryById",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "主键",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "商品数据",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Result"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/model.Product"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/validate": {
            "post": {
                "description": "测试校验器",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "validate"
                ],
                "summary": "ValidateUser",
                "parameters": [
                    {
                        "description": "用户姓名",
                        "name": "name",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "年龄",
                        "name": "age",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "integer"
                        }
                    },
                    {
                        "description": "密码",
                        "name": "pwd",
                        "in": "body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "确认密码",
                        "name": "checkPwd",
                        "in": "body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "用户地址",
                        "name": "addr",
                        "in": "body",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/controller.Address"
                            }
                        }
                    },
                    {
                        "type": "boolean",
                        "description": "缓存值",
                        "name": "value",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "校验是否成功",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Result"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "object",
                                            "additionalProperties": {
                                                "type": "boolean"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "client.Cache": {
            "type": "object",
            "properties": {
                "connectTimeout": {
                    "description": "连接Redis超时时间, 单位: 秒",
                    "type": "integer"
                },
                "database": {
                    "description": "数据库",
                    "type": "integer"
                },
                "defaultTimeOut": {
                    "description": "默认缓存时间, 单位: 秒",
                    "type": "integer"
                },
                "freeMaxLifetime": {
                    "description": "闲置连接存活的最大时间, 单位: 分钟",
                    "type": "integer"
                },
                "host": {
                    "description": "Redis IP",
                    "type": "string"
                },
                "maxFreeConnCount": {
                    "description": "最大闲置连接数",
                    "type": "integer"
                },
                "maxOpenConnCount": {
                    "description": "最大连接数",
                    "type": "integer"
                },
                "port": {
                    "description": "Redis 端口",
                    "type": "integer"
                },
                "pwd": {
                    "description": "密码",
                    "type": "string"
                },
                "readTimeout": {
                    "description": "读取超时时间, 单位: 秒",
                    "type": "integer"
                },
                "writeTimeout": {
                    "description": "写超时时间, 单位: 秒",
                    "type": "integer"
                }
            }
        },
        "client.DB": {
            "type": "object",
            "properties": {
                "createBatchSize": {
                    "description": "批量创建数量",
                    "type": "integer"
                },
                "database": {
                    "description": "数据库名",
                    "type": "string"
                },
                "enableRawSql": {
                    "description": "打印原生SQL",
                    "type": "boolean"
                },
                "freeMaxLifetime": {
                    "description": "闲置连接最大存活时间, 单位: 分钟",
                    "type": "integer"
                },
                "host": {
                    "description": "数据库 IP",
                    "type": "string"
                },
                "location": {
                    "description": "时区",
                    "type": "string"
                },
                "maxFreeConnCount": {
                    "description": "最大闲置连接数量",
                    "type": "integer"
                },
                "maxOpenConnCount": {
                    "description": "最大连接数量",
                    "type": "integer"
                },
                "password": {
                    "description": "密码",
                    "type": "string"
                },
                "port": {
                    "description": "数据库 端口",
                    "type": "integer"
                },
                "singularTable": {
                    "description": "表复数禁用",
                    "type": "boolean"
                },
                "tablePrefix": {
                    "description": "表前缀",
                    "type": "string"
                },
                "user": {
                    "description": "用户名",
                    "type": "string"
                }
            }
        },
        "client.Log": {
            "type": "object",
            "properties": {
                "compress": {
                    "description": "是否压缩日志文件，压缩方法gzip",
                    "type": "boolean"
                },
                "enableConsole": {
                    "description": "日志是否输出到终端",
                    "type": "boolean"
                },
                "fileDir": {
                    "description": "日志文件所在目录",
                    "type": "string"
                },
                "level": {
                    "description": "日志级别",
                    "type": "string"
                },
                "localtime": {
                    "description": "是否使用本地时间，默认使用UTC时间",
                    "type": "boolean"
                },
                "maxAge": {
                    "description": "日志保留的最大天数(只保留最近多少天的日志)",
                    "type": "integer"
                },
                "maxBackups": {
                    "description": "只保留最近多少个日志文件，用于控制程序总日志的大小",
                    "type": "integer"
                },
                "maxSize": {
                    "description": "每个日志文件长度的最大大小，默认100M",
                    "type": "integer"
                },
                "rotated": {
                    "description": "日志是否被分割",
                    "type": "boolean"
                }
            }
        },
        "conf.Config": {
            "type": "object",
            "properties": {
                "application": {
                    "description": "应用名",
                    "type": "string"
                },
                "cache": {
                    "$ref": "#/definitions/client.Cache"
                },
                "db": {
                    "$ref": "#/definitions/client.DB"
                },
                "debug": {
                    "description": "debug",
                    "type": "boolean"
                },
                "env": {
                    "description": "环境标识",
                    "type": "string"
                },
                "host": {
                    "description": "地址",
                    "type": "string"
                },
                "log": {
                    "$ref": "#/definitions/client.Log"
                },
                "port": {
                    "description": "端口",
                    "type": "integer"
                },
                "shutDownTimeout": {
                    "description": "优雅重启, 接收到相关信号后, 处理请求的最长时间, 单位: 秒， 默认5s",
                    "type": "integer"
                },
                "swagger": {
                    "description": "是否启动swagger",
                    "type": "boolean"
                }
            }
        },
        "controller.Address": {
            "type": "object",
            "required": [
                "province",
                "street"
            ],
            "properties": {
                "city": {
                    "type": "string"
                },
                "province": {
                    "type": "string"
                },
                "street": {
                    "type": "string"
                }
            }
        },
        "model.Product": {
            "type": "object",
            "properties": {
                "brand": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                },
                "totalCount": {
                    "type": "integer"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "response.Pageable": {
            "type": "object",
            "properties": {
                "hasNext": {
                    "description": "HasNext 是否还有下一页",
                    "type": "boolean"
                },
                "page": {
                    "description": "Page 当前页数",
                    "type": "integer"
                },
                "pageSize": {
                    "description": "PageSize 每页数量",
                    "type": "integer"
                },
                "rows": {
                    "description": "Rows 记录"
                },
                "totalCount": {
                    "description": "TotalCount 总数量",
                    "type": "integer"
                },
                "totalPage": {
                    "description": "TotalPage 总页数",
                    "type": "integer"
                }
            }
        },
        "response.Result": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "Code 业务状态码",
                    "type": "integer",
                    "example": 0
                },
                "data": {
                    "description": "Data 数据"
                },
                "message": {
                    "description": "Message 信息",
                    "type": "string",
                    "example": "Success"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "127.0.0.1:9090",
	BasePath:    "/api",
	Schemes:     []string{},
	Title:       "CobraExample API",
	Description: "Cobra使用示例",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register("swagger", &s{})
}
