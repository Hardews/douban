function getname() {
    let httpRequest = new XMLHttpRequest();
    httpRequest.open("GET", 'http://49.235.99.195:8090/user/TEST/menu', true);
    httpRequest.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
    httpRequest.send("");
    httpRequest.onreadystatechange = () => {
        if (httpRequest.readyState == 4 && httpRequest.status == 200) {
            let data = JSON.parse(httpRequest.responseText);
            {
                if (data.nickName != null) {
                    let nickName = data.nickName;
                    document.getElementById('nickName').innerHTML = nickName;
                }
                else {
                    document.getElementById('nickName').innerHTML = "无名者";
                }

            }
        }
    }

}

function getin() {
    let xhr = new XMLHttpRequest();
    xhr.open("GET", 'http://49.235.99.195:8090/user/TEST/menu', true);
    xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
    xhr.send("");
    xhr.onreadystatechange = () => {
        if (xhr.readyState == 4 && xhr.status == 200) {
            let data = JSON.parse(xhr.responseText);
            {

                if (data.introduce != undefined) {
                    let introduce = data.introduce;
                    document.getElementById('introduce').innerHTML = introduce;
                }
                else {
                    document.getElementById('introduce').innerHTML = "快来更新你的简介吧";
                }
            }
        }
    }

}

getname();
getin();

// 修改简介

let fix = document.getElementById('fix')
let fixb = document.getElementById('fixbox')
fix.onclick = function () {
    fixb.style.display = "inline"
    fix.style.display = "none"
}

let fixa = document.getElementById('fixa')
fixa.onclick = function () {
    fixb.style.display = "none";
    fix.style.display = "inline"
}

let can = document.getElementById('cancel')
can.onclick = function () {
    fixb.style.display = "none"
    fix.style.display = "inline"
}