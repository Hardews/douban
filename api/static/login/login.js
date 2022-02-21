

// 登录
let btn = document.getElementById("login");
btn.onclick = function () {
    let logName = document.getElementById("loginname");
    let password = document.getElementById("password");
    if (logName.value && password.value && password.value.length > 0 && logName.value.length > 0) {
        getuser(logName.value, password.value);
    }
    else {
        if (!logName.value) {

            logName.focus();
        }

        if (!password.value) {

            password.focus();
        }
    }

    function getuser(logName, password) {
        let xhr = new XMLHttpRequest();

        xhr.open("POST", 'http://49.235.99.195:8090/login', true);
        xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
        xhr.send("logName=" + logName + "&password=" + password);
        xhr.onreadystatechange = () => {
            if (xhr.readyState == 4 && xhr.status == 200) {

                let data = xhr.responseText;
                data = data.replace("}{", ",");
                data = JSON.parse(data)
                if (data.info == '成功') {
                    // 这里还得加上登录跳转

                    alert("登录成功");
                    // window.location.href = "space\space.html"
                    // console.log(data)
                }
                else {
                    // console.log(data)
                    alert("登录失败");
                }

            }
        }

    }
}

//注册
let sign = document.getElementById("sign");
sign.onclick = function () {
    let signName = document.getElementById("signname");
    let signPassword = document.getElementById("signpassword");
    let nickName = document.getElementById("nickname");
    if (signName.value && signPassword.value && signPassword.value.length > 0 && signName.value.length > 0) {
        postuser(signName.value, signPassword.value, nickName.value);
    }
    else {
        if (!signName.value) {

            signName.focus();
        }

        if (!signPassword.value) {

            signPassword.focus();
        }

        if (!nickName.value) {

            nickName.focus();
        }
    }

    function postuser(signName, signPassword, nickName) {


        let httpRequest = new XMLHttpRequest();
        httpRequest.withCredentials = true;
        // let obj2 = [{
        //     "key": "signName",
        //     "value": signName,
        //     "enabled": true,
        //     "type": "text",
        //     "description": ""

        // },
        // {
        //     "key": "signPassword",
        //     "value": signPassword,
        //     "enabled": true,
        //     "type": "text",
        //     "description": ""
        // },
        // {
        //     "key": "nickname",
        //     "value": nickname,
        //     "enabled": true,
        //     "type": "text",
        //     "description": ""
        // }]
        httpRequest.open("POST", 'http://49.235.99.195:8090/register', true);
        httpRequest.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
        httpRequest.send("signName=" + signName + "&signPassword=" + signPassword + "&nickName=" + nickName)
        httpRequest.onreadystatechange = () => {
            if (httpRequest.readyState == 4 && httpRequest.status == 200) {
                let data = JSON.parse(httpRequest.responseText);
                if (data.info == '成功') {
                    // 注册跳转
                    alert("注册成功");
                }
                else {
                    // console.log(data);
                    alert("注册失败");
                }

            }
        }

    }
}



//tab翻页功能

let tab_list = document.querySelector('.change');
let lis = tab_list.querySelectorAll('li');
let items = document.querySelectorAll('.item');

lis[0].className = 'current';
// 预处理

for (var i = 0; i < lis.length; i++) {

    lis[i].setAttribute('index', i);
    lis[i].onclick = function () {
        //清除class这个类
        for (var i = 0; i < lis.length; i++) {
            lis[i].className = '';
        }
        //留下我自己
        this.className = 'current';
        //下面的显示内容模块
        var index = this.getAttribute('index');
        // console.log(index);
        //先让所有的item这些div隐藏
        for (var i = 0; i < items.length; i++) {
            items[i].style.display = 'none';
        }
        //留下我自己的，让对应选中的item，显示出来就可以了 
        items[index].style.display = 'block';

    }
}
