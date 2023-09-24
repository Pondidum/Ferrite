import {
  Box,
  Tabs,
  Tab,
  DialogTitle,
  DialogActions,
  Button,
  Dialog,
  DialogContent,
} from "@mui/material";
import { SyntheticEvent, useState } from "react";
import { Keymap, Binding, Parameter, Actions } from "../keymap";
import EditorKP from "./editors/kp";
import EditorLT from "./editors/lt";
import EditorMO from "./editors/mo";
import EditorMT from "./editors/mt";

export interface Keybinding {
  key: number;
  binding: Binding;
}

export type Options = { [key: string]: Parameter[] };

const selectEditor = (
  keymap: Keymap,
  selected: Actions,
  options: Options,
  setOptions: (options: Options) => void
) => {
  const params = options[selected] ?? [];

  const setParams = (params: Parameter[]) =>
    setOptions({
      ...options,
      [selected]: params,
    });

  switch (selected) {
    case "kp":
      return <EditorKP keymap={keymap} params={params} setParams={setParams} />;

    case "lt":
      return <EditorLT keymap={keymap} params={params} setParams={setParams} />;

    case "mo":
      return <EditorMO keymap={keymap} params={params} setParams={setParams} />;

    case "mt":
      return <EditorMT keymap={keymap} params={params} setParams={setParams} />;

    default:
      return <></>;
  }
};

const BindingEditor = ({
  open,
  keymap,
  target,
  onCancel,
  onConfirm,
}: {
  open: boolean;
  keymap: Keymap;
  target: Keybinding | undefined;
  onCancel: () => void;
  onConfirm: (newBinding: Keybinding) => void;
}) => {
  if (!target) {
    return <></>;
  }

  const binding = target.binding;

  const [selected, setSelected] = useState<Actions>(binding.action);
  const [options, setOptions] = useState<Options>({
    [binding.action]: binding.params,
  });

  const selectTab = (e: SyntheticEvent, newValue: Actions) => {
    setSelected(newValue);
  };

  const editor = selectEditor(keymap, selected, options, setOptions);

  return (
    <Dialog open={open} onClose={onCancel} maxWidth={"sm"} fullWidth>
      <DialogTitle>Configure Key</DialogTitle>

      <DialogContent>
        <Box>
          <Tabs value={selected} onChange={selectTab}>
            <Tab value={"kp"} label="KP" />
            <Tab value={"mt"} label="MT" />
            <Tab value={"lt"} label="LT" />
            <Tab value={"mo"} label="MO" />
            <Tab value={"none"} label="None" />
            <Tab value={"trans"} label="Trans" />
          </Tabs>
        </Box>

        {editor}
      </DialogContent>

      <DialogActions>
        <Button onClick={onCancel}>Cancel</Button>
        <Button
          onClick={() => {
            const newBinding = { action: selected, params: options[selected] };

            onConfirm({ key: target.key, binding: newBinding });
          }}
        >
          Confirm
        </Button>
      </DialogActions>
    </Dialog>
  );
};

export default BindingEditor;
