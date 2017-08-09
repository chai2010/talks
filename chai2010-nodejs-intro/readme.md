<!--
// Copyright 2016 ChaiShushan <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
-->

<!-- ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++  -->

<!-- *** 横向分隔, === 竖向分隔, --- 水平线, Note: 讲稿注释  -->

<!--
Reveal.js 可能会需要 AJAX 异步加载 Markdown 文件, 可以在当前目录启动一个 http 服务.

以下是常见的临时启动 http 服务器的方式:

	NodeJS
	npm install http-server -g
	http-server

	Python2
	python -m SimpleHTTPServer

	Python3
	python -m http.server

	Golang
	go run server.go

启动后, 本地可以访问 http://127.0.0.1:port, 其中 port 为端口号, 命令行有提示.

幻灯片操作: F键全屏, S键显示注解, ESC大纲模式, ESC退出全屏或大纲模式, ?显示帮助

-->

<!-- ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++  -->

# Node.JS & Electron 简介

<h4 style="text-align:center;">
	<a href="https://github.com/chai2010" target="_blank">
		chaishushan@gmail.com
	</a>
</h4>

Note: 讲稿注释

<!-- ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++  -->
***

## Node.JS 是什么?

---

- Node.JS 是不依赖浏览器的独立的 JavaScript 引擎
- 支持模块化的 JavaScript 编程
- 标准库: 文件系统, 网络, ...

---

- 网站: http://nodejs.org
- 安装: 安装包安装

<!-- ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++  -->
***

## Electron 是什么?

---

- Node.JS 是一个命令行 JS 引擎
- Electron 则是一个支持 GUI 对 JS 引擎
- Electron 在 Node.JS 的基础上增加了对 Chrome 对支持
- Electron 是 Node.JS 对超集

---

- Electron 将 Node.JS 的标准库带入了 Chrome 中
- Electron 将 Node.JS 的模块化编程带入了 Chrome 中

---

- 网站: https://github.com/electron/electron
- 安装: `npm install electron -g`

<!-- ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++  -->
***

## 你好, 世界

---

examples/hello-v1.js:

```js
console.log('你好, 世界')
```

---

- NodeJS: `node examples/hello-v1.js`
- Electron: `electron examples/hello-v1.js`

---

- Node.JS 执行完成后退出了
- Electron 执行完成后没有退出

<!-- ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++  -->
***

## 你好, 世界 - V2

---

examples/hello-v2.js:

```js
console.log('你好, 世界')
process.exit(0)
```

---

- process 是 NodeJS 内置的表示当前进程对象
- 调用 `process.exit(0)` 主动退出
- Node.JS 和 Electron 行为一致

---

- Node.JS 和 Electron 其实是类似的工具
- Electron 只是标准库比 Node.JS 多一些包而已


<!-- ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++  -->
***

## 你好, 世界 - V3

---

examples/hello-v3/index.js:

```js
const electron = require('electron')

electron.app.on('ready', () => {
	let win = new electron.BrowserWindow()
	win.loadURL(`file:///${__dirname}/index.html`)
})
```
---

examples/hello-v3/index.html:

```html
<h1>你好, 世界</h1>
```

---

- Electron 运行: `electron examples/hello-v3`

===

## 效果图
---

TODO

===

## 改进: 避免 win 被提前回收

---

```js
const electron = require('electron')

let win = null

electron.app.on('ready', () => {
	if(win == null) {
		win = new electron.BrowserWindow()
		win.loadURL(`file:///${__dirname}/index.html`)
	}
})
```

---

- `win` 改为模块级变量


===

## 打开调试窗口

---

```js
win.webContents.openDevTools()
```

---

截图


<!-- ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++  -->
***

## TODO


<!-- ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++  -->
***

## 参考资源

#### http://nodejs.org

#### https://developer.mozilla.org/en-US/

<!-- ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++  -->
***

## Thank you

#### [chaishushan@gmail.com](https://github.com/chai2010)

<!-- ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++  -->
