/*
配列や文字列が回文かどうかを返すisPalindrome関数を実装せよ。

isPalindrome([1, 2, 3]) // false
isPalindrome('たけやぶやけた') // true
isPalindrome([1, 2, 4, 8, 16, 8, 4, 2, 1]) // true
*/

const isPalindrome = (ary) => {
    let head = 0;
    let foot = ary.length - 1;
    let flg = true;
    ary = [...ary];

    while(flg && foot >= head){
        if(ary[head] != ary[foot]){
            flg = false
        }
        head ++;
        foot --;
    }
    return flg;
}


/*-------テスト-------*/ 
const test1 = [1, 2, 3];
const test2 = 'たけやぶやけた';
const test3 = [1, 2, 4, 8, 16, 8, 4, 2, 1];

console.log(isPalindrome(test1));
console.log(isPalindrome(test2));
console.log(isPalindrome(test3));