/*
配列や文字列をランダムに並び替えるrndPermu関数を実装せよ。
rndPermu('abcdef') // badcef など
*/

const rndPermu = (ary) => {
    switch(typeof ary){
        case 'string':
            return exe([...ary]).join('');
        case 'object':
            return exe(ary);
    }
};

const exe = (ary, accAry = []) => {
    movedElementIndex = Math.floor( Math.random() * ary.length );
    movedElement = ary.splice(movedElementIndex, 1)[0];
    accAry.push(movedElement);

    if(ary.length === 0){
        return accAry;
    }
    return exe(ary, accAry);
};

/*-------テスト-------*/ 
const test1 = 'abcdef';
const test2 = [1, 2, 3, 4, 5];

console.log(rndPermu(test1));
console.log(rndPermu(test2));