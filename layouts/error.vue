<template lang="pug">
  div
    h2(v-if="error.statusCode == 404") 404 页面不存在
    h2(v-else) 500 服务器错误
      nuxt-link(to="/") 返回首页
</template>

<script>
export default {
  props: ['error']
}
</script>