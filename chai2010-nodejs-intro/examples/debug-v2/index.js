const electron = require('electron')

let win = null

electron.app.on('ready', () => {
    if(win == null) {
        win = new electron.BrowserWindow()
		win.webContents.openDevTools() // <-- 调试代码

        win.loadURL(`file:///${__dirname}/index.html`)

		console.log('win ready') // <-- 调试代码
    }
})
