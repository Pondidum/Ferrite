import { Grid, ToggleButton } from "@mui/material";
import { Parameter } from "../keymap";
import { useContext } from "react";
import { ZmkContext } from "../zmk/context";
import "./modifier.css";

const ModifierPicker = ({ param }: { param: Parameter }) => {
  return (
    <>
      <h3>When held, press</h3>
      <ModifierGrid selected={param.modifiers ?? []} />
    </>
  );
};

export const ModifierGrid = ({ selected }: { selected: string[] }) => {
  const zmk = useContext(ZmkContext);
  const modifiers = Object.entries(zmk.keys)
    .filter(([name]) => name.endsWith("(code)"))
    .map(([_, key]) => key);

  return (
    <Grid container spacing={1} columns={4}>
      {modifiers.map((key, i) => (
        <Grid item key={i} xs={1}>
          <ToggleButton
            className="modifier"
            value={key.names[0]}
            selected={key.names.some((name) => selected.includes(name))}
            fullWidth
          >
            {key.names[0]}
          </ToggleButton>
        </Grid>
      ))}
    </Grid>
  );
};

export default ModifierPicker;
