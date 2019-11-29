// https://habr.com/ru/post/130472/

// Option 1

const int8 PNG_SIGNATURE = 0x89;
const int8 DOS_EOF = 0x1A;
const int CR = 0x0A;
const int LF = 0x0D;
const int16 CRLF = 0x0A0D;

struct ChunkChar<prop> {
	@ = char {
		@ = @;
		when @ ~ /[A-Z]/ {
			@.<prop> = true;
		}
		when @ ~ /[a-z]/ {
			@.<prop> = false;
		}
	}

	@.<prop>: bool;
}

struct FourthChunkChar : ChunkChar<copyable> {}

struct PNGChunkType {
	.first: ChunkChar<critical> {
		@.critical = critical;
	};
	.second: ChunkChar<public> {
		@.public = public;
	};
	.third: ChunkChar<diff> {
		@.diff = diff;
	};
	.fourth: FourthChunkChar {
		@.copyable = copyable;
	};

	// Getter
	// get name() {
	// 	= @first
	// 	. @second
	// 	. @third
	// 	. @fourth
	// };

	// Method;
	// method() {}

	// Properties
	@.critical: bool;
	@.public: bool;
	@.copyable: bool;
	@.diff: bool;
}


struct PNGChunk {
	.length: int32;
	.type: PNGChunkType;
	.content: int8[@.length];
	.crc: int32;
}

struct IHDRChunk : PNGChunk {
	override .type: PNGChunkType {
		when (@.name != "ihdr") {
			throw "It's not an IHDR chunk: ";
		}
	};
	
	.width: int32;
	.height: int32;
	.bitDepth: int8;
	.colorType: int8 {
		.usePalette = false;
		when (@ ~~ 1) { // Прочитать первый бит
			@.usePalette = true;
		}

		// Либо так
		@.useColor = (@ ~~ 1);
		@.hasAlpha = (@ ~~ 1);
	};
	.compressionMethod: int8;
	.filterMethod: int8;
	.interlace: bool;

	@.usePalette: bool;
	@.useColor: bool;
	@.hasAlpha: bool;
}

struct IENDChunk : PNGChunk {
	override .type: PNGChunkType {
		when (@.name != "ihdr") {
			throw "It's not an IEND chunk";
		}
	};
}

struct V1 {
	@ = int8 {
		when (@ == PNG_SIGNATURE) {
			@ = int8[3] {
				when (@ == "PNG") {
					@ = int16 {
						when @ = CRLF {
							@ = int8 {
								when @ = DOS_EOF {
									@ = int8 {
										when @ = LF {
											@ = IHDRChunk {
												.header: @;
												.chunks[]: PNGChunk {};
												.end: IENDChunk;
											}
										}
										default { // Not LF
											throw "Wrong PNG signature";
										}
									}
								}
								
								default { // Not DOS_EOF
									throw "Wrong PNG signature";
								}
							}
						}

						default { // Not CRLF
							throw "Wrong PNG signature";
						}
					}
				}

				default { // not PNG
					throw "Wrong PNG signature";
				}
			}
		}
		default { // not PNG_SIGNATURE
			throw "Wrong PNG signature";
		}
	}
}
