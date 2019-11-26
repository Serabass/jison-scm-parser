import {Parser} from 'jison';

export function test() {
// a grammar in JSON
    var grammar = {
        "lex": {
            "rules": [
                ["\\s+", "/* skip whitespace */"],
                ["[a-f0-9]+", "return 'HEX';"]
            ]
        },

        "bnf": {
            "hex_strings": ["hex_strings HEX",
                "HEX"]
        }
    };

// `grammar` can also be a string that uses jison's grammar format
    var parser = new Parser(grammar);

// generate source, ready to be written to disk
    var parserSource = parser.generate();

// you can also use the parser directly from memory

// returns true
    return parser.parse("adfe34bc e82a");
}
