import {
  Box,
  Tabs,
  Tab,
  DialogTitle,
  DialogActions,
  Button,
  Dialog,
  Menu,
  MenuItem,
} from "@mui/material";
import { SyntheticEvent, useState } from "react";
import { Keymap, KeymapBinding, KeymapLayer } from "../App";
import { MouseEvent } from "react";

const LayerPicker = ({
  layers,
  selected,
}: {
  layers: KeymapLayer[];
  selected: string;
}) => {
  const [selectedLayer, setSelectedLayer] = useState<number>(Number(selected));

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
    setSelectedLayer(i);
  };

  return (
    <>
      <Button variant="outlined" onClick={handleClick}>
        {layers[selectedLayer].name}
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

const selectEditor = (
  keymap: Keymap,
  action: string,
  binding: KeymapBinding
) => {
  switch (action) {
    case "kp":
      return (
        <p>
          When tapped <a href="">{binding.first}</a>
        </p>
      );

    case "mt":
      return (
        <>
          <p>
            When tapped <a href="">{binding.first}</a>
          </p>
          <p>
            When held <a href="">{binding.second}</a>
          </p>
        </>
      );

    case "lt":
      return (
        <>
          <p>
            When tapped <a href="">{binding.second}</a>
          </p>
          <p>
            When held switch to layer{" "}
            <LayerPicker layers={keymap.layers} selected={binding.first[0]} />.
          </p>
        </>
      );

    case "mo":
      return (
        <p>
          When tapped switch to layer <a href="">{binding.first}</a>
        </p>
      );

    case "none":
      return <p>Do nothing</p>;

    case "trans":
      return (
        <p>Pass the keypress through to the next layer below in the stack</p>
      );

    default:
      return <></>;
  }
};

const KeyEditor = ({
  open,
  keymap,
  binding,
  onCancel,
  onConfirm,
}: {
  open: boolean;
  keymap: Keymap;
  binding: KeymapBinding | undefined;
  onCancel: () => void;
  onConfirm: (newBinding: KeymapBinding) => void;
}) => {
  if (!binding) {
    return <></>;
  }

  const [tab, setTab] = useState(binding.type);

  const selectTab = (e: SyntheticEvent, newValue: string) => {
    setTab(newValue);
  };

  const editor = selectEditor(keymap, tab, binding);

  return (
    <Dialog open={open}>
      <DialogTitle>Configure Key</DialogTitle>
      <Box>
        <Tabs value={tab} onChange={selectTab}>
          <Tab value={"kp"} label="KP" />
          <Tab value={"mt"} label="MT" />
          <Tab value={"lt"} label="LT" />
          <Tab value={"mo"} label="MO" />
          <Tab value={"none"} label="None" />
          <Tab value={"trans"} label="Trans" />
        </Tabs>
      </Box>
      {editor}

      <DialogActions>
        <Button onClick={onCancel}>Cancel</Button>
        <Button onClick={() => onConfirm(binding)}>Confirm</Button>
      </DialogActions>
    </Dialog>
  );
};

export default KeyEditor;
