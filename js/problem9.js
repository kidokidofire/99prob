/*
配列や文字列の同じ要素の繰り返しを配列としてまとめて返すpack関数を実装せよ。
pack([1, 1, 2, 1, 2, 2, 3, 3, 3, 3]) // [[1, 1], [2], [1], [2, 2], [3, 3, 3, 3]]
pack('aaaabccaadeeee') // ['aaaa', 'b', 'cc', 'aa', 'd', 'eeee']
*/

const pack = (ary) => {
    let packedAry = [];
    let tmpIndex = 0;

    for(let i = 1; i < ary.length; i++){
        if(ary[i] != ary[tmpIndex]){
            packedAry.push(ary.slice(tmpIndex, i));
            tmpIndex = i;
        }
    }
    packedAry.push(ary.slice(tmpIndex, ary.length));

    return packedAry;
}



/*-------テスト-------*/ 
const test1 = [1, 1, 2, 1, 2, 2, 3, 3, 3, 3];
const test2 = 'aaaabccaadeeee';

console.log(pack(test1));
console.log(pack(test2));