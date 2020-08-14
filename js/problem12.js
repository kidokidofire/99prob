/*
ランレングス圧縮した配列をデコードするdecodeModified関数を実装せよ。
decodeModified([[2, 1], 2, 1, [2, 2], [4, 3]]) // [1, 1, 2, 1, 2, 2, 3, 3, 3, 3]
decodeModified([[4, 'a'], 'b', [2, 'c'], [2, 'a'], 'd', [4, 'e']]) // aaaabccaadeeee
*/

const decodeModified = ary => {
    let decodedAry = [];

    ary.forEach(element => {
        if(typeof element === "object"){
            decodedAry.push(...Array(element[0]).fill(element[1]));
        }else{
            decodedAry.push(element);
        }
    });

    if(typeof decodedAry[0] === "string"){
        decodedAry = decodedAry.join("");
    }

    return decodedAry;
}


/*-------テスト-------*/ 
const test1 = [[2, 1], 2, 1, [2, 2], [4, 3]];
const test2 = [[4, 'a'], 'b', [2, 'c'], [2, 'a'], 'd', [4, 'e']];

console.log(decodeModified(test1));
console.log(decodeModified(test2));