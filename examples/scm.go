#endianness BE;

/*

| OPCODE | int |   val1  | float32 |      val2    |
|  0005  |  04 |    01   |   06    |  00 00 00 00 |

| OPCODE | var |   val1  | float32 |      val2    |
|  0005  |  02 |  01 02  |   06    |  00 00 00 00 |

| OPCODE |         string          |
|  034A  | 31 31 31 31 00 CC CC CC |

| OPCODE |
|  0009  |

| OPCODE |  var  |  val  |
|  0019  | 02 01 | 04 04 |

*/

scalar decimalValue {
  if ($$ === 0x04) {
    = int8;
  }

  if ($$ === 0x02) {
    = int16;
  }

  if ($$ === 0x06) {
    = float32;
  }
}

struct opcode0005 {
  val1: decimalValue;
  val2: decimalValue;
}

struct opcode034A {
  name: string8;
}

struct opcode0019 {
  val1: decimalValue;
  val2: decimalValue;
}

struct {
  opcode: int16 { opc ->
    if (opc === 0x0005) {
      = opcode0005;
    }

    if (opc === 0x034A) {
      = opcode034A;
    }

    if (opc == 0x0009) {
      type = "opcode0009";
    }

    if (opc === 0x0019) {
      = opcode0019;
    }
  }
}
