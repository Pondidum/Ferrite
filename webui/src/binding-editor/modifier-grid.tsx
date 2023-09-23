import { Grid, ToggleButton } from "@mui/material";
import { useContext } from "react";
import { ZmkContext } from "../zmk/context";
import "./modifier.css";

export const ModifierGrid = ({
  modifiers,
  setModifiers,
}: {
  modifiers: string[];
  setModifiers: (m: string[]) => void;
}) => {
  const zmk = useContext(ZmkContext);
  const modifierKeys = Object.entries(zmk.keys)
    .filter(([name]) => name.endsWith("(code)"))
    .map(([_, key]) => key);

  return (
    <Grid container spacing={1} columns={4}>
      {modifierKeys.map((key, i) => (
        <Grid item key={i} xs={1}>
          <ToggleButton
            className="modifier"
            value={key.names[0]}
            selected={modifiers.includes(key.names[0])}
            onChange={(event, value: string) => {
              // console.log(key, value);

              const i = modifiers.indexOf(value);
              if (i === -1) {
                setModifiers([...modifiers, value]);
              } else {
                const without = modifiers.slice();
                without.splice(i, 1);
                setModifiers(without);
              }
            }}
            fullWidth
          >
            {key.names[0]}
          </ToggleButton>
        </Grid>
      ))}
    </Grid>
  );
};

export default ModifierGrid;
