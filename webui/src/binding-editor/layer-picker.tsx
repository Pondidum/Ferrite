import { Button, Menu, MenuItem } from "@mui/material";
import { useState, SyntheticEvent, Dispatch, SetStateAction } from "react";
import { Layer, Parameter } from "../keymap";
import { MouseEvent } from "react";

const LayerPicker = ({
  layers,
  param,
  update,
}: {
  layers: Layer[];
  param: Parameter;
  update: (param: Parameter) => void;
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
    update({ number: i, keyCodes: [] });
  };

  return (
    <>
      <Button variant="outlined" onClick={handleClick}>
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
