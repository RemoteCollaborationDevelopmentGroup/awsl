<template lang="pug">
.posts
  .post(v-for="item in 8", :key="item")
    .left
      .avatar
    .right
      .name Last
      .message 这这这..
</template>

<style lang="sass" scoped>
.posts
  .post
    display: flex
    margin: 1rem 0
    .left
      justify-content: center
      align-items: center
      .avatar
        width: 2rem
        height: 2rem
        background: #ccc
        border-radius: 50%
        margin: .5rem auto
    .right
      padding: 0 2rem
      .name
        color: #999
        font-size: .8rem
      .message
        background: #f5f5f5
        padding: 1rem
        border-radius: 1rem
        position: relative

        &::before
          content: ""
          position: absolute
          top: 10px
          left: -1rem
          border-top: 20px solid #f5f5f5
          border-left: 1rem solid transparent
          border-right: 1rem solid transparent
</style>
