// export interface KeymapConfig {
//   device: Device;
// }

export interface Device {
  keymap: Keymap;
}

export interface Keymap {
  layers: Layer[];
}

export interface Layer {
  name: string;
  bindings: Behavior[];
}

export interface Behavior {
  action: string;
  params?: Param[];
}

export interface Param {
  number?: number;
  keyCode?: string;
}
