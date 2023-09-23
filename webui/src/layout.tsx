import { Dispatch, SetStateAction, useContext } from "react";
import { Layer, Binding } from "./keymap";
import Key from "./key";
import { ZmkContext } from "./zmk/context";

interface KeyboardProps {
  bindings: Binding[];
  startEditing: (target: { key: number; binding: Binding }) => void;
}

const Layout = ({ bindings, startEditing }: KeyboardProps) => {
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
          startEditing={(b) => startEditing({ key: i, binding: b })}
        />
      ))}
    </div>
  );
};
export default Layout;
