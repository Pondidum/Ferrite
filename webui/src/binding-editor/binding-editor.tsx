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
import { Keymap, Binding, Parameter } from "../keymap";
import LayerPicker from "./layer-picker";
import KeyPicker from "./key-picker";

const paramOrDefault = (params: Parameter[], index: number): Parameter =>
  params.length > index ? params[index] : { keyCodes: [] };

const selectEditor = (
  keymap: Keymap,
  binding: Binding,
  updateBinding: Dispatch<SetStateAction<Binding>>
) => {
  switch (binding.action) {
    case "kp":
      return (
        <KeyPicker
          param={paramOrDefault(binding.params, 0)}
          update={(p) =>
            updateBinding((b) => ({
              action: b.action,
              params: [p],
            }))
          }
        />
      );

    case "lt":
      return (
        <>
          <KeyPicker
            param={paramOrDefault(binding.params, 1)}
            update={(p) =>
              updateBinding((b) => ({
                ...b,
                params: [paramOrDefault(b.params, 0), p],
              }))
            }
          />
          <LayerPicker
            layers={keymap.layers}
            param={paramOrDefault(binding.params, 0)}
            update={(p) =>
              updateBinding((b) => ({
                action: b.action,
                params: [p, paramOrDefault(b.params, 1)],
              }))
            }
          />
        </>
      );

    case "mo":
      return (
        <LayerPicker
          layers={keymap.layers}
          param={paramOrDefault(binding.params, 0)}
          update={(p) =>
            updateBinding((b) => ({
              action: b.action,
              params: [p],
            }))
          }
        />
      );

    case "mt":
      return (
        <>
          <KeyPicker
            param={paramOrDefault(binding.params, 1)}
            update={(p) =>
              updateBinding((b) => ({
                action: b.action,
                params: [p, paramOrDefault(b.params, 0)],
              }))
            }
          />
          <KeyPicker
            param={paramOrDefault(binding.params, 0)}
            update={(p) =>
              updateBinding((b) => ({
                action: b.action,
                params: [p, paramOrDefault(b.params, 1)],
              }))
            }
          />
        </>
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
  binding: Binding | undefined;
  onCancel: () => void;
  onConfirm: (newBinding: Binding) => void;
}) => {
  if (!binding) {
    return <></>;
  }

  const [newBinding, setBinding] = useState(binding);

  const updateBinding: Dispatch<SetStateAction<Binding>> = (b) => {
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
