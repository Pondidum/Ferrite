import { replace } from "./device";
import { describe, expect, test } from "vitest";

describe("replacing arrays", () => {
  test("negative index", () => {
    const source = ["a", "b", "c", "d", "e"];
    const result = replace(source, -2, "replaced");

    expect(result).toEqual(source);
  });

  test("out of bounds", () => {
    const source = ["a", "b", "c", "d", "e"];
    const result = replace(source, 17, "replaced");

    expect(result).toEqual(source);
  });

  test("middle", () => {
    const source = ["a", "b", "c", "d", "e"];
    const result = replace(source, 2, "replaced");

    expect(result).toEqual(["a", "b", "replaced", "d", "e"]);
    expect(result).not.toBe(source);
  });

  test("first", () => {
    const source = ["a", "b", "c", "d", "e"];
    const result = replace(source, 0, "replaced");

    expect(result).toEqual(["replaced", "b", "c", "d", "e"]);
    expect(result).not.toBe(source);
  });

  test("last", () => {
    const source = ["a", "b", "c", "d", "e"];
    const result = replace(source, 4, "replaced");

    expect(result).toEqual(["a", "b", "c", "d", "replaced"]);
    expect(result).not.toBe(source);
  });
});
