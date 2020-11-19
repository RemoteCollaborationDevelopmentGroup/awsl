<template lang="pug">
div.container_index
  vs-avatar(primary v-for="user in users" @click="open_user(user.id)")
    template(v-if="user.avatar===''" #text) {{ user.name }}
    img(v-else src="user.avatar")

  //vs-avatar(success)
  //  template(#text) Evan You
  //vs-avatar(danger)
  //  i.bx.bx-user
  //vs-avatar(warn)
  //  i.bx.bx-hot
  //vs-avatar(dark)
  //  img(src="/avatars/avatar-5.png" alt="")
</template>

<script>
import users from '~/assets/users.js'
export default {
  data: () => ({
    users: users
  }),
  methods: {
    open_user(id) {
      this.$router.push({path:'/user/' + id})
    }
  }
}

//div.center.con-avatars(style="display:flex;")
//div.container_index
//  div 开发组能力范围 |
//  div 开发组成员 |
//  div 输出能力
</script>

<style>
.container_index {
  margin: 0 auto;
  min-height: calc(100vh - 120px - 64px);
  display: flex;
  justify-content: center;
  align-items: center;
  text-align: center;
}
</style>
