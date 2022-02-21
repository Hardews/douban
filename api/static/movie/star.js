
window.onload = function () {
    let aSpan = document.getElementsByClassName("star")[0];
    let aStxt = document.getElementsByClassName("star-txt")[0];
    let aBstar = aSpan.getElementsByTagName("b");
    let arrBtxt = ["很差", "较差", "还行", "推荐", "力荐"];
    let num = 0;
    let onOff = true; for (let i = 0; i < aBstar.length; i++) {
        aBstar[i].index = i;
        //鼠标移入
        aBstar[i].onmouseover = function () {
            if (onOff) {
                num = this.index;
                aStxt.innerHTML = arrBtxt[num];
                for (let i = 0; i <= this.index; i++) {
                    aBstar[i].style.background = ' url(../模块/star.png)';
                }
            }
        };
        //鼠标移开
        aBstar[i].onmouseout = function () {
            if (onOff) {//设定个开关，等开关为真就执行鼠标移除的代码
                aStxt.innerHTML = "";
                for (let i = 0; i <= this.index; i++) {
                    aBstar[i].style.background = ' url(../模块/staroff.png)';
                }
            }
        };
        //鼠标点击
        aBstar[i].onclick = function () {
            onOff = false;//开关为假就不执行鼠标移除的代码
            //先清空
            aStxt.innerHTML = "";
            for (let i = 0; i < aBstar.length; i++) {
                aBstar[i].style.background = ' url(../模块/staroff.png)';
            }
            //点击当前星星，之前的都点亮包含自己
            num = this.index;
            aStxt.innerHTML = arrBtxt[num];
            for (let i = 0; i <= this.index; i++) {
                aBstar[i].style.background = ' url(../模块/star.png)';

            }

        };
    }
}
