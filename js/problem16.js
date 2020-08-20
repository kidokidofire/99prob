/*
nの倍数の位置の要素を配列や文字列から削除するdrop関数を実装せよ。
drop([1, 2, 3, 4], 2) // [1, 3] 
drop('abcdefghik', 3) // 'abdeghk'
*/

const drop = (props) => {
    const {
        ary, baseIndex,
    } = props;

    switch(typeof ary){
        case 'string':
            return exe([...ary], baseIndex).join('');
        case 'object':
            return exe(ary, baseIndex);   
    }
}

const exe = (ary, baseIndex) => {
    let droppedAry = [];

    for(let i=0; i<ary.length; i++){
        if((i+1) % baseIndex != 0){
            droppedAry.push(ary[i]);
        }
    }
    return droppedAry;
};

/*-------テスト-------*/ 
const test1 = {ary: [1, 2, 3, 4], baseIndex: 2};
const test2 = {ary: 'abcdefghik', baseIndex: 3};

console.log(drop(test1));
console.log(drop(test2));