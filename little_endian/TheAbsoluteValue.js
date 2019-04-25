function hi3(num, base) {
	let hi = num.toString(16);
	if (hi.length % 2 === 1)
		hi = '0' + hi;
	return Number(`0x${hi.match(/.{1,2}/g).reverse().join('')}`)
}
