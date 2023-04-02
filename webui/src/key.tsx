import {
  CSSProperties,
  Dispatch,
  SetStateAction,
  SyntheticEvent,
  useContext,
} from "react";
import "./key.css";
import { KeymapBinding } from "./keymap";
import { ZmkLayoutKey } from "./zmk";
import { ZmkContext } from "./zmk/context";

const DefaultSize = 65;
const DefaultPadding = 5;

const styleKey = (k: ZmkLayoutKey): CSSProperties => {
  const x = k.X * (DefaultSize + DefaultPadding);
  const y = k.Y * (DefaultSize + DefaultPadding);
  const w = DefaultSize;
  const h = DefaultSize;
  const rx = (k.X - Math.max(k.Rx, k.X)) * -(DefaultSize + DefaultPadding);
  const ry = (k.Y - Math.max(k.Ry, k.Y)) * -(DefaultSize + DefaultPadding);
  const a = k.R;

  return {
    top: `${y}px`,
    left: `${x}px`,
    width: `${w}px`,
    height: `${h}px`,
    transformOrigin: `${rx}px ${ry}px`,
    transform: `rotate(${a}deg)`,
  };
};

interface KeyProps {
  zmkKey: ZmkLayoutKey;
  binding: KeymapBinding;
  editBinding: Dispatch<SetStateAction<KeymapBinding | undefined>>;
}

const Key = ({ zmkKey, binding, editBinding }: KeyProps) => {
  const zmk = useContext(ZmkContext);

  const onClick = (e: SyntheticEvent) => {
    editBinding(binding);
  };

  const first = binding.first?.map((b) => zmk.keys[b]?.symbol || b).join(" ");
  const second = binding.second?.map((b) => zmk.keys[b]?.symbol || b).join(" ");

  return (
    <div className="key" style={styleKey(zmkKey)} onClick={onClick}>
      <span className="behaviour-binding">{binding.type}</span>
      {second && <div className="code">{second}</div>}
      <div className="code">{first}</div>
    </div>
  );
};

export default Key;
