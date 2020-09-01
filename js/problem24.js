/*
長さnの1以上m以下の乱数列を生成するdiffSelect関数を実装せよ。
diffSelect(6, 49) // [23, 1, 17, 33, 21, 37] など
*/

const diffSelect = (props) => {
    const {
        length, max,
    } = props;

    return exe(length, max);
}

const exe = (length, max, accAry = []) => {
    newElement = Math.ceil( Math.random() * max );

    accAry.push(newElement);
    if(accAry.length === length){
        return accAry;
    }
    return exe(length, max, accAry);
};

/*-------テスト-------*/ 
const test1 = {length: 6, max: 49};
const test2 = {length: 10, max: 5};

console.log(diffSelect(test1));
console.log(diffSelect(test2));