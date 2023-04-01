export interface Zmk {
  layout: ZmkLayoutKey[];
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
