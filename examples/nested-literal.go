struct Parent {
	a: uint8;

	data: struct {
		title: char[10];
		meta: struct Meta {
			sandbox: uint[255];
		}
	}

	endbyte: uint8;
}