/*
配列や文字列の要素を左にn個ずらすrotate関数を実装せよ。
負の数を渡したときは右にずらすようにすること。
rotate([1, 2, 3], 1) // [2, 3, 1]
rotate('abcdefgh', 3) // defghabc
rotate('abcdefgh', -2) // ghabcdef
*/

const rotate = (props) => {
    const {
        ary, rotateCount,
    } = props;

    switch(typeof ary){
        case 'string':
            if(rotateCount > 0){
                return normal_rotate([...ary], rotateCount).join('');
            }else if(rotateCount < 0){
                return reverse_rotate([...ary], rotateCount).join('');
            };
            
        case 'object':
            if(rotateCount > 0){
                return normal_rotate(ary, rotateCount);
            }else if(rotateCount < 0){
                return reverse_rotate(ary, rotateCount);
            };
    }
    return ary;
}

const normal_rotate = (ary, rotateCount) => {

    if(rotateCount === 0){
        return;
    };

    ary.push(ary.shift());
    rotateCount -= 1;

    normal_rotate(ary, rotateCount);
    return ary;
};

const reverse_rotate = (ary, rotateCount) => {

    if(rotateCount === 0){
        return;
    };

    ary.unshift(ary.pop());
    rotateCount += 1;

    reverse_rotate(ary, rotateCount);
    return ary;
};

/*-------テスト-------*/ 
const test1 = {ary: [1, 2, 3], rotateCount: 1};
const test2 = {ary: 'abcdefgh', rotateCount: 3};
const test3 = {ary: 'abcdefgh', rotateCount: -2};

console.log(rotate(test1));
console.log(rotate(test2));
console.log(rotate(test3));