/*
配列や文字列のn番目の要素を削除するremoveAt関数を実装せよ。
削除した要素と処理後の配列を配列に格納し、呼び出した結果として返しなさい。
removeAt(3, [1, 2, 3]) // [3, [1, 2]]
removeAt(2, 'abcd') // ['b', 'acd']
*/

const removeAt = (props) => {
    const {
        removeIndex, ary,
    } = props;

    switch(typeof ary){
        case 'string':
            return string_remove(removeIndex, [...ary]);
        case 'object':
            return object_remove(removeIndex, ary);
    }
}

const object_remove = (removeIndex, ary) => {
    removedAry = [];
    removedElement = [];

    ary.forEach((element, index) => {
        if(index === removeIndex - 1){
            removedElement.push(element);
        }else{
            removedAry.push(element);
        }
    });
    removedElement.push(removedAry);
    return removedElement;
};

const string_remove = (removeIndex, ary) => {
    removedAry = [];
    removedElement = [];

    ary.forEach((element, index) => {
        if(index === removeIndex - 1){
            removedElement.push(element);
        }else{
            removedAry.push(element);
        }
    });
    removedElement.push(removedAry.join(''));
    return removedElement;
};

/*-------テスト-------*/ 
const test1 = {removeIndex: 3, ary: [1, 2, 3]};
const test2 = {removeIndex: 2, ary: 'abcd'};

console.log(removeAt(test1));
console.log(removeAt(test2));