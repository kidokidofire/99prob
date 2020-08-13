/*
問10の結果を変更し、重複が無ければランレングス圧縮せずに要素を格納するencodeModified関数を実装せよ。
encodeModified([1, 1, 2, 1, 2, 2, 3, 3, 3, 3]) // [[2, 1], 2, 1, [2, 2], [4, 3]]
encodeModified('aaaabccaadeeee')
// [[4, 'a'], 'b', [2, 'c'], [2, 'a'], 'd', [4, 'e']]
*/

const encodeModified = ary => {
    let encodedAry = [];
    let packedAry = pack(ary);

    packedAry.forEach(element => {
        if(element.length === 1){
            encodedAry.push(element[0]);
        }else{
            let encodedElement = [];
            encodedElement.push(element.length);
            encodedElement.push(element[0]);
            encodedAry.push(encodedElement);
        }
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

console.log(encodeModified(test1));
console.log(encodeModified(test2));