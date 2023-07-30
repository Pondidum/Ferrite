import { Button, Menu, MenuItem } from "@mui/material";
import { useState, SyntheticEvent, Dispatch, SetStateAction } from "react";
import { Layer, Parameter } from "../keymap";
import { MouseEvent } from "react";

const LayerPicker = ({
  layers,
  param,
  setParam,
}: {
  layers: Layer[];
  param: Parameter;
  setParam: (param: Parameter) => void;
}) => {
  const [anchorEl, setAnchorEl] = useState<null | HTMLElement>(null);

  const handleClick = (event: MouseEvent<HTMLButtonElement>) => {
    event.preventDefault();
    setAnchorEl(event.currentTarget);
  };

  const handleClose = () => {
    setAnchorEl(null);
  };

  const handleSelect = (e: SyntheticEvent, i: number) => {
    setAnchorEl(null);
    setParam({ number: i });
  };

  return (
    <>
      <Button variant="outlined" onClick={handleClick} fullWidth>
        {layers[param.number || 0].name}
      </Button>
      <Menu anchorEl={anchorEl} open={Boolean(anchorEl)} onClose={handleClose}>
        {layers.map((l, i) => (
          <MenuItem key={i} onClick={(e) => handleSelect(e, i)}>
            {l.name}
          </MenuItem>
        ))}
      </Menu>
    </>
  );
};

export default LayerPicker;
