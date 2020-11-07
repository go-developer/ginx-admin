// 初始化列表
function initSchemeList(page, pageSize) {
    columns = [
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
                var rowName = row.scheme
                var html =
                    '<button type="button"  class="btn btn-primary btn-xs btndo" onclick="editScheme(\'' +
                    rowID + '\', \'' +
                    rowName + '\', \'' +
                    rowStatus + '\')">修改</button>'
                html +=
                    '<button type="button"  class="btn btn-primary btn-xs btndo" onclick="Delete(\'' +
                    rowID + '\')" >删除</button>'
                return html;
            }
        }
    ];

    tableConfig = getTableConfig(columns, "/admin/project/scheme/all", page, pageSize, true, [10, 20, 25, 50, 100], "server", "data", undefined)
    $("#scheme_list").bootstrapTable(tableConfig);
}

initSchemeList(1, 10)
// alert($("#scheme_list").bootstrapTable("getOptions").pageNumber);
// alert($("#scheme_list").bootstrapTable("getOptions").pageSize);
$('#scheme_list').on('page-change.bs.table', function (e, number, size) {
    page = $("#scheme_list").bootstrapTable("getOptions").pageNumber;
    pageSize = $("#scheme_list").bootstrapTable("getOptions").pageSize;
    initSchemeList(page, pageSize);
    // alert('切换页事件 --- 当前页数：第' + number + "页，每页显示数量" + size + "条");
});

/**
 * 编辑scheme
 * @param {integer} schemeID
 * @param {string} currentSchemeName
 * @param {integer} currentSchemeStatus
 */
function editScheme(schemeID, currentSchemeName, currentSchemeStatus) {
    $("#edit_scheme_id").val(schemeID)
    $("#edit_scheme_status").val(currentSchemeStatus)
    $("#edit_scheme_name").val(currentSchemeName)
    // 展示模态框
    $('#editSchemeModal').modal({
        show: true,
        backdrop: 'static'
    })
}

/**
 * 编辑按钮点击事件
 */
$("#edit_scheme_button").click(function () {
    var schemeID = Number($("#edit_scheme_id").val())
    var schemeName = $("#edit_scheme_name").val()
    var schemeStatus = $("#edit_scheme_status").val()
    request(methodPost, "/admin/project/scheme/update", { "scheme_id": schemeID, "scheme": schemeName, "status": schemeStatus }, "json", contentTypeJson)
});
