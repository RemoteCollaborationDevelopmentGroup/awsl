<style lang="sass" scoped>
h1
  margin-bottom: 2rem
.repay
  width: 16rem
  height: 8rem
  background: rgba(245,145,50,.25)
  border-radius: 8px
  margin: 1rem
  padding: 1rem 2rem
  position: relative
  .type
    top: -1rem
    font-size: 8rem
    font-weight: 900
    position: absolute
    color: rgba(255,255,255,.25)
    transform: rotate(7deg)
  .quantity
    font-size: 4rem
    font-weight: 600
    color: rgba(245,145,50,.8)
    text-align: right

.task
  display: flex
  .left
    flex: 1
  .right
    width: 24rem
    border-left: 3px dashed #eee
    padding: 0 1rem
.member-list
  .member-item
    opacity: .5
  .member-item.accept
    opacity: 1
    //background: #f5c5fc
.task-list
  list-style: none
  .task-item
    margin: 1rem
    padding: 1rem
    background: #f5f5f5
    .task-name
      font-size: 1.2rem
      font-weight: 600
.material-list
  color: #aa6655
  padding: 0 1rem
  margin: 1rem 0
  .material-item

.material-list.secret
  color: #66aa55
</style>

<template lang="pug">
AppDrawer
  .task
    .left
      h1 远程项目协作
      div (执行可见)
        button 关注项目(当有进展时向我推送通知)
        button 跟随发布者(当他有新发布时向我推送通知)

      p 项目介绍
      ul.task-list
        li.task-item(v-for="task in tasks", :key="task.id")
          .task-name {{ task.name }}
          p {{ task.name }} + 价格/加价请求及报价说明
          p 任务详情(前置条件)
            ul
              li 1. 基于现有php项目(laraval框架)
              li 2. 实现通讯功能

          p 任务资料
            ul.material-list.secret
              li(v-for="item in task.material", v-if="item.public") {{ item.name }}
            ul.material-list
              li(v-for="item in task.material", v-if="!item.public") {{ item.name }}
          p 验收标准(验收依据)
          p 邀请参与者/接受邀请者/拒绝邀请者(谁来执行/议价过程)
            ul
              li Kanako (议价)
                p Kanako 提出报价,将任务价格设为 400
                p 说明: 功能略复杂且耗时
                button 同意此议价
                button 拒绝此议价

                p Sato 提出报价,将任务价格设为 400
                p 说明: 功能略复杂且耗时
                button 同意此议价
                button 拒绝此议价

                p Last 提出报价,将任务价格设为 400
                p 说明: 功能略复杂且耗时
                button 同意此议价
                button 拒绝此议价

                p Kanako 接受了定价, 但提出更改架构
                p Kanako 的更改方案, 此方案的相对优劣
                button 同意此方案
                button 拒绝此方案(理由)

                p Kanako 正在执行此任务, Ena 请求加入此任务的协作
                button 同意加入协作
                button 拒绝加入

                p Ena 加入了协作
                p 任务分配, 资金分配(由kanako)
                button Ena 同意此方案
                button Ena 拒绝此方案
                button 议价

              li Last (接受邀请)
              li Kana (邀请中..)
              li Sato (已拒绝)
                p 无原因
              li ... (收起的所有邀请)
              li
                button + 邀请参与者 (指定或按照条件检索)
                button 发送邀请链接

      p 项目参与者
      ul
        li 开发(任务)
        li 开发(任务)
        li 开发(聘任)
        li 架构(聘任)
        li 运维(聘任)
        li 运营(聘任)

      p 项目流程(无限分级)
      ul
        li 1. 设计图
        li 2. 评估
        li 3. 开发/构建
        li 4. 测试/debug
        li 5. 上线/运营

      p 公开状态(只允许邀请者查看项目信息)
      p 公开状态(只允许参与者查看项目详情)
      p 公开状态(特定组织成员才有权限查看特定信息)

      p 全部成员, 组成员(只有权限查看组内信息)
      ul.member-list
        li.member-item(
          v-for="member in members",
          :key="member.id",
          :class="{ accept: member.accept }",
          :title="member.accept ? '已经接受邀请' : '尚未接受邀请'"
        )
          span {{ member.name }}
        li 已被拒绝的邀请

      button + 邀请成员

      div
        p 标签(自动筛选)
      .repay
        .type ￥
        .quantity 500

      div
        p 任务物品/附件
      div
        p 发布者其它任务
      div
        p 相关推荐
      div 右侧讨论
    .right
      p 基于议题的多级对话
      Posts
</template>

<script>
export default {
  data: () => ({
    tasks: [
      {
        id: "xx1",
        name: "项目任务分发",
        material: [
          { name: "接受任务前, 可以读取的非敏感资料", public: true },
          { name: "接受任务前, 可以读取的非敏感资料", public: true },
          { name: "接受任务后, 可以读取的任务资料", public: false },
          { name: "接受任务后, 可以读取的任务资料", public: false },
        ],
      },
      //{ id: "xx2", name: "项目列表" },
      //{ id: "xx3", name: "任务3" },
      //{ id: "xx4", name: "任务4" },
      //{ id: "xx5", name: "任务5" },
    ],
    teams: [
      { id: "xx1", name: "小组1", color: "#cc1414" },
      { id: "xx2", name: "小组2", color: "#cc1414" },
      { id: "xx3", name: "小组3", color: "#cc1414" },
      { id: "xx4", name: "小组4", color: "#cc1414" },
      { id: "xx5", name: "小组5", color: "#cc1414" },
    ],
    members: [
      { id: "xx1", name: "成员1", color: "#cc1414", accept: true },
      { id: "xx2", name: "成员2", color: "#cc1414", accept: true },
      { id: "xx3", name: "成员3", color: "#cc1414", accept: true },
      { id: "xx4", name: "成员4", color: "#cc1414", accept: false },
      { id: "xx5", name: "成员5", color: "#cc1414", accept: false },
    ],
  }),
};
</script>
