import { describe, it, expect } from "vitest";
import { keysFromCombo } from "./key";

describe("should match", () => {
  it.each([
    ["H", "H"],
    ["LS(H)", "LS H"],
    ["LGUI(LS(H))", "LGUI LS H"],
  ])("%s", (input, expected) => {
    expect(keysFromCombo(input)).toEqual(expected);
  });
});
