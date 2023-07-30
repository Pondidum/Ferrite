export interface Keymap {
  configs: Configuration[];
  combos: Combo[];
  layers: Layer[];
}

export interface Configuration {
  name: string;
  properties: { [key: string]: Value };
}

export interface Value {
  string?: string;
  number?: number;
}

export interface Combo {
  name: string;
  timeout: number;
  keyPositions: number[];
  layers: number[];
  bindings: Binding[];
}

export type Actions = "kp" | "mt" | "lt" | "mo" | "none" | "trans";

export interface Binding {
  action: Actions;
  params: Parameter[];
}

export interface Parameter {
  number?: number;
  keyCode?: string;
  modifiers?: string[];
}

export interface Layer {
  name: string;
  bindings: Binding[];
}
