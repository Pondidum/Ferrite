import { describe, it, expect } from "vitest";
import { keysFromCombo } from "./key";

describe("should match", () => {
  it.each([
    ["H", ["H"]],
    ["LS(H)", ["LS(code)", "H"]],
    ["LGUI(LS(H))", ["LGUI(code)", "LS(code)", "H"]],
  ])("%s", (input, expected) => {
    expect(keysFromCombo(input)).toEqual(expected);
  });
});
