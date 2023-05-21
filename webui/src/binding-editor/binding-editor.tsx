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
    // case "kp":
    //   return (
    //     <p>
    //       When tapped <a href="">{binding.params[0].keyCode}</a>
    //     </p>
    //   );

    // case "mt":
    //   return (
    //     <>
    //       <p>
    //         When tapped <a href="">{binding.first}</a>
    //       </p>
    //       <p>
    //         When held <a href="">{binding.second}</a>
    //       </p>
    //     </>
    //   );

    // case "lt":
    //   return (
    //     <>
    //       <p>
    //         When tapped <a href="">{binding.second}</a>
    //       </p>
    //       <p>
    //         When held switch to layer{" "}
    //         <LayerPicker
    //           layers={keymap.layers}
    //           binding={binding}
    //           updateBinding={updateBinding}
    //         />
    //         .
    //       </p>
    //     </>
    //   );

    // case "mo":
    //   return (
    //     <p>
    //       When tapped switch to layer
    //       <LayerPicker
    //         layers={keymap.layers}
    //         binding={binding}
    //         updateBinding={updateBinding}
    //       />
    //       .
    //     </p>
    //   );

    // case "none":
    //   return <p>Do nothing</p>;

    // case "trans":
    //   return (
    //     <p>Pass the keypress through to the next layer below in the stack</p>
    //   );

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
