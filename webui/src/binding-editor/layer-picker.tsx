import { Button, Menu, MenuItem } from "@mui/material";
import { useState, SyntheticEvent, Dispatch, SetStateAction } from "react";
import { KeymapBinding, KeymapLayer } from "../App";
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

  return (
    <>
      <Button variant="outlined" onClick={handleClick}>
        {layers[Number(binding.first[0])].name}
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
