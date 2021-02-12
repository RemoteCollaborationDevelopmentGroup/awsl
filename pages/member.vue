<template lang="pug">
.member
  //p 少年, 你渴望力量吗? (签订契约成为魔法少女吧!)
  //p 众人拾柴火焰高 (上午一把火,下午派出所)
  vs-table(v-model="selected")
    template(#thead)
      vs-tr
        vs-th name
        vs-th city
        vs-th roles
        vs-th skill
        vs-th skill_ex
        vs-th time
        vs-th mark
    template(#tbody)
      vs-tr(
        v-for="(user, i) in users",
        :data="user.id",
        :is-selected="selected == user.id",
        :key="i"
      )
        vs-td {{ user.name }}
        vs-td {{ user.city }}
        vs-td {{ user.roles }}
        vs-td {{ user.skill }}
        vs-td {{ user.skill_ex }}
        vs-td {{ user.time }}
        vs-td {{ user.mark }}
    //template(#tfoot)
    //  vs-tr
    //    vs-td footer
</template>

<style lang="sass" scoped>
.member
  max-width: 1280px
  margin: auto
//  table
//    border-radius: 8px
//    overflow: hidden
//    transition: opacity 0.3s cubic-bezier(0.22, 0.61, 0.35, 1)
//    thead
//      background: rgba(0,0,0,.05)
//      tr
//        th
//          padding: .5rem 1rem
//    tbody
//      tr
//        td
//          padding: .5rem 1rem
//          span
//            margin: 0 .25rem
</style>

<script>
import users from "~/assets/users.json";
export default {
  data: () => ({
    selected: null,
    users: [],
  }),
  mounted() {
    //console.log(users);
    this.users = users;
  },
};
</script>