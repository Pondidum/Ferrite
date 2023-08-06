import LayerPicker from "../layer-picker";
import { EditorProps, paramOrDefault } from "./util";

const EditorKP = ({ keymap, options, selected, setOptions }: EditorProps) => {
  const params = options[selected] ?? [];

  return (
    <LayerPicker
      layers={keymap.layers}
      param={paramOrDefault(params, 0)}
      setParam={(p) => {
        setOptions({
          ...options,
          [selected]: [p],
        });
      }}
    />
  );
};

export default EditorKP;
