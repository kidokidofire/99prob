/*
配列や文字列のn番目の要素を取り出すelementAt関数を実装せよ。
ただし、最初の要素は0番目ではなく、1番目と数える。
elementAt([1, 2, 3], 2) // 2
elementAt('JavaScript', 5) // S
*/

const elementAt = (ary, index) => {
    if(ary.length < index){
        return "index number is over"
    }
    return ary[index - 1];
}


/*-------テスト-------*/ 
const test1 = {
    ary : [1, 2, 3],
    index : 2
};
const test2 = {
    ary : 'JavaScript',
    index : 5
};


console.log(elementAt(test1.ary, test1.index));
console.log(elementAt(test2.ary, test2.index));