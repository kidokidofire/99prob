/*
配列や文字列の最後の要素を取り出すlast関数を実装せよ。
last([1, 2, 3, 4]) // 4
last('xyz') // z
*/

const last = (ary) => {
    return ary[ary.length - 1];
}

test1 = [1, 2, 3, 4];
test2 = 'xyz';

console.log(last(test1));
console.log(last(test2));