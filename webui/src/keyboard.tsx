import { Dispatch, SetStateAction } from "react";
import { KeymapBinding, KeymapLayer } from "./keymap";
import { ZmkLayoutKey } from "./zmk";
import Key from "./key";

interface KeyboardProps {
  layout: ZmkLayoutKey[];
  layer: KeymapLayer;
  editBinding: Dispatch<SetStateAction<KeymapBinding | undefined>>;
}

const Keyboard = ({ layout, layer, editBinding }: KeyboardProps) => {
  return (
    <div
      style={{
        position: "relative",
        width: "975px",
        height: "301.476",
        padding: "40px",
      }}
    >
      {layout.map((key, i) => (
        <Key
          key={key.Label}
          zmkKey={key}
          binding={layer.bindings[i]}
          editBinding={editBinding}
        />
      ))}
    </div>
  );
};
export default Keyboard;
