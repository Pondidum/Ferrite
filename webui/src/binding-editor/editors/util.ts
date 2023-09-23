import { Keymap, Parameter } from "../../keymap";

export const paramOrDefault = (params: Parameter[], index: number): Parameter =>
  params.length > index ? params[index] : {};

export interface EditorProps {
  keymap: Keymap;
  params: Parameter[];
  setParams: (params: Parameter[]) => void;
}
