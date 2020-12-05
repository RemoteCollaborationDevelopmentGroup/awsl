export default {
    user: new Map(),
    task: new Map(),
    project: new Map(),
    group: new Map(),
    tag: new Map(),
}

function demo() {
    let satori = user(12)
}

// 初始化时, 从缓存载入数据到内存
// 当使用某个数值时, 从内存读取数值
// 当内存未命中时, 加载数据从网络
// 当网络返回数据时, 写入内存和缓存
function Usre(id) {

}

class 数据 {
    constructor() {
        // 从缓存载入

    }
    读取数据(id) {
        return this.从内存读取() || this.从网络读取()
    }
    从网络读取(id) {
        var data = "2333"
        写入内存(data)
        写入缓存(data)
        return data
    }
    从内存读取(id) { }
}

class 模型 {
    get 缓存() { }
    set 缓存() { }
}

class 内存 {
    constructor() {
        this.list = new Map()
    }
    写入(key, value) { }
    读取(key) { }
    移除(key) { }
}

class 缓存 {
    写入() { }
    读取() { }
    移除() { }
}

// 缓存又分为本地存储和cookie, 一般避免使用cookie
// 缓存只存储重要数据, 或是刷新页面时的数据交换