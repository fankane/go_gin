var APP = {
    init: function (config) {
        this.config = config || {
            webApiBaseURL: baseReqURL
            // webApiBaseURL: "https://apiconn-support.dev.klook.io"
        };
        this.config.vueBootstrapMap = {
            '/assets/page/download_pdf/create_download_task.html': CreateDownloadPDFTask,
            '/assets/page/upload_img/upload_img.html': UploadImg,
            '/assets/page/upload_img/img_list.html': ImgList,
        };
    },
}

function loadhtmlToContainer(html) {
    $('#main-container').load(html, function () {
        var bootstrapFunc = APP.config.vueBootstrapMap[html];
        if (bootstrapFunc) {
            bootstrapFunc();
        } else {
            console.log(html + " not found");
        }
    });
    window.location.hash = html;
}

window.onload = function (en) {
    APP.init({
        debug:true,
        webApiBaseURL: baseReqURL
        // webApiBaseURL: "https://apiconn-support.dev.klook.io"
    });
}
