import { Autocomplete, Grid, TextField } from "@mui/material";
import { Parameter } from "../keymap";
import { ModifierGrid } from "./modifier-grid";
import { useContext } from "react";
import { ZmkContext } from "../zmk/context";

const KeyPicker = ({
  param,
  setParam,
}: {
  param: Parameter;
  setParam: (param: Parameter) => void;
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
            onChange={(x, v) => {
              setParam({
                ...param,
                keyCode: v ?? undefined,
              });
            }}
          />
        </Grid>
      </Grid>

      <h4>With Modifiers</h4>
      <ModifierGrid
        modifiers={param.modifiers ?? []}
        setModifiers={(m) => {
          setParam({
            ...param,
            modifiers: m,
          });
        }}
      />
    </>
  );
};

export default KeyPicker;
