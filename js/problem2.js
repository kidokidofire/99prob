/*
配列や文字列の最後の一つ前の要素を取り出すbutLast関数を実装せよ。
butLast([1, 2, 3, 4]) // 3
butLast('abcdefghijklmnopqrstuvwxyz') // y
*/

const butlast = (ary) => {
    return ary[Math.max(ary.length - 2, 0)];
}


/*-------テスト-------*/ 
test1 = [1, 2, 3, 4];
test2 = 'abcdefghijklmnopqrstuvwxyz';

console.log(butlast(test1));
console.log(butlast(test2));