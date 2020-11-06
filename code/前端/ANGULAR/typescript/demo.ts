var b: boolean;     // 显式指定类型
var yes = true;     // 等同于 yes: boolean = true
var no = false;     // 等同于 no: boolean = false

console.log('b type is: ' + typeof(b));   // b type is undefined
console.log('yes type is: ' + typeof(yes));
console.log('no type is: ' + typeof(no));

// 赋值
b = false;
console.log('b type is: ' + typeof(b));