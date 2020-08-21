/*
配列や文字列を指定した位置で分けるsplit関数を実装せよ。
split([1, 2, 3, 4], 2) // [[1, 2], [3, 4]]
split('abcdefghik', 3) // ['abc', 'defghik']
*/

const split = (props) => {
    const {
        ary, splitIndex,
    } = props;

    switch(typeof ary){
        case 'string':
            splitAry = exe([...ary], splitIndex);
            return [splitAry[0].join(''), splitAry[1].join('')];
        case 'object':
            return exe(ary, splitIndex);
    }
}

const exe = (ary, splitIndex) => {
    let frontAry = [];
    let backAry = [];
    let index = 0;

    ary.forEach(element => {
        if(index < splitIndex){
            frontAry.push(element);
        }else if(index >= splitIndex){
            backAry.push(element);
        }
        index += 1;
    });
    return [frontAry, backAry];
};

/*-------テスト-------*/ 
const test1 = {ary: [1, 2, 3, 4], splitIndex: 2};
const test2 = {ary: 'abcdefghik', splitIndex: 3};

console.log(split(test1));
console.log(split(test2));