// 项目管理相关js文件

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
            var rowName = row.scheme
            var html =
                '<button type="button"  class="btn btn-primary btn-xs btndo" onclick="editScheme(\'' +
                index + '\',\'' +
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