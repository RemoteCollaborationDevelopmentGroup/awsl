<template lang="pug">
.member
  p meber list
  table
    thead
      tr
        th name
        th city
        th roles
        th skill
        th skill_ex
        th time
        th mark
    tbody
      tr(v-for="user in users")
        td {{ user.name }}
        td {{ user.city }}
        td {{ user.roles }}
        td {{ user.skill }}
        td {{ user.skill_ex }}
        td {{ user.time }}
        td {{ user.mark }}
    //tfoot
    //  tr
    //    td footer
</template>

<style lang="sass" scoped>
.member
  table
    border-radius: 8px
    overflow: hidden
    transition: opacity 0.3s cubic-bezier(0.22, 0.61, 0.35, 1)
    thead
      background: rgba(0,0,0,.05)
      tr
        th
          padding: .5rem 1rem
    tbody
      tr
        td
          padding: .5rem 1rem
          span
            margin: 0 .25rem
</style>

<script>
import users from "~/assets/users.json";
export default {
  data: () => ({
    users: [],
  }),
  mounted() {
    //console.log(users);
    this.users = users;
  },
};
</script>