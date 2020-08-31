/*
配列や文字列から指定された数ぶんだけランダムに要素を取り出すrndSelect関数を実装せよ。
rndSelect('abcdefgh', 3) // eda など
*/

const rndSelect = (props) => {
    const {
        ary, selectNum,
    } = props;

    switch(typeof ary){
        case 'string':
            return exe([...ary], selectNum).join('');
        case 'object':
            return exe(ary, selectNum);
    }
}

const exe = (originAry, selectNum, selectedAry = []) => {
    selectIndex = Math.floor( Math.random() * originAry.length );

    selectedAry.push(...originAry.splice(selectIndex, 1));
    selectNum -= 1;
    if(selectNum === 0){
        return selectedAry;
    }
    return exe(originAry, selectNum, selectedAry);
};

/*-------テスト-------*/ 
const test1 = {ary: 'abcdefgh', selectNum: 3};
const test2 = {ary: [1, 2, 3, 4, 5], selectNum: 2};

console.log(rndSelect(test1));
console.log(rndSelect(test2));