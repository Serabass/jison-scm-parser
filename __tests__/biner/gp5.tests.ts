import * as fs from "fs";
import { load } from "../../src/util";

describe("Generic", () => {
  it("Test gp5", () => {
    let b = fs.readFileSync("./2chords.gp5");
    let pr = load("gp5", b);
    let result = pr.run();

  });
});
