var UploadImg = function () {
    var form = new Vue({
        el: "#upload_img",
        data(){
          return {
              fileList:[],
              chooseStatus:false,
              uploadStatus:false,
          }
        },
        methods: {
            handleRemove2(file, fileList) {
                console.log(file, fileList);
                this.chooseStatus = false;
                this.uploadStatus = false;
            },
            handlePreview2(file) {
                console.log("hufan preview");
                this.$notify({
                    title: '提示',
                    message: '文件已选择',
                    type: 'info'
                });
            },
            uploadImg() {
                this.$refs.upload.submit();
            },
            handlerUploadSuccess2(response, file, fileList) {
                console.log("上传完成,file:", file.name)
                this.$notify({
                    title: '提示',
                    message: '文件上传成功',
                    type: 'info'
                });
            },
            imgFileUpload(obj){
                console.log("开始上传：file:",obj.file.name);
                var fd = new FormData();
                fd.append('file',obj.file);

                var fileName = obj.file.name;
                fileName = fileName.trim();

                if (!fileName.endsWith("jpg") && !fileName.endsWith("jpeg") &&  !fileName.endsWith("png")) {
                    this.$notify({
                        title: '提示',
                        message: '仅支持 jpg/png 格式的文件',
                        type: 'warning'
                    });
                    this.fileList = [];//清空列表
                    return
                }
                var res = false;
                var message = "";
                $.ajax({
                    url:baseReqURL + "/v1/file/upload/image",
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
                this.chooseStatus = false;
                this.uploadStatus = false;
            },
            beforeAvatarUpload(file) {
                const isLt20M = file.size / 1024 / 1024 < 20;
                if (!isLt20M) {
                    this.$notify({
                        title: '错误',
                        message: '上传头像图片大小不能超过 20MB!',
                        type: 'error'
                    });
                }
                return isLt20M;
            }
        }
    });
}
