/*
ネストしている配列を平坦（一次元配列）にして返すflatten関数を実装せよ。
ただし、Array.flatは使用しないこと。
flatten([1, [2, [3, 4], 5]]) // [1, 2, 3, 4, 5]
*/

const flatten = (ary) => {
    let flg = true;

    while(flg){
        flatAry = [];
        ary.forEach(function(elem){
            flatAry = flatAry.concat(elem);
            if(typeof elem === 'object'){
                flg = false;
            }
        })
        flg = !flg;
        ary = flatAry;
    }
    return flatAry;
}


/*-------テスト-------*/ 
const test1 = [1, [2, [3, 4], 5]];

console.log(flatten(test1));