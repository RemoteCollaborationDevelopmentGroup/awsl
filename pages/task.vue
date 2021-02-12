<template lang="pug">
.task
  h1 Task
  .header
    .type
      span 开发
      span 设计
      span 产品
      span 文案
      span 市场

    p 筛选器(每次选择都以最佳筛选为主)
    // 任务类型: 前端, 后端 (大分类/小标签/项目风格/架构)
    // 排序方法: 时间, 价格, 热度
    .taglist
  .list
    .item(v-for="i in 4", :key="i")
      .name 任务名 (关注/不喜欢)
      .content 实施过程和方式
      .content 项目架构和技术栈
      .content 验收标准
      .info --某项目
</template>

<style lang="sass" scoped>
.task
  .header
    .type
      span
        margin: 0 1rem
  .list
    .item
      padding: 0 1rem
      margin: 2rem 0
      border-left: 2px solid rgba(0,0,0,.1)
</style>

<script>
export default {
  data: () => ({
    type: [
      { name: "开发", tag: [{ name: "" }] },
      { name: "设计", tag: [{ name: "UI" }, { name: "UX" }] },
      { name: "产品", tag: [{ name: "" }] },
      { name: "文案", tag: [{ name: "" }] },
      { name: "市场", tag: [{ name: "" }] },
    ],
  }),
};
</script>