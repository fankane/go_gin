var APP = {
    init: function (config) {
        this.config = config || {
            webApiBaseURL: "http://127.0.0.1:9002"
            // webApiBaseURL: "https://apiconn-support.dev.klook.io"
        };
        this.config.vueBootstrapMap = {
            '/assets/page/download_pdf/create_download_task.html': CreateDownloadPDFTask,

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
        webApiBaseURL: "http://127.0.0.1:9002"
        // webApiBaseURL: "https://apiconn-support.dev.klook.io"
    });
}
