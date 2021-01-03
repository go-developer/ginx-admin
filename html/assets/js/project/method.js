methodColumns = [
    {
        "field": "id",      // 对应数据结果中的字段名
        "visible": true,    // 是否显示
        "vclass": "colStyle",
        "align": "center",
        "valign": "middle",
        "title": "ID",       // 页面表格展示的title
        //自定义方法
        formatter: function (value) {
            return "<a target='_blank' href='/page/project/method_detail.html#" + value + "'>" + value + "</a>";
        }
    },
    {
        "field": "method",
        "visible": true,
        "vclass": "colStyle",
        "align": "center",
        "valign": "middle",
        "title": "method"
    },
    {
        "field": "status",
        "visible": true, "vclass": "colStyle",
        "align": "center",
        "valign": "middle",
        "title": "状态",
        "sortable": true,
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
        "title": "创建时间",
        "sortable": true,
    },
    {
        "field": "modify_time",
        "visible": true,
        "vclass": "colStyle",
        "align": "center",
        "valign": "middle",
        "title": "更新时间",
        "sortable": true,
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
            var rowID = row.id;
            var rowStatus = row.status
            var rowName = row.method
            var html =
                '<button type="button"  class="btn btn-primary btn-xs btndo" onclick="editMethod(\'' +
                index + '\',\'' +
                rowID + '\', \'' +
                rowName + '\', \'' +
                rowStatus + '\')">修改</button>'
            //html +=
            //    '<button type="button"  class="btn btn-primary btn-xs btndo" onclick="Delete(\'' +
            //    rowID + '\')" >删除</button>'
            return html;
        }
    }
];
let methodTableID = "#method_list"
let defaultPage = 1
let defaultSize = 10
let methodTableConfig = getTableConfig(methodColumns, "/admin/project/method/list", defaultPage, defaultSize, true, [10, 20, 25, 50, 100], "server", "data", undefined)
$("#method_list").bootstrapTable(methodTableConfig);

// 分页页码点击事件
$('#method_list').on('page-change.bs.table', function (e, number, size) {
    // alert(123)
    //page = $("#method_list").bootstrapTable("getOptions").pageNumber;
    //size = $("#method_list").bootstrapTable("getOptions").pageSize;
    refreshTableEvent(methodTableID, number, size, undefined);
    // alert('切换页事件 --- 当前页数：第' + number + "页，每页显示数量" + size + "条");
});

/**
 * 编辑method
 * @param {integer} methodID
 * @param {string} currentMethodName
 * @param {integer} currentMethodStatus
 */
function editMethod(methodID, currentMethodName, currentMethodStatus) {
    $("#edit_method_id").val(methodID)
    $("#edit_method_status").val(currentMethodStatus)
    $("#edit_method_name").val(currentMethodName)
    // 展示模态框
    $('#editMethodModal').modal({
        show: true,
        backdrop: 'static'
    })
}

/**
 * 编辑按钮点击事件
 */
$("#edit_method_button").click(function () {
    var methodID = Number($("#edit_method_id").val())
    var methodName = $("#edit_method_name").val()
    var methodStatus = $("#edit_method_status").val()
    request(
        methodPost,
        "/admin/project/method/update",
        { "method_id": methodID, "method": methodName, "status": methodStatus },
        "json",
        contentTypeJson,
        function (data) {
            // 关闭模态框
            $('#editMethodModal').modal("hide")
            // 前台更新数据
            refreshTableEvent(methodmethodTableID, currentPage, currentPageSize, undefined)
        }
    )
});
