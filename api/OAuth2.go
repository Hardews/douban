package api

import (
	"douban/middleware"
	"douban/model"
	"douban/service"
	"douban/tool"
	"encoding/json"
	"net/url"

	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	var html = `<html><body><a href="/tLogin">GitHub login</a></body></html>`
	_, err := fmt.Fprint(c.Writer, html)
	if err != nil {
		fmt.Println("set html failed,err:", err)
		tool.RespInternetError(c)
		return
	}
}

type URL struct {
	AuthURL  string
	TokenURL string
}

type config struct {
	ClientID     string
	ClientSecret string
	Endpoint     URL
	RedirectURL  string
	Scopes       []string
}

var (
	githubOauthConfig = &config{
		ClientID:     os.Getenv("GITHUB_OAUTH2_CLIENT_ID"),
		ClientSecret: os.Getenv("GITHUB_OAUTH2_CLIENT_SECRET"),
		Endpoint: URL{
			AuthURL:  "https://github.com/login/oauth/authorize",
			TokenURL: "https://github.com/login/oauth/access_token",
		},
		RedirectURL: "http://49.235.99.195:8090/callback",
	}
	state = "Hardews" // 应该是随机字符串
)

func LoginByThirdParty(c *gin.Context) {
	//重定向到授权网页
	OAuth2Url := "https://github.com/login/oauth/authorize?" +
		"client_id=" + githubOauthConfig.ClientID +
		"&redirect_uri=" + githubOauthConfig.RedirectURL +
		"&response_type=code" +
		"&state=" + state
	c.Redirect(http.StatusMovedPermanently, OAuth2Url)
}

func CallBack(c *gin.Context) {
	if c.Query("state") != state {
		tool.RespErrorWithDate(c, "state is not valid")
		return
	}
	// 获取code
	code := c.Query("code")

	// 用code换token

	// 形成请求
	var err error
	var tokenStr = "https://github.com/login/oauth/access_token" // github的access token 获取接口

	// 形成表单
	postData := url.Values{}
	postData.Add("grant_type", "authorization_code")
	postData.Add("code", code)
	postData.Add("client_id", githubOauthConfig.ClientID)
	postData.Add("client_secret", githubOauthConfig.ClientSecret)

	body := strings.NewReader(postData.Encode())

	//发送请求
	var req *http.Request
	req, err = http.NewRequest(http.MethodPost, tokenStr, body)
	if err != nil {
		fmt.Println("get token failed,err:", err)
		tool.RespInternetError(c)
		return
	}
	// 设置表单类型和返回格式
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("accept", "application/json")
	var resp *http.Response
	var client = http.Client{}
	// 获取响应
	if resp, err = client.Do(req); err != nil {
		fmt.Println("get resp failed,err:", err)
		tool.RespInternetError(c)
		return
	}

	defer resp.Body.Close()

	// 将响应的数据写入 tokenInfo 中，并返回
	var tokenInfo = make(map[string]interface{})
	if err = json.NewDecoder(resp.Body).Decode(&tokenInfo); err != nil {
		fmt.Println("write token in failed,err:", err)
		tool.RespInternetError(c)
		return
	}

	// 用access token获取用户信息
	// 形成请求 发送请求并获取响应
	var requ *http.Request
	var userInfoUrl = "https://api.github.com/user" // github用户信息获取接口
	if requ, err = http.NewRequest(http.MethodGet, userInfoUrl, nil); err != nil {
		fmt.Println("get userinfo failed,err,", err)
		tool.RespInternetError(c)
		return
	}
	// 设置请求头和返回信息格式
	requ.Header.Set("accept", "application/json")
	requ.Header.Set("Authorization", fmt.Sprintf("token %s", tokenInfo["access_token"].(string)))

	// 将响应的数据写入 userInfo 中，并返回
	var res *http.Response
	var tClient = http.Client{}
	if res, err = tClient.Do(requ); err != nil {
		fmt.Println("do get info request failed,err,", err)
		tool.RespErrorWithDate(c, "连接超时")
		return
	}

	defer res.Body.Close()

	var userInfo = make(map[string]interface{})
	if err = json.NewDecoder(res.Body).Decode(&userInfo); err != nil {
		fmt.Println("get userinfo failed err:", err)
		tool.RespInternetError(c)
		return
	}
	var user model.User
	for s, i := range userInfo {
		if s == "id" {
			id := strconv.FormatFloat(i.(float64), 'f', 0, 64)
			user.Username = id
		}
		if s == "login" {
			nickname := userInfo["login"].(string)
			user.Nickname = nickname
		}
	}

	// 验证是否注册
	err, flag := service.CheckUsername(user)
	if err != nil {
		tool.RespInternetError(c)
		fmt.Println("check username failed, error: ", err)
		return
	}

	// 已注册直接登录
	if flag == false {
		tool.RespSuccessfulWithDate(c, "第三方登录的账号已存在!为您登录")
		var identity = "用户"
		flag = service.CheckAdministratorUsername(user.Username)
		if flag {
			identity = "管理员"
		}
		token, flag := middleware.SetToken(user.Username, identity)
		if !flag {
			tool.RespInternetError(c)
			return
		}
		c.JSON(200, gin.H{
			"msg": token,
		})
		tool.RespSuccessful(c)
		return
	}

	// 未注册就注册新账号
	tool.RespSuccessfulWithDate(c, "检测到新账号，自动为您注册")
	// 令密码等于账号
	user.Password = user.Username

	//加密
	err, user.Password = service.Encryption(user.Password)
	if err != nil {
		tool.RespInternetError(c)
		fmt.Println("encryption failed , err :", err)
		return
	}

	err = service.WriteIn(user)
	if err != nil {
		tool.RespInternetError(c)
		fmt.Println("register failed,err:", err)
		return
	}

	var identity = "用户"
	flag = service.CheckAdministratorUsername(user.Username)
	if flag {
		identity = "管理员"
	}
	tokenMsg, flag := middleware.SetToken(user.Username, identity)
	if !flag {
		tool.RespInternetError(c)
		return
	}

	c.JSON(200, gin.H{
		"msg":      "初始密码与id一致，已登录!",
		"token":    tokenMsg,
		"id":       user.Username,
		"nickname": user.Nickname,
	})
	//c.Redirect(301, "http://49.235.99.195")
}
