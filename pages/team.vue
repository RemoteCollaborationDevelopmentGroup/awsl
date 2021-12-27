<template lang="pug">
.team
  h1 Team (小组)
  p 条件检索: 标签, 特性, 风格, 节奏
  ul.team-list
    li.team-item(v-for="team in teams") {{ team.name }}
  //CssConesInner
</template>

<script>
export default {
  data: () => ({
    teams: [
      { id: "x1", name: "Lisya Rust", member: 2 },
      { id: "x2", name: "Lisya Go", member: 2 },
    ],
  }),
};
</script>
