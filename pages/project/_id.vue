<template lang="pug">
AppDrawer
  .task
    .left
      h1 SAJCIOASJOIC
      div 
        p 任务介绍1
        p 任务介绍2
        p 任务判定标准
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
      Posts
</template>

<style lang="sass" scoped>
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
    width: 32rem
    border-left: 3px dashed #eee
    padding: 0 1rem
</style>