export interface Zmk {
  layout: ZmkLayoutKey[];
  keys: { [key: string]: ZmkKey };
}

export interface ZmkLayoutKey {
  Label: string;
  Row: number;
  Col: number;
  X: number;
  Y: number;
  R: number;
  Rx: number;
  Ry: number;
}

export interface ZmkKey {
  names: string[];
  symbol: string;
  description: string;
  context: string;
  clarify: boolean;
  documentation: string;
  os: Os;
  footnotes: { [key: string]: string[] };
}

export interface Os {
  windows: boolean;
  linux: boolean;
  android: boolean;
  macos: boolean;
  ios: boolean;
}
