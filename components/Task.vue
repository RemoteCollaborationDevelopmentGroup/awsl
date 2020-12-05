<template lang="pug">
div.task(@click="click")
    span {{ name }}
    span {{ content }}
    span 
</template>
<style>
.task {
    padding: .5em;
    background: #ff9999;
}
</style>

<script>
// 组件中应当只有交互
// 没有对数据的 API 操作
export default {
    data: () => ({
        name: '任务名称',
        content: '主体信息详情和资源的引用',
        partner: [121],
        status: '', // 开放 | 执行 | 暂停 | 终止 | 结束 | 完成
    }),
    methods: {
        hover() {
            // hover 时预装载数据
        },
        click() {
            // click 时展开项, 但列表中的其它项目应能快速切换
            this.name = 'new task name'
            this.content = '0'
        }
        // 关注任务/参与任务
        // 任务独占/开放协作
        // 任务进度/任务完成
        // 更新任务
        // 回复 / 会话 ? IM信息
        // 节点信息 / 标签信息 / 源project / 发布者
    }
}
</script>