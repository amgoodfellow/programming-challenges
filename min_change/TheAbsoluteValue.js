function hi4(amt, coins) {
	const temp = amt;
	if (!coins.length) return 0;
	let out = 0;
	for (let i = 0; i < coins.length; i++) {
		if (amt === 0) break;
		const blah = Math.floor(amt / coins[i]);
		if (blah === 0) continue;
		amt -= blah * coins[i];
		out += blah;
    	}
	return Math.min(...([out, hi4(temp, coins.slice(1))]).filter(val => val != 0))
}
