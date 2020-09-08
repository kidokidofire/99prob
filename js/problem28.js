/*
配列の配列を、要素の長さでソートするlsort関数を実装せよ。
また、要素の長さの頻度順にソートするlfsort関数を実装せよ。
lsort(["abc","de","fgh","de","ijkl","mn","o"]) // ["o","de","de","mn","abc","fgh","ijkl"]
lfsort(["abc", "de", "fgh", "de", "ijkl", "mn", "o"]) // ["ijkl","o","abc","fgh","de","de","mn"]
*/

const lsort = (ary, length=1, returnAry=[]) => {
    returnAry.push(...ary.filter((element) => element.length === length));
    if(returnAry.length === ary.length){
        return returnAry;
    };

    return lsort(ary, length+1, returnAry);
};

const lfsort = (ary) => {
    lfpack = pack(ary);
    return lfspread(lfpack, ary.length);
};

const pack = (ary, length=1, returnAry=[]) => {
    returnAry.push(ary.filter((element) => element.length === length));
    if(returnAry.length === ary.length){
        return returnAry;
    };

    return pack(ary, length+1, returnAry);
}

const lfspread = (ary, elementCount, length=1, returnAry=[]) => {
    ary.map((element) => {
        if(element.length === length){
            returnAry = [...returnAry, ...element];
        }
    });
    if(returnAry.length === elementCount){
        return returnAry;
    };

    return lfspread(ary, elementCount, length+1, returnAry);
}

/*-------テスト-------*/ 
const test1 = ["abc","de","fgh","de","ijkl","mn","o"];
const test2 = ["abc", "de", "fgh", "de", "ijkl", "mn", "o"];

console.log(lsort(test1));
console.log(lfsort(test2));