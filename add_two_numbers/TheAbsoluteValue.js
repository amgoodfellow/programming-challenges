function hi(list1, list2) {
    return String(+list1.reverse().join('') + +list2.reverse().join('')).split('').reverse();
}
