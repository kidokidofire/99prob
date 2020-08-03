/*
配列や文字列の長さを返すlength関数を実装せよ。
ただし、Array.lengthやString.lengthは使用しないこと。
また、絵文字の長さは1と数えるようにすること。
length([123, 456, 789]) // 3
length('💃Hello, World!💃') // 15
*/

const length = (ary) => {
    let i = 0;
    let flg = true;
    ary = [...ary];

    while(flg){
        flg = ary[i];
        i ++;
    }
    return i - 1;
}


/*-------テスト-------*/ 
const test1 = [123, 456, 789];
const test2 = '💃Hello, World!💃';


console.log(length(test1));
console.log(length(test2));