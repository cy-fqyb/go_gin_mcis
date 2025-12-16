 $(function () 
 {
            $("#loginBtn").click(async function () {
                const user_name = $("#username").val().trim();
                const password = $("#password").val().trim();
                const $msg = $("#msg");

                if (!user_name || !password) {
                    $msg.text("请输入用户名和密码").removeClass("hidden");
                    return;
                }

                try {
                    const token = await api.post("/api/public/login", { user_name, password });
                    api.setToken(token);
                    $msg.text("登录成功！正在跳转...").removeClass("hidden").removeClass("text-red-500").addClass("text-green-600");

                    // 1秒后跳转首页
                    setTimeout(() => location.href = "/system/home.html", 800);
                } catch (err) {
                    console.error("登录失败:", err);
                    $msg.text(typeof err === "string" ? err : "登录失败，请检查用户名或密码").removeClass("hidden");
                }
            });
        }
    );