
function sum(n) {
	let s = 0
	for(let i = 0; i <= n; i++) {
		s += i
	}
	return s
}

if(require.main === module) {
	let args = process.argv.splice(2)
	if(args.length > 0) {
		console.log(sum(args[0]|0))
	}
}
