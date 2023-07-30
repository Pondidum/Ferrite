import {
  CSSProperties,
  Dispatch,
  SetStateAction,
  SyntheticEvent,
  useContext,
} from "react";
import "./key.css";
import { Binding, Parameter } from "./keymap";
import { ZmkKey, ZmkLayoutKey } from "./zmk";
import { ZmkContext } from "./zmk/context";

const DefaultSize = 65;
const DefaultPadding = 5;

const styleKey = (k: ZmkLayoutKey): CSSProperties => {
  const x = k.Col * (DefaultSize + DefaultPadding);
  const y = k.Y * (DefaultSize + DefaultPadding);
  const w = DefaultSize;
  const h = DefaultSize;
  const rx = (k.Col - Math.max(k.Rx, k.Col)) * -(DefaultSize + DefaultPadding);
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
  binding: Binding;
  editBinding: Dispatch<SetStateAction<Binding | undefined>>;
}

const trySymbol = (lookup: { [key: string]: ZmkKey }, key: string) =>
  lookup[key] ? lookup[key].symbol || key : key;

const render = (lookup: { [key: string]: ZmkKey }, param: Parameter) => {
  const keys = [];

  if (param.modifiers) {
    keys.push(...param.modifiers);
  }

  if (param.keyCode) {
    keys.push(param.keyCode);
  }

  return keys.map((m) => trySymbol(lookup, m)).join(" ");
};

const Key = ({ zmkKey, binding, editBinding }: KeyProps) => {
  const zmk = useContext(ZmkContext);

  const onClick = (e: SyntheticEvent) => {
    editBinding(binding);
  };

  const params = binding.params || [];

  const first =
    params.length > 0 && (params[0].number || render(zmk.keys, params[0]));
  const second =
    params.length > 1 && (params[1].number || render(zmk.keys, params[1]));

  return (
    <div className="key" style={styleKey(zmkKey)} onClick={onClick}>
      <span className="behaviour-binding">{binding.action}</span>
      {second && <div className="code">{second}</div>}
      <div className="code">{first}</div>
    </div>
  );
};

export default Key;
