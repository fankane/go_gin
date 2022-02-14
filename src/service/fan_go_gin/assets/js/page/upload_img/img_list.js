var ImgList = function () {
    var form = new Vue({
        el: "#img_list",
        data() {
            return {
                currentPage:1,
                pageSize: 10,
                pageSizes: [5, 10, 20, 50],
                imgTotal: 30,
                previewURLs:[],
                imgInfos: [],
            }
        },
        mounted() {
            var self = this;
            //检查处理进度
            var url = baseReqURL + "/v1/file/get/imageList";
            var getBody = {
              page: 1,
              pageSize: self.pageSize,
            };
            $.get(url, getBody, function (resp) {
                if (resp.success == true) {
                    console.log("result:", resp.result);
                    self.imgTotal = resp.result.imgTotal;
                    self.imgInfos = resp.result.imgInfos;
                    self.previewURLs = resp.result.previewURLs;
                } else {
                    self.$notify({
                        title: '提示',
                        message: '查询失败',
                        type: 'error'
                    });
                }
            });
            console.log("第一次加载 total:", self.imgTotal);
        },
        methods: {
            // 翻页方法
            handleSizeChange(val) {
                console.log(`每页 ${val} 条`);
                this.currentPage = 1;
                this.pageSize = val;
                this.handleCurrentChange(1);
            },
            handleCurrentChange(val) {
                console.log(`当前页: ${val}`);
                var self = this;
                this.currentPage = val;
                var getBody = {
                    page: val,
                    pageSize: self.pageSize,
                };
                console.log(`当前size: ${self.pageSize}`);
                //检查处理进度
                var url = baseReqURL + "/v1/file/get/imageList";
                $.get(url, getBody, function (resp) {
                    console.log("resp:", resp);
                    if (resp.success == true) {
                        console.log("result:", resp.result);
                        self.imgTotal = resp.result.imgTotal;
                        self.imgInfos = resp.result.imgInfos;
                        self.previewURLs = resp.result.previewURLs;
                    } else {
                        self.$notify({
                            title: '提示',
                            message: '查询失败',
                            type: 'error'
                        });
                    }
                });
                console.log("当前 total:", self.imgTotal);
                console.log("当前this total:", this.imgTotal);
            },
        }
    });
}
