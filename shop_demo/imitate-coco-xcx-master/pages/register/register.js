var app = getApp()
Page({
  data: {
    vId: "",
    name:"",
    tel:"",
    pwd:""
  },
  onReady: function (options) {
    // 页面初始化 options为页面跳转所带来的参数
    this.initUserInfo()
    
  },
  initUserInfo:function(){
    var that=this
    //获取用户vId
    wx.getStorage({
      key: 'vId',
      success: function(res) {
         that.setData({vId:res.data})
      },
    })
    //获取用户nicName
    wx.getUserInfo({
      success: function (res) {
        var userInfo = res.userInfo
        that.setData({
          name: userInfo.nickName          
        })
       console.log(that.data)
      }
    })

  },
  registerUser:function(e){
    var that=this
    that.setData()
    wx.request({
      //获取openid接口  
      url: 'https://mushangyun.com/register',
      data: {
        vId: that.data.vId,
        name: that.data.name,
        tel: that.data.tel,
        pwd: that.data.pwd,
      },
      method: 'POST',
      success: function (res) {
        console.log(res.data)
        if(res.data.status=="200"){
          wx.navigateTo({
            url: '../index/index',
          })
        } else if (res.data.status == "400"){
           wx.showToast({
             title:res.data.msg,
           })
        }

      }
    })
  },
  watchName:function(e){
   var that=this
   that.setData({ name: e.detail.value})
  },
  watchTel: function (e) {
    var that = this
    that.setData({ tel: e.detail.value })
  },
  watchPwd: function (e) {
    var that = this
    that.setData({ pwd: e.detail.value })
  },


})