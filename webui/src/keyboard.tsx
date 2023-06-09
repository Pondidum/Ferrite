import { Dispatch, SetStateAction, useContext } from "react";
import { Layer, Binding } from "./keymap";
import Key from "./key";
import { ZmkContext } from "./zmk/context";

interface KeyboardProps {
  layer: Layer;
  editBinding: Dispatch<SetStateAction<Binding | undefined>>;
}

const Keyboard = ({ layer, editBinding }: KeyboardProps) => {
  const zmk = useContext(ZmkContext);
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
          binding={layer.bindings[i]}
          editBinding={editBinding}
        />
      ))}
    </div>
  );
};
export default Keyboard;
