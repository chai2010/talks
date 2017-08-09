const electron = require('electron')

electron.app.on('ready', () => {
	let win = new electron.BrowserWindow()
	win.loadURL(`file:///${__dirname}/index.html`)
})
