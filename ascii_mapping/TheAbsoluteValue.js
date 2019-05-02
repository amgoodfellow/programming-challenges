function hi2(str) {
	return str.split('').reduce((a,b) => a + b.charCodeAt(0), 0) / str.length;
}
