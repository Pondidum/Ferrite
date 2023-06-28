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

export interface Binding {
  action: string;
  params: Parameter[];
}

export interface Parameter {
  number?: number;
  keyCodes: string[];
}

export interface Layer {
  name: string;
  bindings: Binding[];
}
