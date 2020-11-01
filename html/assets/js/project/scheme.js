// 初始化列表
function initSchemeList() {
    var schemeList = $("#scheme_list").bootstrapTable({
        // height: 100,   // 不配置，动态设置表格高度
        url: "/admin/project/scheme/all",                    // 服务端数据接口
        method: "get",              // 请求方法
        contentType: "application/x-www-form-urlencoded",
        sortOrder: "asc",           // 排序方式
        pageNumber: 1,              // 初始化加载第一页,默认第一页
        pagination: false,          // 是否分页
        pageSize: 20,               // 每页记录数量
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
        queryParams: function queryParams(params) { // 处理请求参数
            var param = {
                /* pageNumber: params.pageNumber,
                 pageSize: params.pageSize,
                 orgId: ztreeId,
                 nodeId: ztreeId,
                 citizenName: $("#fullname").val().trim(),
                 sex: $("#sex").val(),
                 age: $("#age").val().trim(),
                 identityCode: $("#idCard").val().trim(),
                 cellPhone: $("#isMobile").val().trim(),
                 adress: $("#adress").val().trim(), */
            };
            return param;
        },
        onLoadSuccess: function (data) { // 加载成功时执行, 可进行数据格式化
            result = JSON.stringify(data.data)
            logger("scheme列表数据", result)
            //$("#scheme_list").bootstrapTable('load', data.code);
            //return result;
        },

        columns: [
            {
                "field": "id",      // 对应数据结果中的字段名
                "visible": true,    // 是否显示
                "vclass": "colStyle",
                "align": "center",
                "valign": "middle",
                "title": "ID"       // 页面表格展示的title
            },
            {
                "field": "scheme",
                "visible": true,
                "vclass": "colStyle",
                "align": "center",
                "valign": "middle",
                "title": "SCHEME"
            },
            {
                "field": "status",
                "visible": true,
                "vclass": "colStyle",
                "align": "center",
                "valign": "middle",
                "title": "状态",
                //自定义方法
                formatter: function (value) {
                    switch (value) {
                        case 0:
                            return "<span style='color:orange'>待激活</span>";
                        case 1:
                            return "<span style='color:green'>使用中</span>";
                        case 2:
                            return "<span style='color:red'>已删除</span>";
                        default:
                            return "未知状态 - " + value;
                    }
                }
            },
            {
                "field": "create_time",
                "visible": true,
                "vclass": "colStyle",
                "align": "center",
                "valign": "middle",
                "title": "创建时间"
            },
            {
                "field": "modify_time",
                "visible": true,
                "vclass": "colStyle",
                "align": "center",
                "valign": "middle",
                "title": "更新时间"
            },
            {
                "field": "create_user_id",
                "visible": true,
                "align": "center",
                "valign": "middle",
                "title": "创建人"
            },
            {
                "field": "modify_user_id",
                "visible": true,
                "vclass": "colStyle",
                "align": "center",
                "valign": "middle",
                "title": "更新人"
            },
            {
                "field": 'click',
                "align": 'center',
                "class": 'colStyle',
                "valign": 'middle',
                "title": '操作',
                "align": 'center',
                formatter: function (value, row, index) {
                    //console.log(row)
                    var value = value;
                    var index = index;
                    var rowStr = row.id;
                    var html =
                        '<button type="button"  class="btn btn-primary btn-xs btndo" onclick="Edit(\'' +
                        rowStr + '\')">修改</button>'
                    html +=
                        '<button type="button"  class="btn btn-primary btn-xs btndo" onclick="Delete(\'' +
                        rowStr + '\')" >删除</button>'
                    return html;
                }
            }
        ],

        // 没数据得处理
        formatNoMatches: function () {
            return "暂无数据";
        },
    });
}

initSchemeList()


    // 渲染表格
    //$("#scheme_list").bootstrapTable(schemeList);
