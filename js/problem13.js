/*
pack関数のように重複を含む配列を作らずに、直接ランレングス圧縮するencodeDirect関数を実装せよ。
encodeDirect([1, 1, 2, 1, 2, 2, 3, 3, 3, 3]) // [[2, 1], 2, 1, [2, 2], [4, 3]]
encodeDirect('aaaabccaadeeee')
// [[4, 'a'], 'b', [2, 'c'], [2, 'a'], 'd', [4, 'e']]
*/

const encodeDirect = ary => {
    let encodedAry = [];
    let count = 0;

    if(typeof ary === 'string'){
        ary = [...ary];
    }

    let lookingElement = ary[0];

    ary.forEach(element => {
        if(element != lookingElement){
            insertAry(count, lookingElement, encodedAry);
            count = 0;
            lookingElement = element;
        }
        count ++;
    });
    insertAry(count, lookingElement, encodedAry);


    return encodedAry;
}

const insertAry = (count, element, ary) => {
    if(count === 1){
        ary.push(element);
    }else{
        ary.push([count, element]);
    }
};


/*-------テスト-------*/ 
const test1 = [1, 1, 2, 1, 2, 2, 3, 3, 3, 3];
const test2 = 'aaaabccaadeeee';

console.log(encodeDirect(test1));
console.log(encodeDirect(test2));