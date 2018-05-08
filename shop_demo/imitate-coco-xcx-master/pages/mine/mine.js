// pages/mine/mine.js
const APP_ID = 'wx7db7b53153ae27dd';//输入小程序appid  
const APP_SECRET = '22cacdf349b14646350ecf6026d79129';//输入小程序app_secret  
var OPEN_ID = ''//储存获取到openid  
var SESSION_KEY = ''//储存获取到session_key  
Page({

  /**
   * 页面的初始数据
   */
  data: {
    nickName:"",
    avatarUrl:""
  },

  /**
   * 生命周期函数--监听页面加载
   */
  onLoad: function (options) {
    var that=this;
    wx.getUserInfo({
      success: function (res) {
        var userInfo = res.userInfo
        that.setData({
          nickName: userInfo.nickName,
          avatarUrl: userInfo.avatarUrl,
        })
      }
    })
  },
  bitphone:function(){
    wx.makePhoneCall({
      phoneNumber: '1340000' 
    })
  },
  //获取openId
  getOpenId: function () {
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
    
          }
        })
      }
    })
  },
  //登录用户
  loginUser:function(){

  }
 

 
})