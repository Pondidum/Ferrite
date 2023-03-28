import { CSSProperties } from "react";
import { KeymapBinding, KeymapLayer, ZmkLayoutKey } from "./App";
import "./key.css";

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

interface KeyboardProps {
  layout: ZmkLayoutKey[];
  layer: KeymapLayer;
}

interface KeyProps {
  zmkKey: ZmkLayoutKey;
  binding: KeymapBinding;
}

const Key = ({ zmkKey, binding }: KeyProps) => (
  <div className="key" style={styleKey(zmkKey)}>
    <span className="behaviour-binding">{binding.type}</span>
    <span className="code">{binding.first.join(" ")}</span>
  </div>
);

const Keyboard = ({ layout, layer }: KeyboardProps) => {
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
        <Key key={key.Label} zmkKey={key} binding={layer.bindings[i]} />
      ))}
    </div>
  );
};
export default Keyboard;
