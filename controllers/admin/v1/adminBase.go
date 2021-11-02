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
	"github.com/garyburd/redigo/redis"
)

type adminBaseController struct {
	beego.Controller
	o       orm.Ormer
	user_id int64
}

type returnObj struct {
	code int16
	msg  string
}

//openssl genrsa -out rsa_private_key.pem 1024
var adminPrivateKey = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIICWwIBAAKBgQCuxk2OafXt888I0msU/k9xmyhdIHZuouFfDmtegGGdqoaM+sIt
A3C30gUvGL56yjysQYlf1kCdacXr2bF4f6z0u6rv84E59w8WeXfS0I3YeUjwGQ5t
mI8R+Rf8idre1DCkTpp5ffsA5z0SFzz1gsVyd5DoWl4oJY227H0HY3r5xQIDAQAB
AoGAAtHDaBiDEOZCUz124FQJDQwKpJQ7jgEzpjz68/+HLwdv1zkVM9BDUpc8quLC
FijcmUXqzgvn1RSHhXpECiSEjaXGLxo072FOFyslTQ/TbmKSN1EONOkLQKwy3r/f
ExxAvvs+saDWM4pbvOXDJmEbTFJyBLXzlWcedmyVRpBoovECQQD73OwogIFP8olM
Y1LrD3T8eWLZeGaOqT/pRrB3TZll9k3eLIHjH1dXPQDkT7U2J+i7weuWMwo7Umrn
N2Zcs1ANAkEAsaU50WSwbk0FiNfps2N1mJTFp4wkc1js/IE2GYwmwSDfdhhHPwW0
UtNdKdojpYsv0j7j3f1M0yyD/0GmXGYqmQJAFoJC9MevRtbVIGeMBIfoG5w5klfp
SnyjwpRXtwHPYMZnZSCzJvopExnXl4/sEP/2E7mb9VtwYabW+P0Bf+1ijQJAYdkC
ScXOMFMYU1GqFfcYlNyNKkZU5Xv7vPFm3ReHWSVEMIYa6Cm6M0zcqerPa6WIx6OA
W4vjvwVsBzMf8RENMQJAIYP30rtjf2boAFrkqrjc5SnfW1YrHFVUdX4iFq5FjSyG
nvLJgPlOW7F++ahqEm8VCn6Tv0vv/BpiX88u5XbBvA==
-----END RSA PRIVATE KEY-----
`)

//openssl
//openssl rsa -in rsa_private_key.pem -pubout -out rsa_public_key.pem
var adminPublicKey = []byte(`-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCuxk2OafXt888I0msU/k9xmyhd
IHZuouFfDmtegGGdqoaM+sItA3C30gUvGL56yjysQYlf1kCdacXr2bF4f6z0u6rv
84E59w8WeXfS0I3YeUjwGQ5tmI8R+Rf8idre1DCkTpp5ffsA5z0SFzz1gsVyd5Do
Wl4oJY227H0HY3r5xQIDAQAB
-----END PUBLIC KEY-----`)

//初始化方法
func (p *adminBaseController) Prepare() {

	//util.Logger.Error("Error")

}

/**
 * @description: 公共用户身份验证接口
 * @param {*baseController} yxm --- 2020-10-29
 * @return {*} *returnObj
 */
func (p *adminBaseController) verifAdminIdentity() *returnObj {

	signBytes, err := base64.StdEncoding.DecodeString(p.Ctx.Input.Header("sign"))

	if err != nil {
		return &returnObj{10001, "身份验证失败！错误：" + err.Error()}
	}

	origData, _ := p.AdminRsaDecrypt(signBytes)

	sign_map := make(map[string]string)
	err = json.Unmarshal(origData, &sign_map)
	if err != nil {
		return &returnObj{10001, "身份验证失败！错误：" + err.Error()}
	}

	redis_pool, err := util.GetRedisConnection(2)
	if err != nil {
		return &returnObj{10001, "身份验证失败！错误：" + err.Error()}
	}
	defer redis_pool.Close() //函数运行结束 ，把连接放回连接池
	redis_key := "login:admin:uuid:" + sign_map["admin_uuid"]

	r, error := redis.String(redis_pool.Do("Get", redis_key))
	if error != nil {
		return &returnObj{10001, "身份验证失败！错误：" + err.Error()}
	}

	userinfo := make(map[string]interface{})

	err = json.Unmarshal([]byte(r), &userinfo)
	if err != nil {
		return &returnObj{10001, "身份验证失败！错误：" + err.Error()}
	}

	userid, ok := userinfo["admin_id"].(float64)
	if ok {
		p.user_id = int64(userid)
	}
	if userinfo["admin_token"].(string) == sign_map["admin_token"] {
		return &returnObj{200, "成功"}
	} else {
		return &returnObj{10001, "身份验证失败！错误：" + err.Error()}
	}

}

/**
 * @description: 获取用户IP地址
 * @param {*baseController} yxm --- 2020-10-29
 * @return {*} string
 */
func (p *adminBaseController) getClientIp() string {
	s := strings.Split(p.Ctx.Request.RemoteAddr, ":")
	return s[0]
}

/**
 * @description: Rsa加密
 * @param {*baseController} yxm --- 2020-10-29
 * @return {*} []byte, err
 */
func (p *adminBaseController) AdminRsaEncrypt(origData []byte) ([]byte, error) {
	//解密pem格式的公钥
	block, _ := pem.Decode(adminPublicKey)
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
func (p *adminBaseController) AdminRsaDecrypt(ciphertext []byte) ([]byte, error) {
	//解密
	block, _ := pem.Decode(adminPrivateKey)
	if block == nil {
		return nil, errors.New("private key error!")
	}
	//p.Ctx.Output.JSON(err, false, true)
	//解析PKCS1格式的私钥
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 解密
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}
