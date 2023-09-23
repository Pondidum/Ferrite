import KeyPicker from "../key-picker";
import { EditorProps, paramOrDefault } from "./util";

const EditorMO = ({ params, setParams }: EditorProps) => {
  return (
    <KeyPicker
      param={paramOrDefault(params, 0)}
      setParam={(p) => setParams([p])}
    />
  );
};

export default EditorMO;
