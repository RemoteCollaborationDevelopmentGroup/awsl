<template lang="pug">
div.container_remote
  div
    div
      img.NuxtImg(src="/emmm.jpg" alt="awsl" style="height:10rem;")
      p 施工中..
    div
      span 软件 网络 交互 产品设计 服务
    div
      Task(v-for="i in 8" :key="i") {{ i }}
</template>

<script>
import Logo from '~/components/Logo.vue'
import VuesaxLogo from '~/components/VuesaxLogo.vue'

export default {
  components: {
    Logo,
    VuesaxLogo
  }
}
</script>

<style lang="styl">
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