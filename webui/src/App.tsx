import { useEffect, useState } from "react";
import "./App.css";
import Keyboard from "./keyboard";

interface Zmk {
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

function App() {
  const [zmk, setZmk] = useState<Zmk>({ layout: [] });

  const [keymap, setKeymap] = useState<Keymap>({ layers: [] });

  useEffect(() => {
    fetch("http://localhost:5656/api/zmk")
      .then((r) => r.json())
      .then((j) => setZmk(j));
  });

  useEffect(() => {
    fetch("http://localhost:5656/api/keymap")
      .then((r) => r.json())
      .then((j) => setKeymap(j));
  });

  return (
    <div className="App">
      <Keyboard layout={zmk.layout} layer={keymap.layers[0]} />
    </div>
  );
}

export default App;
