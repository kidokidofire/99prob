/*
pack関数を用いて、配列や文字列をランレングス圧縮するencode関数を実装せよ。
encode([1, 1, 2, 1, 2, 2, 3, 3, 3, 3]) // [[2, 1], [1, 2], [1, 1], [2, 2], [4, 3]]
encode('aaaabccaadeeee')
// [[4, 'a'], [1, 'b'], [2, 'c'], [2, 'a'], [1, 'd'], [4, 'e']]
*/

const encode = ary => {
    let encodedAry = [];
    let packedAry = pack(ary);

    packedAry.forEach(element => {
        let tmp = [];
        tmp.push(element.length);
        tmp.push(element[0]);
        encodedAry.push(tmp);
    });
    return encodedAry;
}

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

console.log(encode(test1));
console.log(encode(test2));