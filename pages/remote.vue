<template lang="pug">
div.container_remote
  div
    img(src="/emmm.jpg" alt="awsl" style="height:10rem;")
    p 施工中..
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
</style>