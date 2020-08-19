/*
指定された回数だけ配列や文字列の要素を複製するrepli関数を実装せよ。
repli([1, 2, 3], 3) // [1, 1, 1, 2, 2, 2, 3, 3, 3]
repli('abc', 3) // aaabbbccc
*/

const repli = (props) => {
    switch(typeof props.ary){
        case 'string':
            props.ary = [...props.ary];
            return exe(props).join('');
        case 'object':
            return exe(props);   
    }
}

const exe = (props) => {
    let repliAry = [];

    props.ary.forEach(element => {
        repliElement = Array(props.count).fill(element);
        repliAry.push(...repliElement);
    })
    return repliAry;
};

/*-------テスト-------*/ 
const test1 = {ary: [1, 2, 3], count: 3};
const test2 = {ary: 'abc', count: 3};

console.log(repli(test1));
console.log(repli(test2));