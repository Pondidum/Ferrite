import { Grid } from "@mui/material";
import KeyPicker from "../key-picker";
import { EditorProps, paramOrDefault } from "./util";
import LayerPicker from "../layer-picker";

const EditorLT = ({ keymap, options, selected, setOptions }: EditorProps) => {
  const params = options[selected] ?? [];

  return (
    <>
      <KeyPicker
        param={paramOrDefault(params, 1)}
        setParam={(p) => {
          setOptions({
            ...options,
            [selected]: [paramOrDefault(params, 0), p],
          });
        }}
      />
      <h3>When held, switch to layer</h3>

      <Grid container spacing={1} columns={4}>
        <Grid item xs={2}>
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
        </Grid>
      </Grid>
    </>
  );
};

export default EditorLT;
