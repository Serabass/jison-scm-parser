// Структуры не могут ссылаться на самих себя, а также не может быть вложенных рекурсий

struct A {
	a: uint8;
	b: A;
}

struct B {
	c: C;
	u: uint8;
}

struct C {
	a: A;
}