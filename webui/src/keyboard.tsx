import { Dispatch, SetStateAction, useContext } from "react";
import { Layer, Binding } from "./keymap";
import Key from "./key";
import { ZmkContext } from "./zmk/context";

interface KeyboardProps {
  bindings: Binding[];
  editBinding: Dispatch<SetStateAction<number | undefined>>;
}

const Keyboard = ({ bindings, editBinding }: KeyboardProps) => {
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
          binding={bindings[i]}
          editBinding={() => editBinding(i)}
        />
      ))}
    </div>
  );
};
export default Keyboard;
