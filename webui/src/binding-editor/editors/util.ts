import { Actions, Parameter } from "../../keymap";
import { Options } from "../binding-editor";

export const paramOrDefault = (params: Parameter[], index: number): Parameter =>
  params.length > index ? params[index] : {};

export interface EditorProps {
  keymap: Keymap;
  selected: Actions;
  options: Options;
  setOptions: (options: Options) => void;
}
