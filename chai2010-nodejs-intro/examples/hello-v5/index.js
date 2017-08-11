const electron = require('electron')

let win = null

electron.app.on('ready', () => {
    if(win == null) {
        win = new electron.BrowserWindow()
		win.webContents.openDevTools() // <-- 打开调试窗口

        win.loadURL(`file:///${__dirname}/index.html`)
    }
})
