var tj = [];



function getrecommend() {
    let httpRequest = new XMLHttpRequest();
    httpRequest.open("GET", 'http://49.235.99.195:8090/home/recommend', true);
    httpRequest.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
    httpRequest.send("");
    httpRequest.onreadystatechange = () => {
        if (httpRequest.readyState == 4 && httpRequest.status == 200) {
            let res = httpRequest.responseText;
            let dx = res.replace(/}}/g, "}},");
            let item = JSON.parse("[" + dx.substring(0, dx.length - 1) + "]");
            tj = tj.concat(item);
            let oOl = document.getElementById('ol1');
            let outStr = "";
            for (let i = 0; i < 8; i++) {
                outStr += '<li><span><img src=' + tj[i].data.Img + 'style=" width=135px height=200px"></span> </br> <span class="name">' + tj[i].data.Name + "</span></li > ";
            }
            oOl.innerHTML = outStr;

            // 实现翻页功能
            let aBtn = document.getElementsByClassName('pageChange');
            for (let j = 0; j < aBtn.length; j++) { // 为每个页数按钮都设置点击事件
                aBtn[j].onclick = function (event) {
                    let pageNum = parseInt(event.target.getAttribute("pageNum")); // 获取自定义属性pageNum的值 并将其转换为数字类型  释：pageNum表明当前时第几页
                    outStr = "";
                    for (let i = pageNum * 8; i < pageNum * 8 + 8; i++) { //通过页数获取data中相应段的值
                        outStr += '<li><a href=' + tj[i].data.Url + ' target="_blank"><span><img src=' + tj[i].data.Img + ' width=135px height=200px></span></a> </br> <span class="name">' + tj[i].data.Name + "</span></li > ";
                    }
                    oOl.innerHTML = outStr;
                }
            }

            let pagez = document.getElementById('pagez')
            pagez.onclick = function () {
                pageNum += 1;
                outStr = "";
                for (let i = pageNum * 8; i < pageNum * 8 + 8; i++) { //通过页数获取data中相应段的值
                    outStr += '<li><a href=' + tj[i].data.Url + ' target="_blank"><span><img src=' + tj[i].data.Img + ' width=180px height=280px></span></a> </br> <span class="name">' + tj[i].data.Name + "</span></li > ";
                }
                oOl.innerHTML = outStr;
            }

            let pagej = document.getElementById('pagej')
            pagej.onclick = function () {
                pageNum -= 1;
                outStr = "";
                for (let i = pageNum * 6; i < pageNum * 6 + 6; i++) { //通过页数获取data中相应段的值
                    outStr += '<a href=' + tj[i].data.Url + ' target="_blank"><li><span><img src=' + tj[i].data.Img + ' width=180px height=280px></span></a> </br> <span class="name">' + tj[i].data.Name + "</span></li > ";
                }
                oOl.innerHTML = outStr;
            }
        }
    }
}

for (let i = 0; i <= 4; i++) {
    getrecommend()
    // console.log(data)
}


