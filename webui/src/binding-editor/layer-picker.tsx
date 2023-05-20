import { Button, Menu, MenuItem } from "@mui/material";
import { useState, SyntheticEvent, Dispatch, SetStateAction } from "react";
import { Behavior, Layer, Param } from "../keymap";
import { MouseEvent } from "react";

const LayerPicker = ({
  layers,
  binding,
  updateBinding,
}: {
  layers: Layer[];
  binding: Behavior;
  updateBinding: Dispatch<SetStateAction<Behavior>>;
}) => {
  const [anchorEl, setAnchorEl] = useState<null | HTMLElement>(null);
  const open = Boolean(anchorEl);

  const handleClick = (event: MouseEvent<HTMLButtonElement>) => {
    event.preventDefault();
    setAnchorEl(event.currentTarget);
  };

  const handleClose = () => {
    setAnchorEl(null);
  };

  const handleSelect = (e: SyntheticEvent, i: number) => {
    setAnchorEl(null);

    const [_, ...rest] = binding.params;

    const newSelection: Param = { number: i };

    updateBinding({ ...binding, params: [newSelection, ...rest] });
  };

  const layerIndex =
    (binding.params.length > 0 && binding.params[0].number) || 0;

  return (
    <>
      <Button variant="outlined" onClick={handleClick}>
        {layers[layerIndex].name}
      </Button>
      <Menu anchorEl={anchorEl} open={open} onClose={handleClose}>
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
