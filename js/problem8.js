/*
配列や文字列から同じ要素の繰り返しを排除して返すcompress関数を実装せよ。
compress([1, 1, 2, 1, 2, 2, 3, 3, 3, 3]) // [ 1, 2, 1, 2, 3 ]
compress('aaaabccaadeeee') // abcade
*/

const compress = (ary) => {
    switch(typeof ary){
        case 'string':
            ary = [...ary];
            return exe(ary).join('');
        case 'object':
            return exe(ary);   
    }
}

const exe = (ary) => {
    let compAry = [];
    let tmp;

    ary.forEach(function(elem){
        if(elem != tmp){
            compAry.push(elem);
        }
        tmp = elem;
    })
    return compAry;
}


/*-------テスト-------*/ 
const test1 = [1, 1, 2, 1, 2, 2, 3, 3, 3, 3];
const test2 = 'aaaabccaadeeee';

console.log(compress(test1));
console.log(compress(test2));