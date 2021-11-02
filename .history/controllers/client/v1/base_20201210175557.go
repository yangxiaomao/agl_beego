package controllers

import (
	"beeapi/util"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/beego/i18n"
	"github.com/garyburd/redigo/redis"
)

var langTypes []string // Languages that are supported.

func init() {
	// Initialize language type list.
	langTypes = strings.Split(beego.AppConfig.String("lang_types"), "|")

	// Load locale files according to language types.
	for _, lang := range langTypes {
		beego.Trace("Loading language: " + lang)
		if err := i18n.SetMessage(lang, "conf/"+"locale_"+lang+".ini"); err != nil {
			beego.Error("Fail to set message file:", err)
			return
		}
	}
}

type baseController struct {
	beego.Controller
	i18n.Locale
	o       orm.Ormer
	user_id int64
}

type returnObj struct {
	code int16
	msg  string
}

//openssl genrsa -out rsa_private_key.pem 1024
var privateKey = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIICXAIBAAKBgQCgZf00dse3Ww3xlz+vIsgw6zS28ltf7B7wqQYf9tHD8LO2UhPq
RLePMllheXQExm4Q68aaWasNipi2bRYzNVwx5GW795/hS/4ZFprfarz6XOdihb8C
11fBHgeJ55tRwFebWuNIYtiHmwsB5guKlayLYh35mTMjqLLmIPRLsJa3BwIDAQAB
AoGAM4C8HoH7W+0xW38w1DgLZuXHTe0hGPpU7vqe/FmA/nUGB4dwXJtHA4RrvchF
UBk1E1rZsQsUySrVIVKCu9uo59q1ZKypCvG5nOI/siC9eyj9R837oPk11jeFEd1V
OgW+2cNcCH0GaOg9wBTRdc8kJSBN320SSH2QgqVvLSF+EAkCQQDrNKYRQJva5wQ9
d9jmYm7DCZMXwoqZZta2NxTrZnRC4x8Td7c0WyEK3qZfhLtpvcrdAR4KZcM5Nwj9
HSBypOGJAkEArpQ/AwTL/IixDuA+5ET9WlJkgwuMkOX9+y/h/9lRMcCGQIAooF05
bA16Zm8gxNtWzejlnoLbiBzIlmKD6aGADwJBAKcLmFouacKZSgCh6pENHZ81YJuC
Vk9Go318N0ZMWIvcpIh+AMaVZi1DHiQ+r6AU5Ev77Cr0RSeQd7jUg+QISAkCQDHU
Wo/wBJFWKsOGOi1Ji88GCW5mE38gRng12JoAW26J90fdzLrJISS4TCGEwqOtv38S
ZMfzrdmR7acPA3vh/v0CQAQnT6bon5mci4GH4Cl40MQb882qLe3iZaUYJUu4A1eu
icckhmBWHfmSMaSSGJc8Y0tLdnTzDaT1ufzqcw45Cs8=
-----END RSA PRIVATE KEY-----
`)

//openssl
//openssl rsa -in rsa_private_key.pem -pubout -out rsa_public_key.pem
var publicKey = []byte(`-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCgZf00dse3Ww3xlz+vIsgw6zS2
8ltf7B7wqQYf9tHD8LO2UhPqRLePMllheXQExm4Q68aaWasNipi2bRYzNVwx5GW7
95/hS/4ZFprfarz6XOdihb8C11fBHgeJ55tRwFebWuNIYtiHmwsB5guKlayLYh35
mTMjqLLmIPRLsJa3BwIDAQAB
-----END PUBLIC KEY-----`)

// Prepare implemented Prepare() method for baseController.
// It's used for language option check and setting.
func (p *baseController) Prepare() {
	// Reset language option.
	p.Lang = "" // This field is from i18n.Locale.

	// 1. Get language information from 'Accept-Language'.
	al := p.Ctx.Request.Header.Get("Accept-Language")
	if len(al) > 4 {
		al = al[:5] // Only compare first 5 letters.
		if i18n.IsExist(al) {
			p.Lang = al
		}
	}

	// 2. Default language is English.
	if len(p.Lang) == 0 {
		p.Lang = "en-US"
	}

	// Set template level language option.
	p.Data["Lang"] = p.Lang
}

/**
 * @description: 公共用户身份验证接口
 * @param {*baseController} yxm --- 2020-10-29
 * @return {*} *returnObj
 */
func (p *baseController) verifUserIdentity() *returnObj {

	signBytes, err := base64.StdEncoding.DecodeString(p.Ctx.Input.Header("sign"))

	if err != nil {
		return &returnObj{10001, "身份验证失败！"}
	}

	origData, _ := p.RsaDecrypt(signBytes)

	sign_map := make(map[string]string)
	err = json.Unmarshal(origData, &sign_map)
	if err != nil {
		return &returnObj{10001, "身份验证失败！"}
	}

	redis_pool, err := util.GetRedisConnection(2)
	if err != nil {
		return &returnObj{10001, "身份验证失败！"}
	}
	defer redis_pool.Close() //函数运行结束 ，把连接放回连接池
	redis_key := "login:token:uuid:" + sign_map["user_uuid"]

	r, error := redis.String(redis_pool.Do("Get", redis_key))
	if error != nil {
		return &returnObj{10001, "身份验证失败！"}
	}

	userinfo := make(map[string]interface{})

	err = json.Unmarshal([]byte(r), &userinfo)
	if err != nil {
		return &returnObj{10001, "身份验证失败！"}
	}

	userid, ok := userinfo["user_id"].(float64)
	if ok {
		p.user_id = int64(userid)
	}

	if userinfo["user_token"].(string) == sign_map["user_token"] {
		return &returnObj{200, "成功"}
	} else {
		return &returnObj{10001, "身份验证失败！"}
	}

}

/**
 * @description: 获取用户IP地址
 * @param {*baseController} yxm --- 2020-10-29
 * @return {*} string
 */
func (p *baseController) getClientIp() string {
	s := strings.Split(p.Ctx.Request.RemoteAddr, ":")
	return s[0]
}

/**
 * @description: Rsa加密
 * @param {*baseController} yxm --- 2020-10-29
 * @return {*} []byte, err
 */
func (p *baseController) RsaEncrypt(origData []byte) ([]byte, error) {
	//解密pem格式的公钥
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}
	// 解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 类型断言
	pub := pubInterface.(*rsa.PublicKey)
	//加密
	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
	//return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}

/**
 * @description: Rsa解密
 * @param {*baseController} yxm --- 2020-10-29
 * @return {*} []byte, err
 */
func (p *baseController) RsaDecrypt(ciphertext []byte) ([]byte, error) {
	//解密
	block, _ := pem.Decode(privateKey)

	if block == nil {
		return nil, errors.New("private key error!")
	}
	//p.Ctx.Output.JSON(err, false, true)
	//解析PKCS1格式的私钥
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	//p.Ctx.Output.JSON(err, false, true)
	if err != nil {
		return nil, err
	}
	// 解密
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}

// AppController handles the welcome screen that allows user to pick a technology and username.
type AppController struct {
	baseController // Embed to use methods that are implemented in baseController.
}

// Get implemented Get() method for AppController.
func (this *AppController) Get() {
	this.TplName = "welcome.html"
}

// Join method handles POST requests for AppController.
func (this *AppController) Join() {
	// Get form value.
	uname := this.GetString("uname")
	tech := this.GetString("tech")

	// Check valid.
	if len(uname) == 0 {
		this.Redirect("/", 302)
		return
	}

	switch tech {
	case "longpolling":
		this.Redirect("/lp?uname="+uname, 302)
	case "websocket":
		this.Redirect("/ws?uname="+uname, 302)
	default:
		this.Redirect("/", 302)
	}

	// Usually put return after redirect.
	return
}
