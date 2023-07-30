import { Autocomplete, Grid, TextField } from "@mui/material";
import { Parameter } from "../keymap";
import { ModifierGrid } from "./modifier-picker";
import { useContext } from "react";
import { ZmkContext } from "../zmk/context";

const KeyPicker = ({
  param,
  update,
}: {
  param: Parameter;
  update: (param: Parameter) => void;
}) => {
  const zmk = useContext(ZmkContext);
  const keys = [...new Set(Object.values(zmk.keys).flatMap((k) => k.names))];

  return (
    <>
      <h3>When tapped, press</h3>

      <Grid container spacing={1} columns={4}>
        <Grid item xs={2}>
          <Autocomplete
            disablePortal
            options={keys}
            renderInput={(p) => <TextField {...p} label="Key" />}
            value={param.keyCode}
          />
        </Grid>
      </Grid>

      <h4>With Modifiers</h4>
      <ModifierGrid selected={param.modifiers ?? []} />
    </>
  );
};

export default KeyPicker;
