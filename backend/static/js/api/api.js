// js/api.js
const api = {
    baseUrl: '', // 可选，比如 '/api' 或完整地址
    tokenKey: 'access_token',

    // 通用请求（JSON）
    request(method, url, data = null) {
        return new Promise((resolve, reject) => {
            $.ajax({
                url: this.baseUrl + url,
                type: method,
                contentType: "application/json",
                dataType: "json",
                data: data ? JSON.stringify(data) : null,
                headers: this._getAuthHeader(),
                success: res => {
                    if (res.code === 0) {
                        resolve(res.data);
                    } else {
                        reject(res.msg || "请求失败");
                    }
                },
                error: xhr => {
                    let msg = xhr.responseJSON?.msg || xhr.responseText || "网络错误";
                    reject(msg);
                }
            });
        });
    },

    get(url, params) {
        const query = params ? '?' + new URLSearchParams(params).toString() : '';
        return this.request('GET', url + query);
    },

    post(url, data) {
        return this.request('POST', url, data);
    },

    put(url, data) {
        return this.request('PUT', url, data);
    },

    delete(url, data) {
        return this.request('DELETE', url, data);
    },

    // 🔥 新增上传文件方法
    upload(url, file, extraData = {}) {
        return new Promise((resolve, reject) => {
            const formData = new FormData();
            formData.append('file', file);

            // 附加额外字段（比如 userId、类型等）
            for (const key in extraData) {
                formData.append(key, extraData[key]);
            }

            $.ajax({
                url: this.baseUrl + url,
                type: 'POST',
                data: formData,
                processData: false,     // 不让 jQuery 处理 FormData
                contentType: false,     // 让浏览器自动设置 multipart 边界
                headers: this._getAuthHeader(),
                success: res => {
                    if (res.code === 0) {
                        resolve(res.data);
                    } else {
                        reject(res.msg || "上传失败");
                    }
                },
                error: xhr => {
                    let msg = xhr.responseJSON?.msg || xhr.responseText || "网络错误";
                    reject(msg);
                }
            });
        });
    },

    // 带进度回调的版本（可选）
    uploadWithProgress(url, file, extraData = {}, onProgress) {
        return new Promise((resolve, reject) => {
            const formData = new FormData();
            formData.append('file', file);
            for (const key in extraData) formData.append(key, extraData[key]);

            $.ajax({
                url: this.baseUrl + url,
                type: 'POST',
                data: formData,
                processData: false,
                contentType: false,
                headers: this._getAuthHeader(),
                xhr: function () {
                    const xhr = $.ajaxSettings.xhr();
                    if (xhr.upload && onProgress) {
                        xhr.upload.onprogress = e => {
                            if (e.lengthComputable) {
                                const percent = (e.loaded / e.total * 100).toFixed(2);
                                onProgress(percent);
                            }
                        };
                    }
                    return xhr;
                },
                success: res => {
                    if (res.code === 0) resolve(res.data);
                    else reject(res.msg || "上传失败");
                },
                error: xhr => {
                    let msg = xhr.responseJSON?.msg || xhr.responseText || "网络错误";
                    reject(msg);
                }
            });
        });
    },

    _getAuthHeader() {
        const token = localStorage.getItem(this.tokenKey);
        return token ? { Authorization: `Bearer ${token}` } : {};
    },

    setToken(token) {
        localStorage.setItem(this.tokenKey, token);
    },

    clearToken() {
        localStorage.removeItem(this.tokenKey);
    }
};
