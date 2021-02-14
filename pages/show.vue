<template lang="pug">
.show
  h2 UI
  .uilist
    a(v-for="(item, i) in uilist", :key="i", :href="item.url", target="_blank")
      img.NuxtImgx(:src="item.img")
      p {{ item.name }}
</template>

<style lang="sass" scoped>
.show
  h2
    margin: 0 1rem
    padding: 0 1rem
    border-left: 4px solid #108775
    transition: all .25s
    -moz-transition: all .25s
    -webkit-transition: all .25s
    -o-transition: all .25s
  h2:hover
    color: #108775
    border-left: 4px solid #00C58E
  .uilist
    list-style: none
    display: flex
    flex-wrap: wrap
    padding: 0
    a
      margin: 1rem
      color: #555
      text-decoration: none
    a:hover
      color: #108775
.NuxtImgx
  animation: 1s appearxsx
  margin: auto
@keyframes appearxsx
  0%
    opacity: 0
</style>

<script>
export default {
  data: () => ({
    uilist: [
      {
        name: "Education系统界面",
        img: "/show/ui/1.jpg",
        url: "https://www.zcool.com.cn/work/ZNDk5NDA3NzY=.html",
        author: "ouyangzetao",
      },
      {
        name: "音频剪接app安卓界面UI设计",
        img:
          "https://img.zcool.cn/community/010e2c60126ac911013f79282d39d8.jpg@260w_195h_1c_1e_1o_100sh.jpg",
        url: "https://www.zcool.com.cn/work/ZNDk1MTk1ODg=.html",
        author: "ouyangzetao",
      },
      {
        name: "区块链交易所",
        img:
          "https://img.zcool.cn/community/0183a25fc21fb711013fdcc7778b7e.jpg@260w_195h_1c_1e_1o_100sh.jpg",
        url: "https://www.zcool.com.cn/work/ZNDkyNzk0MTI=.html",
      },
      {
        name: "banner广告",
        img:
          "https://img.zcool.cn/community/01f6bf5fa0f1a711013fdcc7c7331a.jpg@260w_195h_1c_1e_1o_100sh.jpg",
        url: "https://www.zcool.com.cn/work/ZNDg3MTYyMDg=.html",
      },
      {
        name: "CARPOOL小程序界面设计",
        img:
          "https://img.zcool.cn/community/0183185f0025ada8012066213f4028.jpg@260w_195h_1c_1e_1o_100sh.jpg",
        url: "https://www.zcool.com.cn/work/ZNDU5OTE4ODA=.html",
      },
      {
        name: "第一次给自己设计宣传banner",
        img:
          "https://img.zcool.cn/community/0185885ec15531a8012148143745e6.jpg@260w_195h_1c_1e_1o_100sh.jpg",
        url: "https://www.zcool.com.cn/work/ZNDQ3MDMzNTY=.html",
      },
      {
        name: "感谢以前客户。这是如下评价内容..",
        img:
          "https://img.zcool.cn/community/01d7c95ebd0e13a801214814dae805.jpg@260w_195h_1c_1e_1o_100sh.jpg",
        url: "https://www.zcool.com.cn/work/ZNDQ2MDEyNDg=.html",
      },
      {
        name: "",
        img: "",
        url: "",
      },
      {
        name: "",
        img: "",
        url: "",
      },
      {
        name: "",
        img: "",
        url: "",
      },
      {
        name: "",
        img: "",
        url: "",
      },
      {
        name: "",
        img: "",
        url: "",
      },
      {
        name: "",
        img: "",
        url: "",
      },
      {
        name: "",
        img: "",
        url: "",
      },
      {
        name: "",
        img: "",
        url: "",
      },
      {
        name: "",
        img: "",
        url: "",
      },
      {
        name: "",
        img: "",
        url: "",
      },
      {
        name: "",
        img: "",
        url: "",
      },
      {
        name: "",
        img: "",
        url: "",
      },
      {
        name: "",
        img: "",
        url: "",
      },
      {
        name: "",
        img: "",
        url: "",
      },
      {
        name: "",
        img: "",
        url: "",
      },
      {
        name: "",
        img: "",
        url: "",
      },
      {
        name: "",
        img: "",
        url: "",
      },
      {
        name: "",
        img: "",
        url: "",
      },
      {
        name: "",
        img: "",
        url: "",
      },
      {
        name: "",
        img: "",
        url: "",
      },
      {
        name: "",
        img: "",
        url: "",
      },
      {
        name: "",
        img: "",
        url: "",
      },
      {
        name: "",
        img: "",
        url: "",
      },
      {
        name: "",
        img: "",
        url: "",
      },
      {
        name: "",
        img: "",
        url: "",
      },
      {
        name: "",
        img: "",
        url: "",
      },
      {
        name: "",
        img: "",
        url: "",
      },
      {
        name: "",
        img: "",
        url: "",
      },
      {
        name: "",
        img: "",
        url: "",
      },
      {
        name: "",
        img: "",
        url: "",
      },
      {
        name: "",
        img: "",
        url: "",
      },
      {
        name: "",
        img: "",
        url: "",
      },

      {
        name: "",
        img: "",
        url: "",
      },
      {
        name: "",
        img: "",
        url: "",
      },
      {
        name: "",
        img: "",
        url: "",
      },
      {
        name: "",
        img: "",
        url: "",
      },
      {
        name: "",
        img: "",
        url: "",
      },
      {
        name: "",
        img: "",
        url: "",
      },
      {
        name: "",
        img: "",
        url: "",
      },
      {
        name: "",
        img: "",
        url: "",
      },
      {
        name: "",
        img: "",
        url: "",
      },
      {
        name: "",
        img: "",
        url: "",
      },
      {
        name: "",
        img: "",
        url: "",
      },
      {
        name: "",
        img: "",
        url: "",
      },
      {
        name: "",
        img: "",
        url: "",
      },
      {
        name: "",
        img: "",
        url: "",
      },
      {
        name: "",
        img: "",
        url: "",
      },
    ],
  }),
};
</script>
