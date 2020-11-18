<template lang="pug">
div.container_index
  div 开发组能力范围 |
  div 开发组成员 |
  div 输出能力
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

<style>
.container_index {
  margin: 0 auto;
  min-height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
  text-align: center;
}
</style>
