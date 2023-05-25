import {
  CSSProperties,
  Dispatch,
  SetStateAction,
  SyntheticEvent,
  useContext,
} from "react";
import "./key.css";
import { Behavior } from "./keymap";
import { ZmkKey, ZmkLayoutKey } from "./zmk";
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
  binding: Behavior;
  editBinding: Dispatch<SetStateAction<Behavior | undefined>>;
}

export const keysFromCombo = (input: string | undefined) => {
  if (!input) {
    return [];
  }

  const keys = [];
  let current = "";
  for (let i = 0; i < input.length; i++) {
    const char = input[i];

    if (char === "(") {
      keys.push(current);
      current = "";
    } else if (char === ")") {
      break;
    } else {
      current += char;
    }
  }

  if (current !== "") {
    keys.push(current);
  }

  return keys;
};

const trySymbol = (lookup: { [key: string]: ZmkKey }, key: string) =>
  lookup[key] ? lookup[key].symbol || key : key;

const render = (lookup: { [key: string]: ZmkKey }, combo: string | undefined) =>
  keysFromCombo(combo)
    .map((k) => trySymbol(lookup, k))
    .join(" ");

const Key = ({ zmkKey, binding, editBinding }: KeyProps) => {
  const zmk = useContext(ZmkContext);

  const onClick = (e: SyntheticEvent) => {
    editBinding(binding);
  };

  const params = binding.params || [];

  const first =
    params.length > 0 &&
    (params[0].number || render(zmk.keys, params[0].keyCode));
  const second =
    params.length > 1 &&
    (params[1].number || render(zmk.keys, params[1].keyCode));

  return (
    <div className="key" style={styleKey(zmkKey)} onClick={onClick}>
      <span className="behaviour-binding">{binding.action}</span>
      {second && <div className="code">{second}</div>}
      <div className="code">{first}</div>
    </div>
  );
};

export default Key;
