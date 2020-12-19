
/**
 * 生成动态表格的配置
 * @param {JSON} columns  字段配置列表
 * @param {string} url 服务端URL
 * @param {integer} defaultPageSize 默认每页数量
 * @param {*} pagination 是否分页
 * @param {*} pageList 允许选择的一页数据量
 * @param {*} sidePagination 分页方式(client/server)
 * @param {*} dataField 数据字段
 * @param {*} queryParams 获取请求参数的过滤函数
 * @param {*} onLoadSuccess 数据加载成功之后的回调函数
 */
function getTableConfig(columns, url, defaultPage, defaultPageSize, pagination, pageList, sidePagination, dataField, queryParams, onLoadSuccess) {
    if (pagination === undefined) {
        // 默认需要分页
        pagination = true
    }
    if (defaultPage === undefined || defaultPage <= 0) {
        logger("未定义页码,使用默认值 0")
        defaultPage = 1
    }
    if (undefined === defaultPageSize || 0 >= defaultPageSize) {
        logger("未定义每页数量,使用默认值 10")
        defaultPageSize = 10
    }
    if (undefined === pageList) {
        // 默认的页码列表
        pageList = [10, 20, 25, 50, 100]
    }
    if (undefined === sidePagination || ("client" !== sidePagination && "server" !== sidePagination)) {
        sidePagination = "server"
    }
    if (undefined === dataField || dataField.length === 0) {
        dataField = "data"
    }
    if (undefined === queryParams) {
        queryParams = { "page": defaultPage, "size": defaultPageSize }
    }

    if (undefined === onLoadSuccess) {
        onLoadSuccess = function (data) {
            logger(url + "接口数据加载成功, 数据列表", data.data)
        }
    }

    logger("表格分页,页码与大小 : ", queryParams, defaultPageSize, defaultPage)
    return {
        // height: 100,   // 不配置，动态设置表格高度
        url: url,                    // 服务端数据接口
        method: "get",              // 请求方法
        contentType: "application/x-www-form-urlencoded",
        sortOrder: "asc",           // 排序方式
        pageNumber: defaultPage,              // 初始化加载第一页,默认第一页
        pagination: pagination,          // 是否分页
        pageSize: defaultPageSize,               // 每页记录数量
        pageList: [10, 25, 50, 100],//可供选择的每页的行数（*）
        sidePagination: "server",   //分页方式：client客户端分页，server服务端分页（*）
        toolbar: '#toolbar',        //Bstable工具导航条
        dataField: "data",          //数据节点   默认为rows 改成data后需要后台返回的数组为data
        minimumCountColumns: 2,     // 允许最少的列数
        showExport: false,          // 是否显示导出
        exportDataType: "basic",    // 导出数据类型 basic/all/selected
        align: "center",            // 对齐方式
        clickToSelect: true,        // 是否启用点击选中行
        uniqueId: "id",             // 每一行的唯一标识，一般为主键列
        showToggle: false,           //是否显示详细视图和列表视图的切换按钮
        cardView: false,            //是否显示详细视图
        detailView: false,          //是否显示父子表
        queryParams: queryParams,   // 处理请求参数(匿名函数)

        onLoadSuccess: function (data) { // 加载成功时执行, 可进行数据格式化

            //$("#scheme_list").bootstrapTable('load', data.code);
            //return result;
        },

        columns: columns,

        // 没数据得处理
        formatNoMatches: function () {
            return "暂无数据";
        },
    }
}

/**
 * 公共的刷新表格方法
 *
 * @param {string} tableID
 * @param {int} page
 * @param {int} size
 * @param {Map} params
 */
function refreshTableEvent(tableID, page, size, params) {
    if (undefined === params) {
        params = {}
    }
    params["page"] = page
    params["size"] = size
    $(tableID).bootstrapTable("refreshOptions", {
        "queryParams": params
    })
}
