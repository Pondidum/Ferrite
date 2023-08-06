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

const paramOrDefault = (params: Parameter[], index: number): Parameter =>
  params.length > index ? params[index] : {};

export type Options = { [key: string]: Parameter[] };

const selectEditor = (
  keymap: Keymap,
  selected: Actions,
  options: Options,
  setOptions: (options: Options) => void
) => {
  const params = options[selected] ?? [];

  switch (selected) {
    case "kp":
      return (
        <EditorKP
          keymap={keymap}
          options={options}
          selected={selected}
          setOptions={setOptions}
        />
      );

    case "lt":
      return (
        <EditorLT
          keymap={keymap}
          options={options}
          selected={selected}
          setOptions={setOptions}
        />
      );

    case "mo":
      return (
        <EditorMO
          keymap={keymap}
          options={options}
          selected={selected}
          setOptions={setOptions}
        />
      );

    // case "mt":
    //   return (
    //     <>
    //       <KeyPicker param={paramOrDefault(params, 1)} setParam={(p) => {}} />
    //     </>
    //   );

    default:
      return <></>;
  }
};

// type BindingOptions = {
//   [key in Actions]: Parameter[];
// };

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
          onClick={() =>
            onConfirm({ action: selected, params: options[selected] })
          }
        >
          Confirm
        </Button>
      </DialogActions>
    </Dialog>
  );
};

export default BindingEditor;
