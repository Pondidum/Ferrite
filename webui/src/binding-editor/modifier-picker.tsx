import {
  Button,
  ButtonGroup,
  Container,
  Grid,
  TextField,
  ToggleButton,
  ToggleButtonGroup,
} from "@mui/material";
import { Param } from "../keymap";
import { useContext, useMemo } from "react";
import { ZmkContext } from "../zmk/context";

const ModifierPicker = ({
  param,
  update,
}: {
  param: Param;
  update: (param: Param) => void;
}) => {
  const zmk = useContext(ZmkContext);
  const modifiers = Object.entries(zmk.keys)
    .filter(([name]) => name.endsWith("(code)"))
    .map(([_, key]) => key);

  return (
    <>
      <h3>When held, press</h3>

      <Grid container spacing={1} columns={2}>
        {modifiers.map((key, i) => (
          <Grid item key={i} xs={1}>
            <ToggleButton value={key.names[0]} fullWidth>
              {key.names[0]}
            </ToggleButton>
          </Grid>
        ))}
      </Grid>
    </>
  );
};

export default ModifierPicker;
