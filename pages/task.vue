<template lang="pug">
.task
  h1 Task
  p this task list
  div
    NuxtLink(v-for="(item, i) in list", :key="i", :to="`/task/${item.name}`")
      h5 {{ item.name }}
      p {{ item.info }}
      div
        button 关注
        button 领取
  NuxtChild
</template>

<style lang="sass" scoped>
.task
  a
    display: block
    background: rgba(0, 0, 0, 0.05)
    border-radius: 5px
    margin: 0.5rem
    padding: 0.5rem 1rem
    color: #555555
    text-decoration: none
    button
      margin: .25rem .25rem 0 0
</style>

<script>
export default {
  data: () => ({
    list: [
      {
        name: "Fecyo 小程序",
        info: "基于fecyo目前的系统，搞一套uniapp的单商户系统",
      },
      {
        name: "WPS SDK",
        info: "基于 wps sdk 实现供 app 调用的文档页面",
      },
    ],
  }),
};
</script>