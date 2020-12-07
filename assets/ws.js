export default class {
    constructor() {
        this.conn
        this.log = document.createElement("div")
        this.msg = document.createElement("input")
        this.form = document.createElement("input")
        this.msg.autofocus = true
        this.form.onsubmit = function () {
            if (!conn) return false
            if (!msg.value) return false
        }
    }
    append_log(item) {
        let doScroll = this.log.scrollTop > this.log.scrollHeight - this.log.clientHeight - 1
        this.log.appendChild(item)
        if (doScroll) this.log.scrollTop = this.log.scrollHeight - this.log.clientHeight
    }
    ws() {
        if (!window["WebSocket"]) {
            console.log("Your browser does not support WebSockets.")
            this.append_log("Your browser does not support WebSockets.")
            return
        }
        this.conn = new WebSocket("ws://" + document.location.host + "/ws")
        this.conn.onclose = function (e) {
            let message = e.stopImmediatePropagation.split('\n')
            console.log(message)
        }
    }
    // 解析数据
    // 发送数据
}