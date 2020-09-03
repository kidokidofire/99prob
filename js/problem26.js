/*
m個の要素からn個を選んだ組み合わせを返すcombinations関数を実装せよ。
combinations('abcdef', 2) // ['ab', 'ac', 'ad', 'ae', 'af', 'bc', 'bd', 'be', 'bf', 'cd', 'ce', 'cf', 'de', 'df', 'ef']
*/

const combinations = (props) => {
    const {
        ary, count,
    } = props;

    const combIndexes = [...Array(count).keys()]
    return exe(ary, count, combIndexes)
};

const exe = (ary, count, combIndexes, accAry = []) => {
    accAry.push(combIndexes.map(i => ary[i]).join(''));

    if(combIndexes[0] === ary.length - count){
        return accAry;
    };
    combIndexes = updateCombIndexes(combIndexes, count, ary.length);

    return exe(ary, count, combIndexes, accAry);
};

const updateCombIndexes = (combIndexes, size, max) => {
    if(combIndexes[size-1] + 1 < max){
        combIndexes[size-1] = combIndexes[size-1] + 1;
    }else{
        for(let i=2; i<=size; i++){
            if(combIndexes[size-i] + i < max){
                combIndexes[size-i] = combIndexes[size-i] + 1;
                for(let sub=size-i+1; sub<size; sub++){
                    combIndexes[sub] = combIndexes[sub-1]+1;
                }
            }
        }
    }
    return combIndexes;
};

/*-------テスト-------*/ 
const test1 = {ary: 'abcdef', count: 2};

console.log(combinations(test1));