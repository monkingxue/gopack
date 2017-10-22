var foo = require('bar');
var foo = 2
var Person = function (name, age) {
    //共有属性
    this.name = name;
    //共有方法
    this.sayName = function () {
        console.log(this.name);
    };
    //静态私有属性(只能用于内部调用)
    var home = "China";

    //静态私有方法(只能用于内部调用)
    function sayHome() {
        console.log(home);
    }

    //构造器
    this.setAge = function (age) {
        console.log(age + 12);
    };
    this.setAge(age);
}
//静态方法（只能被类来访问）
Person.sayAge = function () {
    console.log("your age is 12");
}
//静态属性（只能被类来访问）
Person.drink = "water";
//静态共有方法（类和实例都可以访问）
Person.prototype.sayWord = function () {
    console.log("ys is a boy");
}