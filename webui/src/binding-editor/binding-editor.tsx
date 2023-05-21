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
import { Dispatch, SetStateAction, SyntheticEvent, useState } from "react";
import { Keymap, Behavior } from "../keymap";
import LayerPicker from "./layer-picker";

const selectEditor = (
  keymap: Keymap,
  binding: Behavior,
  updateBinding: Dispatch<SetStateAction<Behavior>>
) => {
  switch (binding.action) {
    case "lt":
      return (
        <LayerPicker
          layers={keymap.layers}
          binding={binding}
          updateBinding={updateBinding}
        />
      );

    default:
      return <></>;
  }
};

const BindingEditor = ({
  open,
  keymap,
  binding,
  onCancel,
  onConfirm,
}: {
  open: boolean;
  keymap: Keymap;
  binding: Behavior | undefined;
  onCancel: () => void;
  onConfirm: (newBinding: Behavior) => void;
}) => {
  if (!binding) {
    return <></>;
  }

  const [newBinding, setBinding] = useState(binding);

  const updateBinding: Dispatch<SetStateAction<Behavior>> = (b) => {
    console.log("new binding:", b);
    setBinding(b);
  };

  const selectTab = (e: SyntheticEvent, newValue: string) => {
    updateBinding({ ...newBinding, action: newValue });
  };

  const editor = selectEditor(keymap, newBinding, updateBinding);

  return (
    <Dialog open={open} onClose={onCancel}>
      <DialogTitle>Configure Key</DialogTitle>
      <Box>
        <Tabs value={newBinding.action} onChange={selectTab}>
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
        <Button onClick={() => onConfirm(newBinding)}>Confirm</Button>
      </DialogActions>
    </Dialog>
  );
};

export default BindingEditor;
