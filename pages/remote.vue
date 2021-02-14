<template lang="pug">
.container_remote
  div
    div
      img.NuxtImg(src="/emmm.jpg", alt="awsl", style="height: 10rem")
      p 施工中..
    div
      span 软件 网络 交互 产品设计 服务
    div
      TaskItem(v-for="id in tasklist", :key="id", :data="id")
</template>

<script>
import Logo from "~/components/Logo.vue";
import VuesaxLogo from "~/components/VuesaxLogo.vue";

export default {
  data: () => ({
    tasklist: [1, 2, 3, 4, 5, 6],
    //[
    //  { id: 1, name: "一条任务信息", info: "其它参数信息" },
    //  { id: 5, name: "一条任务信息", info: "其它参数信息" },
    //  { id: 4, name: "一条任务信息", info: "其它参数信息" },
    //  { id: 3, name: "一条任务信息", info: "其它参数信息" },
    //  { id: 2, name: "一条任务信息", info: "其它参数信息" },
    //],
  }),
  components: {
    Logo,
    VuesaxLogo,
  },
};
</script>

<style lang="sass" scoped>
.container_remote
  margin: 0 auto
  height: calc(100vh - 120px - 64px)
  display: flex
  justify-content: center
  align-items: center
  text-align: center

.NuxtImg
  animation: 1s appearx
  margin: auto
@keyframes appearx
  0%
    opacity: 0
</style>
