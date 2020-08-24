/*
選択した範囲を配列や文字列から取り出すslice関数を実装せよ。
ただし、Array.slice、Array.spliceは使用しないこと。
slice([1, 2, 3, 4], 2, 4) // [2, 3, 4]
slice('abcdefghik', 3, 7) // cdefg
*/

const slice = (props) => {
    const {
        ary, startIndex, endIndex
    } = props;

    switch(typeof ary){
        case 'string':
            return exe([...ary], startIndex, endIndex).join('');
        case 'object':
            return exe(ary, startIndex, endIndex);
    }
}

const exe = (ary, startIndex, endIndex) => {
    let slicedAry = [];

    ary.forEach((element, index) => {
        if(index >= startIndex - 1 && index < endIndex){
            slicedAry.push(element);
        }
    });
    return slicedAry;
};

/*-------テスト-------*/ 
const test1 = {ary: [1, 2, 3, 4], startIndex: 2, endIndex: 4};
const test2 = {ary: 'abcdefghik', startIndex: 3, endIndex: 7};

console.log(slice(test1));
console.log(slice(test2));