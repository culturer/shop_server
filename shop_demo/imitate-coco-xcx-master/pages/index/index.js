//index.js
//获取应用实例
//const app = getApp()
const APP_ID = 'wx7db7b53153ae27dd';//输入小程序appid  
const APP_SECRET = '22cacdf349b14646350ecf6026d79129';//输入小程序app_secret  
var OPEN_ID = ''//储存获取到openid  
var SESSION_KEY = ''//储存获取到session_key  
Page({
  data: {
    //轮播图
    imgUrls: [
      '../../images/seafood_1.jpg',
      '../../images/seafood_2.jpg',
      '../../images/seafood_3.jpg',
      '../../images/seafood_4.jpg',
      '../../images/seafood_5.jpg'
    ],
    indicatorDots: true,
    autoplay: true,
    interval: 2000,
    duration: 500,
    circular:true
  },
  onLoad: function () {
    this.getOpenId(this.loginUser)
  },
  golist: function () {
    wx.navigateTo({
      url: '../list/list'
    })
  },
  //获取openId
  getOpenId: function (loginUser) {
    var that = this;
    wx.login({
      success: function (res) {
        wx.request({
          //获取openid接口  
          url: 'https://api.weixin.qq.com/sns/jscode2session',
          data: {
            appid: APP_ID,
            secret: APP_SECRET,
            js_code: res.code,
            grant_type: 'authorization_code'
          },
          method: 'GET',
          success: function (res) {
            console.log(res.data)
            OPEN_ID = res.data.openid;//获取到的openid 
            wx.setStorage({
              key: 'vId',
              data: OPEN_ID,
            }) 
            SESSION_KEY = res.data.session_key;//获取到session_key  
       
            loginUser(OPEN_ID)

          }
        })
      }
    })
  },
  //登录用户
  loginUser: function (OPEN_ID) {
    wx.request({
      //获取openid接口  
      url: 'https://mushangyun.com/login',
      data: {
        options: 2,
        vId: OPEN_ID,
       
      },
      method: 'POST',
      success: function (res) {
      
        console.log(res.data)
        //var date = JSON.parse(res.data)
        if (res.data.status=="400"){
          wx.navigateTo({
            url: '../register/register',
          })
        } else if (res.data.status == "200"){
          wx.showToast({
            title: res.data.msg,
          })
        }

      },
      fail:function(res){
        console.log(res)
        // wx.navigateTo({
        //   url: '../register/register',
        // })
      }
    })
  },
  tapLogin:function(){
    this.loginUser(OPEN_ID)
  }
})
