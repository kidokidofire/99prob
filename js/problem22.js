/*
指定された範囲内のすべての整数または文字を含む配列を生成するrange関数を実装せよ。
range(4, 9) // [4, 5 ,6, 7, 8, 9]
range('a', 'z') // ['a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 
*/

const range = (props) => {
    const {
        start, end
    } = props;

    switch(typeof start){
        case 'string':
            return range_string(start, end, []);
        case 'number':
            return range_number(start, end, []);
    }
}

const range_number = (start, end, rangeAry) => {
    rangeAry.push(start)

    if(start === end){
        return rangeAry;
    }
    return range_number(start + 1, end, rangeAry);
};

const range_string = (start, end, rangeAry) => {
    rangeAry.push(start)

    if(start === end){
        return rangeAry;
    }
    start = String.fromCharCode(start.charCodeAt(0) + 1);
    return range_string(start, end, rangeAry);
};


/*-------テスト-------*/ 
const test1 = {start: 4, end: 9};
const test2 = {start: 'a', end: 'z'};

console.log(range(test1));
console.log(range(test2));