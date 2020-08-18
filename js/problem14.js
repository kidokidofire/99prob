/*
配列や文字列の要素を複製するdupli関数を実装せよ。
dupli([1, 2, 3]) // [1, 1, 2, 2, 3, 3]
dupli('abc') // aabbcc
*/

const dupli = (ary) => {
    switch(typeof ary){
        case 'string':
            ary = [...ary];
            return exe(ary).join('');
        case 'object':
            return exe(ary);   
    }
}

const exe = (ary) => {
    let dupliAry = [];

    ary.forEach(element => {
        dupliAry.push(element);
        dupliAry.push(element);
    })
    return dupliAry;
}

/*-------テスト-------*/ 
const test1 = [1, 2, 3];
const test2 = 'abc';

console.log(dupli(test1));
console.log(dupli(test2));