/*
配列や文字列の指定した位置に要素を挿入するinsertAt関数を実装せよ。
insertAt(5, [1, 2, 3, 4], 3) // [1, 2, 5, 3, 4]
insertAt('X', 'abcd', 2) // aXbcd
*/

const insertAt = (props) => {
    const {
        element, ary, insertIndex
    } = props;

    switch(typeof ary){
        case 'string':
            return insert_string(element, [...ary], insertIndex);
        case 'object':
            return insert_object(element, [...ary], insertIndex);
    }
}

const insert_object = (element, ary, insertIndex) => {

    frontAry = ary.slice(0, insertIndex - 1);
    backAry = ary.slice(insertIndex - 1);
    frontAry.push(element);

    return frontAry.concat(backAry);
};

const insert_string = (element, ary, insertIndex) => {

    frontAry = ary.slice(0, insertIndex - 1);
    backAry = ary.slice(insertIndex - 1);
    frontAry.push(element);

    return frontAry.concat(backAry).join('');
};


/*-------テスト-------*/ 
const test1 = {element: 5, ary: [1, 2, 3, 4], insertIndex: 3};
const test2 = {element: 'X', ary: 'abcd', insertIndex: 2};

console.log(insertAt(test1));
console.log(insertAt(test2));