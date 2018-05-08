// pages/details/details.js
var goods = require('data.js')
var app = getApp()
Page({

  /**
   * 页面的初始数据
   */
  data: {
    listData: [],

    head_hidden:true,
    headPacity:0,
    bodyTop:0,
    active:false,

    indicatorDots: false,
    autoplay: true,
    interval: 5000,
    duration: 1000,
    circular:true,

    toView:'test',

    modalNum:'',
    title:'',
    content:'',
    showModalStatus: false,

    buyNum:1,

    evaluteAll:6,
    evalute_1:2,
    evalute_2:1,
    evalute_3:3,
    
    server_1:"4.7",
    server_2:"4.7",
    server_3:"4.7",
    score_1:"低",
    score_2:"低",
    score_3:"低",
  },
  jumpTo: function (e) {
    // 获取标签元素上自定义的 data-opt 属性的值
    let target = e.currentTarget.dataset.opt;
    this.setData({
      toView: target
    })
  },

  //监听商品滚动
  scroll:function(e){
    console.log(e.detail.scrollTop);
    var scrollValue = e.detail.scrollTop;
    if (scrollValue <= 10){
      console.log(scrollValue);
      this.setData({
        head_hidden: true,
        bodyTop: 0
      })
    }else{
      console.log(scrollValue);
      this.setData({
        head_hidden: false,
        headPacity: 1,
        bodyTop: scrollValue > 200 ? 11 : (scrollValue / 20),
        headPacity: scrollValue > 200 ? 1 : (scrollValue / 200)
      })
    };
    while(true){
      if (scrollValue >= 0 && scrollValue<404){

      } else if (scrollValue >=404 && scrollValue < 580){

      } else if (scrollValue >= 580 && scrollValue < 1398){

      } else if (scrollValue >= 1398){

      }
      return false;
    }
  },

  addToCart:function(){
    console.log("加入购物车");
  },

  goPay:function(){
    console.log("跳转支付页面");
  },

  showModals0: function (e) {
    this.setData(
      {
        modalNum:0,
        title: this.data.listData[0].pro_sale[0].title,
        content: this.data.listData[0].pro_sale[1].content
      }
    );
    var currentStatu = e.currentTarget.dataset.statu;
    this.util(currentStatu)
  },
  showModals1: function (e) {
    this.setData(
      {
        modalNum:1,
        title: this.data.listData[0].pro_sale[0].title,
        content: this.data.listData[0].pro_sale[1].content
      }
    );
    var currentStatu = e.currentTarget.dataset.statu;
    this.util(currentStatu)
  },
  showModals2: function (e) {
    this.setData(
      {
        modalNum:2,
        title: this.data.listData[0].pro_serv[0].title,
        content: this.data.listData[0].pro_serv[1].content
      }
    );
    var currentStatu = e.currentTarget.dataset.statu;
    this.util(currentStatu)
  },
  showModals3: function (e) {
    this.setData(
      {
        modalNum:3,
        title: this.data.listData[0].pro_param[0].title,
        content: this.data.listData[0].pro_param[1].content
      }
    );
    var currentStatu = e.currentTarget.dataset.statu;
    this.util(currentStatu)
  },
  util: function (currentStatu) {
    /* 动画部分 */
    // 第1步：创建动画实例   
    var animation = wx.createAnimation({
      duration: 200,  //动画时长  
      timingFunction: "linear", //线性  
      delay: 0  //0则不延迟  
    });

    // 第2步：这个动画实例赋给当前的动画实例  
    this.animation = animation;

    // 第3步：执行第一组动画  
    animation.opacity(0).rotateX(-100).step();

    // 第4步：导出动画对象赋给数据对象储存  
    this.setData({
      animationData: animation.export()
    })

    // 第5步：设置定时器到指定时候后，执行第二组动画  
    setTimeout(function () {
      // 执行第二组动画  
      animation.opacity(1).rotateX(0).step();
      // 给数据对象储存的第一组动画，更替为执行完第二组动画的动画对象  
      this.setData({
        animationData: animation
      })

      //关闭  
      if (currentStatu == "close") {
        this.setData(
          {
            showModalStatus: false
          }
        );
      }
    }.bind(this), 200)

    // 显示  
    if (currentStatu == "open") {
      this.setData(
        {
          showModalStatus: true
        }
      );
    }
  },

  /**
   * 生命周期函数--监听页面加载
   */
  onLoad: function (options) {
    var that = this;
    var sysinfo = wx.getSystemInfoSync().windowHeight;
    wx.showLoading({
      title: '努力加载中',
    })
    //将本来的后台换成了easy-mock 的接口，所有数据一次请求完 略大。。
    that.setData({
      listData: goods.goods,
      loading: true
    })
    wx.hideLoading();
  },

  /**
   * 生命周期函数--监听页面初次渲染完成
   */
  onReady: function () {
  },

  /**
   * 生命周期函数--监听页面显示
   */
  onShow: function () {
  
  },

  /**
   * 生命周期函数--监听页面隐藏
   */
  onHide: function () {
  
  },

  /**
   * 生命周期函数--监听页面卸载
   */
  onUnload: function () {

  },

  /**
   * 页面相关事件处理函数--监听用户下拉动作
   */
  onPullDownRefresh: function () {
  
  },

  /**
   * 页面上拉触底事件的处理函数
   */
  onReachBottom: function () {
  
  },

  /**
   * 用户点击右上角分享
   */
  onShareAppMessage: function () {
  
  }
})
