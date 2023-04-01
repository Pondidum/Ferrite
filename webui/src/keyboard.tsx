import { Dispatch, SetStateAction } from "react";
import { KeymapBinding, KeymapLayer } from "./keymap";
import { Zmk, ZmkLayoutKey } from "./zmk";
import Key from "./key";

interface KeyboardProps {
  zmk: Zmk;
  layer: KeymapLayer;
  editBinding: Dispatch<SetStateAction<KeymapBinding | undefined>>;
}

const Keyboard = ({ zmk, layer, editBinding }: KeyboardProps) => {
  return (
    <div
      style={{
        position: "relative",
        width: "975px",
        height: "301.476",
        padding: "40px",
      }}
    >
      {zmk.layout.map((key, i) => (
        <Key
          key={key.Label}
          zmkKey={key}
          symbols={zmk.symbols}
          binding={layer.bindings[i]}
          editBinding={editBinding}
        />
      ))}
    </div>
  );
};
export default Keyboard;
