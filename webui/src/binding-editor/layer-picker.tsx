import { Button, Menu, MenuItem } from "@mui/material";
import { useState, SyntheticEvent, Dispatch, SetStateAction } from "react";
import { KeymapBinding, KeymapLayer } from "../keymap";
import { MouseEvent } from "react";

const LayerPicker = ({
  layers,
  binding,
  updateBinding,
}: {
  layers: KeymapLayer[];
  binding: KeymapBinding;
  updateBinding: Dispatch<SetStateAction<KeymapBinding>>;
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
    updateBinding({ ...binding, first: [String(i)] });
  };

  const layerIndex = binding.first ? Number(binding.first[0]) : 0;

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
