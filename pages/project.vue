<template lang="pug">
.task
  h1 Task(project 与 task 可以合并)
  p this task list
  ul.nav
    li.active 招募中
    li 执行中
    li 归档
  .list
    NuxtLink.item(
      v-for="item in list",
      :key="item.id",
      :to="`/project/${item.name}`"
    )
      h5 {{ item.name }}
      p {{ item.info }}
      div
        button 关注
        button 领取
  .project
    ul
      li(v-for="project in projects", :key="project.id")
        .header
          .name {{ project.name }}
          button 关注项目
        .task
          ul
            li(v-for="task in project.tasks")
              .name
                span {{ task.name }}
                button 领取任务
  NuxtChild
</template>

<style lang="sass" scoped>
.list
  width: 600px
  margin: 1rem auto
  padding: 0
  .item
    display: block
    background: rgba(0, 0, 0, 0.05)
    border-radius: 5px
    margin: 0.5rem
    padding: 0.5rem 1rem
    color: #555555
    text-decoration: none
    button
      margin: .25rem .25rem 0 0
.project
  >ul
    width: 600px
    margin: 1rem auto
    padding: 0
    list-style: none
    >li
      background: rgba(0,0,0,.05)
      border-radius: 8px
      padding: 1rem
      margin: 1rem 0
      min-height: 4rem
      >.header
        display: flex
        .name
          flex: 1
      .task
        >ul
          padding: 1rem
          list-style: none
          >li
            background: #ffffff
            padding: .5rem 1rem
            button
              margin: .25rem .25rem 0 0
              background: rgba(255,255,255,0.5)
              border-radius: 4px
              border: none
ul.nav
  list-style: none
  display: flex
  background: rgba(0,0,0,.05)
  border-radius: 2rem
  margin: auto
  padding: 0
  width: 600px
  li
    border-radius: 2rem
    padding: .5rem 1rem
    flex: 1
    text-align: center
  li.active
    color: rgb(0, 197, 142)
    background: rgb(16, 135, 117)
</style>

<script>
export default {
  data: () => ({
    projects: [
      {
        id: 1,
        name: "Fecyo 小程序",
        tasks: [{ id: 1, name: "T1" }],
      },
      {
        id: 2,
        name: "本地SEO",
        tasks: [{ id: 2, name: "组织规划", developer: "Last" }],
      },
      {
        id: 3,
        name: "0x00000000",
        tasks: [{ id: 3, name: "poeject 版块设计" }],
      },
    ],
    list: [
      {
        id: "12",
        name: "SPASXASXMK",
        info: "基于 wps sdk 实现供 app 调用的文档页面",
        role: ["前端", "后端"],
        member: [],
        status: "执行中",
        reward: "￥0.00",
      },
    ],
  }),
};
</script>