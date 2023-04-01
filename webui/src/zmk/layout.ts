export interface Zmk {
  layout: ZmkLayoutKey[];
  symbols: SymbolMap;
}

export interface SymbolMap {
  [key: string]: string;
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
