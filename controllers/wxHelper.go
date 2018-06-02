package controllers

import (
	"bytes"
	_ "bytes"
	"crypto/md5"
	"crypto/rand"
	_ "encoding/base64"
	"encoding/hex"
	_ "encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
	"net/http"
	_ "shop/models"
	"sort"
	"strconv"
	"strings"
	"time"
)

type WxHelperController struct {
	BaseController
}

func (this *WxHelperController) Get() {
	this.TplName = "product_edit.html"
	//this.Ctx.Output.Body([]byte(`你好，欢迎使用微信助手控制器`))
}

func (this *WxHelperController) Post() {

	act := this.Input().Get("act")
	//检查请求的方法
	if act != "" {
		switch act {
		//获取openid
		case "getOpenId":
			this.getOpenId()
			//小程序发起支付
		case "goPay":
			this.goPay()
			//APP端发起支付
		case "appGoPay":
			this.appGoPay()
			//获取支付签名
		case "getPaySign":
			this.getPaySign()
		default:
			this.Data["json"] = map[string]interface{}{"status": 400, "msg": "没有对应处理方法", "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()
			return

		}
	}
	// this.Data["json"] = map[string]interface{}{"status": 400, "msg": "没有对应处理方法", "time": time.Now().Format("2006-01-02 15:04:05")}
	// this.ServeJSON()

}

//获取openid--------------------------
func (this *WxHelperController) getOpenId() {

	appid := this.GetString("appid")
	secret := this.GetString("secret")
	js_code := this.GetString("js_code")
	grant_type := this.GetString("grant_type")
	// resp, err := http.PostForm("https://api.weixin.qq.com/sns/jscode2session", url.Values{"appid": {appid},
	// 	"secret":     {secret},
	// 	"js_code":    {js_code},
	// 	"grant_type": {grant_type}})
	resp, err := http.Get(fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%v&secret=%v&js_code=%v&grant_type=%v", appid, secret, js_code, grant_type))
	if err != nil {
		// handle error
		this.Data["json"] = map[string]interface{}{"status": 400, "msg": "获取openid出问题", "time": time.Now().Format("2006-01-02 15:04:05")}
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	// buf := new(bytes.Buffer)
	// buf.ReadFrom(body)
	beego.Info(fmt.Sprintf(":::::::::::::%v", string(body)))
	// var result []byte
	// var ba base64.Encoding
	// ba.Decode(result, []byte(body))
	this.Data["json"] = map[string]interface{}{"status": 200, "openid": string(body), "time": time.Now().Format("2006-01-02 15:04:05")}
	this.ServeJSON()
	return

}

//获取openid--------------------------
func (this *WxHelperController) getPaySign() {

	appId := this.GetString("appId")
	timeStamp := this.GetString("timeStamp")
	nonceStr := this.GetString("nonceStr")
	prepay_id := this.GetString("prepay_id")
	appSecret := this.GetString("mchKey")
	signStrings := fmt.Sprintf("appId=%v&nonceStr=%v&package=prepay_id=%v&signType=MD5&timeStamp=%v&key=%v", appId, nonceStr, prepay_id, timeStamp, appSecret)
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(signStrings))
	cipherStr := md5Ctx.Sum(nil)
	upperSign := strings.ToUpper(hex.EncodeToString(cipherStr))
	beego.Info(signStrings)

	this.Data["json"] = map[string]interface{}{"status": 200, "paySign": upperSign, "time": time.Now().Format("2006-01-02 15:04:05")}
	this.ServeJSON()
	return

}

// //发起支付---------------------------------------------------------
func (this *WxHelperController) goPay() {
	//获取下单参数
	appid := this.GetString("appid")
	mch_id := this.GetString("mch_id")
	spbill_create_ip := this.Ctx.Input.IP()
	//spbill_create_ip := "192.168.100.186"
	openid := this.GetString("openid")
	appSecret := this.GetString("mchKey")

	//-----------构造下单数据-----------------------------

	modOrder, ok := this.GetSession("confirmOrder").(confirmOrder)
	//out_trade_no := this.GetString("out_trade_no")
	if !ok {
		this.Data["json"] = map[string]interface{}{"status": 400, "msg": "确定订单获取失败,请重新下单", "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
		return
	}
	tmpOrder := modOrder.TmpOrder
	total_fee := tmpOrder.RealPrice * 100
	//total_fee, _ := strconv.ParseFloat("0.01", 64)
	out_trade_no := tmpOrder.OrderNum
	wxOrderOBJ := map[string]string{"total_fee": strconv.FormatFloat(total_fee, 'f', -1, 32), "spbill_create_ip": spbill_create_ip, "appid": appid, "mch_id": mch_id, "nonce_str": randStr(32, "alphanum"), "body": "小程序购物", "out_trade_no": out_trade_no, "notify_url": "https://mushangyun.com/wxhelper", "trade_type": "JSAPI", "openid": openid}
	wxOrderOBJ["attach"] = "0"
	wxOrderOBJ["detail"] = "中联国际"
	wxOrderOBJ["fee_type"] = "CNY"
	wxOrderOBJ["goods_tag"] = "WXG"
	wxOrderOBJ["product_id"] = "100"
	wxOrderOBJ["time_start"] = TimeConvert(1)
	wxOrderOBJ["time_expire"] = TimeConvert(2)

	//构造签名---------------------------------------------------------
	sorted_keys := make([]string, 0)
	for k, _ := range wxOrderOBJ {
		sorted_keys = append(sorted_keys, k)
	}

	sort.Strings(sorted_keys)

	//STEP2, 对key=value的键值对用&连接起来，略过空值
	var signStrings string
	for _, k := range sorted_keys {
		fmt.Printf("k=%v, v=%v\n", k, wxOrderOBJ[k])
		value := fmt.Sprintf("%v", wxOrderOBJ[k])
		if value != "" {
			signStrings = signStrings + k + "=" + value + "&"
		}
	}

	//STEP3, 在键值对的最后加上key=API_KEY

	signStrings = signStrings + "key=" + appSecret

	// var sign string
	// for key, item := range wxOrderOBJ {

	// 	sign += fmt.Sprintf("%v=%v", key, item) + `&`
	// }
	// sign = sign + "key=" + appSecret
	beego.Info(signStrings)
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(signStrings))
	cipherStr := md5Ctx.Sum(nil)
	upperSign := strings.ToUpper(hex.EncodeToString(cipherStr))
	beego.Info(upperSign)
	wxOrderOBJ["sign"] = upperSign
	//构造xml文档----------------------------------------------
	var _wxOrder wxOrder
	_wxOrder.Appid = wxOrderOBJ["appid"]
	_wxOrder.Mch_id = wxOrderOBJ["mch_id"]
	_wxOrder.Nonce_str = wxOrderOBJ["nonce_str"]
	_wxOrder.Sign = wxOrderOBJ["sign"]
	_wxOrder.Body = wxOrderOBJ["body"]
	_wxOrder.Out_trade_no = wxOrderOBJ["out_trade_no"]
	_wxOrder.Total_fee, _ = strconv.Atoi(wxOrderOBJ["total_fee"])
	_wxOrder.Spbill_create_ip = wxOrderOBJ["spbill_create_ip"]
	_wxOrder.Notify_url = wxOrderOBJ["notify_url"]
	_wxOrder.Trade_type = wxOrderOBJ["trade_type"]
	_wxOrder.Openid = wxOrderOBJ["openid"]
	_wxOrder.Attach = wxOrderOBJ["attach"]
	_wxOrder.Detail = wxOrderOBJ["detail"]
	_wxOrder.Fee_type = wxOrderOBJ["fee_type"]
	_wxOrder.Goods_tag = wxOrderOBJ["goods_tag"]
	_wxOrder.Product_id = wxOrderOBJ["product_id"]
	_wxOrder.Time_start = wxOrderOBJ["time_start"]
	_wxOrder.Time_expire = wxOrderOBJ["time_expire"]
	bytes_req, err := xml.Marshal(_wxOrder)
	if err != nil {
		beego.Info(err)
		this.Data["json"] = map[string]interface{}{"status": 400, "msg": "构造xml失败", "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
		return
	}

	str_req := string(bytes_req)
	str_req = strings.Replace(str_req, "wxOrder", "xml", -1)
	beego.Info(str_req)
	//发送请求---------------------------------------------------------------
	bytes_req = []byte(str_req)
	req, err := http.NewRequest("POST", "https://api.mch.weixin.qq.com/pay/unifiedorder", bytes.NewReader(bytes_req))
	if err != nil {
		beego.Info(err)
		this.Data["json"] = map[string]interface{}{"status": 400, "msg": "构造xml失败", "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
		return
	}
	req.Header.Set("Accept", "application/xml")
	//这里的http header的设置是必须设置的.
	req.Header.Set("Content-Type", "application/xml;charset=utf-8")

	c := http.Client{}
	resp, _err := c.Do(req)
	if _err != nil {
		beego.Info(err)
		this.Data["json"] = map[string]interface{}{"status": 400, "msg": "构造xml失败", "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
		return
	}
	//返回统一下单数据
	xmlResp := UnifyOrderResp{}
	body, _ := ioutil.ReadAll(resp.Body)
	beego.Info(string(body))
	_err = xml.Unmarshal(body, &xmlResp)
	if xmlResp.Return_code == "FAIL" {
		this.Data["json"] = map[string]interface{}{"status": 400, "msg": "微信后台下单成功", "data": xmlResp, "wxOrderOBJ": wxOrderOBJ, "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
		return
	}

	this.Data["json"] = map[string]interface{}{"status": 200, "msg": "微信后台下单成功", "data": xmlResp, "wxOrderOBJ": wxOrderOBJ, "time": time.Now().Format("2006-01-02 15:04:05")}
	this.ServeJSON()
	return

}

//APP发起支付----------------------------------------------
func (this *WxHelperController) appGoPay() {
	// appid := this.GetString("appid")
	// attach := this.GetString("attach")
	// body := this.GetString("body")
	// mch_id := this.GetString("mch_id")
	// nonce_str := this.GetString("nonce_str")
	// notify_url := this.GetString("notify_url")
	// out_trade_no := this.GetString("out_trade_no")
	// spbill_create_ip := this.GetString("spbill_create_ip")
	// total_fee := this.GetString("total_fee")
	// trade_type := this.GetString("trade_type")
	// sign := this.GetString("sign")
}

func randStr(strSize int, randType string) string {

	var dictionary string

	if randType == "alphanum" {
		dictionary = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	}

	if randType == "alpha" {
		dictionary = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	}

	if randType == "number" {
		dictionary = "0123456789"
	}

	var bytes = make([]byte, strSize)
	rand.Read(bytes)
	for k, v := range bytes {
		bytes[k] = dictionary[v%byte(len(dictionary))]
	}
	return string(bytes)
}

//下单结构体
type wxOrder struct {
	Attach           string `xml:"attach"`
	Appid            string `xml:"appid"`
	Body             string `xml:"body"`
	Detail           string `xml:"detail"`
	Fee_type         string `xml:"fee_type"`
	Goods_tag        string `xml:"goods_tag"`
	Mch_id           string `xml:"mch_id"`
	Nonce_str        string `xml:"nonce_str"`
	Notify_url       string `xml:"notify_url"`
	Product_id       string `xml:"product_id"`
	Time_start       string `xml:"time_start"`
	Time_expire      string `xml:"time_expire"`
	Trade_type       string `xml:"trade_type"`
	Spbill_create_ip string `xml:"spbill_create_ip"`
	Total_fee        int    `xml:"total_fee"`
	Out_trade_no     string `xml:"out_trade_no"`
	Sign             string `xml:"sign"`
	// Appid            string `xml:"appid"`
	// Mch_id           string `xml:"mch_id"`
	// Nonce_str        string `xml:"nonce_str"` //随机吗
	// Sign             string `xml:"sign"`
	// Body             string `xml:"body"`             //商品描述
	// Out_trade_no     string `xml:"out_trade_no"`     //商户订单号
	// Total_fee        string `xml:"total_fee"`        //订单金额，单位分
	// Spbill_create_ip string `xml:"spbill_create_ip"` //小程序ip地址
	// Notify_url       string `xml:"notify_url"`       //回调接受订单结果地址
	// Trade_type       string `xml:"trade_type"`       //值JSAPI
	Openid string `xml:"openid"`
}
type UnifyOrderResp struct {
	Return_code string `xml:"return_code"`
	Return_msg  string `xml:"return_msg"`
	Attach      string `xml:"attach"`
	Appid       string `xml:"appid"`
	Mch_id      string `xml:"mch_id"`
	Nonce_str   string `xml:"nonce_str"`
	Sign        string `xml:"sign"`
	Result_code string `xml:"result_code"`
	Prepay_id   string `xml:"prepay_id"`
	Trade_type  string `xml:"trade_type"`
	Code_url    string `xml:"code_url"`
}

func TimeConvert(span int) string {
	_now := time.Now()
	if span == 2 {
		_now = _now.Add(time.Hour * 2)
	}
	return _now.Format("20060102150405")
}

// //openid 机构体
// type openidObject struct {
// 	Session_key string
// 	Openid      string
// }
