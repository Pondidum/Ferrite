export interface Keymap {
  layers: KeymapLayer[];
}

export interface KeymapLayer {
  name: string;
  bindings: KeymapBinding[];
}

export interface KeymapBinding {
  type: string;
  first: string[];
  second: string[];
}
