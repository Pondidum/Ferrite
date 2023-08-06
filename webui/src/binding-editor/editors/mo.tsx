import KeyPicker from "../key-picker";
import { EditorProps, paramOrDefault } from "./util";

const EditorMO = ({ options, selected, setOptions }: EditorProps) => {
  const params = options[selected] ?? [];

  return (
    <KeyPicker
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

export default EditorMO;
