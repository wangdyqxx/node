## go slice

slice结构主要包括3个部分：data、len、cap。结构中的len和cap分别指示元素数量和容量。append使得len变大，但如果append之后的元素个数大于cap，会引发扩容机制。此时，会重新创建一个容量适合的底层数组。data为指针，指向底层数组。

## go slice 能否值传递

* 1、go中函数传递都是值传递
* 2、slice、map、channel都是引用类型，即便是值传递，结构内部还是指向原来的引用对象，所以函数体内可以直接修改元素。
* 3、如果slice触发扩容，data会指向新的底层数组，而不指向外部的底层数组了。所以之后再修改slice，不会对外部的slice造成影响。


