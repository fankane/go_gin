var CreateDownloadPDFTask = function () {
    var form = new Vue({
        el: "#create_download_task",
        data(){
          return {
              fileList:[],
              percent: 0 ,
              chooseStatus:false,
              uploadStatus:false,
              total:0,
              success:0,
              failed:0,
          }
        },
        methods: {
            handleRemove(file, fileList) {
                console.log(file, fileList);
                this.chooseStatus = false;
                this.uploadStatus = false;
            },
            handlePreview(file) {
                console.log("hufan preview");
                this.$notify({
                    title: '提示',
                    message: '文件已选择',
                    type: 'info'
                });
            },
            submitUpload() {
                this.$refs.upload.submit();

            },
            handlerUploadSuccess(response, file, fileList) {
                console.log("上传完成,file:", file.name)
                this.$notify({
                    title: '提示',
                    message: '文件上传成功',
                    type: 'info'
                });
            },
            fileUpload(obj){
                console.log("开始上传：file:",obj.file.name);
                var fd = new FormData();
                fd.append('file',obj.file);

                var fileName = obj.file.name;
                fileName = fileName.trim();

                if (!fileName.endsWith("csv")) {
                    this.$notify({
                        title: '提示',
                        message: '仅支持 csv 格式的文件',
                        type: 'warning'
                    });
                    this.fileList = [];//清空列表
                    return
                }
                var res = false;
                var message = "";
                $.ajax({
                    url:"http://127.0.0.1:9002/v1/file/upload/download/url",
                    type:"post",
                    data:fd,
                    cache: false,
                    processData: false,
                    contentType: false,
                    async: false,
                    success:function(data){
                        console.log(data);
                        res = true;
                        if (data.success == false) {
                            res = false;
                            message = data.error.message;
                        }
                        console.log("successres:", res);

                    },
                    error:function(e){
                        console.log(e);
                        res = false;
                        console.log("error res:", res);
                    }
                });

                console.log("hufan res:", res);
                if (res) {
                    this.$notify({
                        title: '提示',
                        message: '文件上传成功',
                        type: 'info'
                    });
                } else  {
                    this.$notify({
                        title: '错误',
                        message: '文件上传失败:'+message,
                        type: 'error'
                    });
                }
                this.fileList = [];//清空列表
                this.chooseStatus = true;
                this.uploadStatus = true;
                $("#downloadProcess").show();
                this.percent = 1;
            },
            checkProcess() {
                var self = this;
                //检查处理进度
                var url = "http://127.0.0.1:9002/v1/file/upload/download/process";
                $.get(url, function (resp) {
                    console.log("resp:", resp);
                    if (resp.success == true) {
                        console.log("result:", resp.result);
                        self.percent = resp.result.percent;
                        self.total = resp.result.total;
                        self.success = resp.result.success;
                        self.failed = resp.result.failed;
                        console.log("temp percent:", self.percent);
                    } else {
                        self.$notify({
                            title: '提示',
                            message: '查询进度失败',
                            type: 'error'
                        });
                    }
                });
                console.log("当前 percent:", self.percent);
                console.log("当前 this.percent:", this.percent);
            }
        }
    });
}
