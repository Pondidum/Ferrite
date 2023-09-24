import { Grid } from "@mui/material";
import KeyPicker from "../key-picker";
import { EditorProps, paramOrDefault } from "./util";
import LayerPicker from "../layer-picker";

const EditorMT = ({ keymap, params, setParams }: EditorProps) => {
  return (
    <>
      <KeyPicker
        param={paramOrDefault(params, 1)}
        setParam={(p) => setParams([paramOrDefault(params, 0), p])}
      />
      <h3>When held press</h3>

      <KeyPicker
        param={paramOrDefault(params, 0)}
        setParam={(p) => setParams([p, paramOrDefault(params, 1)])}
      />
    </>
  );
};

export default EditorMT;
