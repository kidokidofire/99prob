/*
配列や文字列を逆順にして返すreverse関数を実装せよ。
ただし、Array.reverseは使用しないこと。
また、絵文字の長さは1と数えるようにすること。
reverse('A man, a plan, a canal, panama!💃') // 💃!amanap ,lanac a ,nalp a ,nam A
reverse([1, 2, 3, 4]) // [4, 3, 2, 1]
*/

const reverse = (ary) => {
    let i = 0;
    let flg = ary[0];
    let rev = [];
    //ary = [...ary];

    while(flg != undefined){
        flg = ary[i+1];
        rev.unshift(ary[i]);
        i ++;
    }
    if(typeof ary === "string"){
        rev = rev.join('')
    }
    return rev;
}


/*-------テスト-------*/ 
const test1 = 'A man, a plan, a canal, panama!💃';
const test2 = [1, 2, 3, 4];


console.log(reverse(test1));
console.log(reverse(test2));
console.log(test1);