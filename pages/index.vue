<style lang="sass" scoped>
ol.start
  list-style-type: none
  counter-reset: sectioncounter
  width: 200px
  li
    font-size: 1.5rem
  li:before
    content: counter(sectioncounter) ""
    counter-increment: sectioncounter
    color: #ffffff
    background: #991144
    border-radius: 50%
    width: 1em
    height: 1em
    display: inline-block
    text-align: center
    line-height: 1em
ul
  list-style: none
  margin: 0
  padding: 0
  display: flex
  li
    height: 12rem
    width: 16rem
    padding: 1rem
    margin: 1rem
    background: rgba(0,0,0,0.05)
</style>

<template lang="pug">
.container_index
  .header
    h1 86 位社区成员在远程协作
  .info
    h2 远程工作 / 自由职业
    ol
      li 通过任务拆组协作来完成项目
      li 创造自由和按照自己的时间表和兴趣/心情/状态调节工作时间
    button 加入我们
    p 发布任务/ 发布项目/ 发布idea/ 创建组织/ 发布服务/
    p 如何提升自身价值而不是被调用?
    p 项目类型/ 沉淀/研发
  //vs-avatar(
  //  primary,
  //  v-for="user in users",
  //  :key="user.id",
  //  @click="open_user(user.id)"
  //)
  //  template(v-if="user.avatar === ''", #text) {{ user.name }}
  //  img(v-else, src="user.avatar")
  div 社区成员总数 / 活跃人数 / 小组
  div 项目组 开发组 产品组 数据组 文案组 设计组 交互组
  div 专长 web应用 / 桌面应用 / APP
  h3 接受任务委托(雇佣)/ >> 更多
  ul
    li task item
    li task item
    li task item
    li task item
  h3 参与项目(合作)/ >> 更多
  ul
    li porject item
    li porject item
    li porject item
    li porject item
  p 自身是否符合要求
  p 项目是否具备预期价值
  h3 加入组织(他们正在招募)/ >> 更多
  ul
    li org item
    li org item
    li org item
    li org item

  h2 我们是怎样的组织?
  ol
    li 松散的, 自由的
  h2 我们如何运作?
  ol
    li 通过兴趣小组和任务调度
    li 各取所长并没有职位限制的
  h2 我们期望实现什么?
  ol
    li 想要一个假期? 立刻 >> 查看如何实现调度
    li 远程
    li 自由
    li 协作
    li 组队
  h3 发挥你的特长
  ol
    li 信息孤岛
    li 避免重复再生产
  div 资金回流与分配
    p 通过参与项目的小组作综合评估, 资金分配到各组, 然后在组内评估资金分配到个人
    p 或直接分配到个人
  ol.start
    li 参与兴趣项目
    li 运用时间价值
    li 自由调度任务(分流你的工作)
    li 多人协作, 取长补短(加入组队)
    li 构建管道收入
  div 如何发挥价值(每个人的)? 如何产生价值(生产总值)?
  div 解决哪些问题?
    p 1. 远程起步门槛
    p 2. 资历精力占比
    p 3. 产能效率 >> 如何提升
    p 4. 时间自由调度 >> 如何实现
    p 5. 各取所长 互通有无
  div 我们正在为此怎么做
  div 你是否需要我们?
    p 订阅
    p 网站建设
    p APP开发
    P 小程序
    p 桌面应用程序
    p 微信联系 / 致电
  div 加入我们吗?
    p 电子邮件
    p 微信群
    p qq
  //vs-avatar(success)
  //  template(#text) Evan You
  //vs-avatar(danger)
  //  i.bx.bx-user
  //vs-avatar(warn)
  //  i.bx.bx-hot
  //vs-avatar(dark)
  //  img(src="/avatars/avatar-5.png" alt="")
  CssConesInner
</template>

<script>
export default {
  data: () => ({
    users: [
      {
        id: "2",
        name: "vvluo",
        avatar: "",
        roles: ["前端"],
        skill: ["javascript", "typescript", "react", "vue", "html5", "小程序"],
        skill_ex: [
          "nodejs",
          "java web",
          "python flask",
          "react native",
          "flutter",
        ],
        city: "成都",
        time: "955",
      },
      {
        id: "3",
        name: "Edmond",
        avatar: "",
        roles: ["后端"],
        skill: ["java"],
        skill_ex: ["javascript", "vue"],
        city: "深圳",
        time: "955",
      },
      {
        id: "4",
        name: "Kr",
        avatar: "",
        roles: ["文案"],
        skill: ["写作"],
        skill_ex: ["运营", "html", "css", "javascript"],
        city: "绍兴",
        time: "0",
      },
      {
        id: "5",
        name: "觉",
        avatar: "",
        roles: ["前端", "后端"],
        skill: ["go", "javascript", "php", "python", "nodejs"],
        skill_ex: ["dart", "java", "c"],
        city: "杭州",
        time: "0",
      },
      {
        id: "6",
        name: "fmrx",
        avatar: "",
        roles: ["后端"],
        skill: ["java", "python"],
        skill_ex: ["flutter", "小程序"],
        city: "北京",
        time: "996",
      },
    ],
  }),
  methods: {
    open_user(id) {
      this.$router.push({ path: "/user/" + id });
    },
  },
};

//div.center.con-avatars(style="display:flex;")
//div.container_index
//  div 开发组能力范围 |
//  div 开发组成员 |
//  div 输出能力
</script>

<style>
.container_index_back {
  margin: 0 auto;
  min-height: calc(100vh - 120px - 64px);
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  text-align: center;
}
</style>
