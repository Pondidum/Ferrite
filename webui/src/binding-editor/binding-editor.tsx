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
  Container,
  DialogContent,
  Grid,
} from "@mui/material";
import { Dispatch, SetStateAction, SyntheticEvent, useState } from "react";
import { Keymap, Binding, Parameter, Actions } from "../keymap";
import LayerPicker from "./layer-picker";
import KeyPicker from "./key-picker";
import ModifierPicker from "./modifier-grid";

const paramOrDefault = (params: Parameter[], index: number): Parameter =>
  params.length > index ? params[index] : {};

type Options = { [key: string]: Parameter[] };

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
        <KeyPicker
          param={paramOrDefault(params, 0)}
          setParam={(p) => {
            setOptions({
              ...options,
              [selected]: [p],
            });
          }}
        />
      );

    case "lt":
      return (
        <>
          <KeyPicker
            param={paramOrDefault(params, 1)}
            setParam={(p) => {
              setOptions({
                ...options,
                [selected]: [paramOrDefault(params, 0), p],
              });
            }}
          />
          <h3>When held, switch to layer</h3>

          <Grid container spacing={1} columns={4}>
            <Grid item xs={2}>
              <LayerPicker
                layers={keymap.layers}
                param={paramOrDefault(params, 0)}
                setParam={(p) => {
                  setOptions({
                    ...options,
                    [selected]: [p],
                  });
                }}
              />
            </Grid>
          </Grid>
        </>
      );

    case "mo":
      return (
        <LayerPicker
          layers={keymap.layers}
          param={paramOrDefault(params, 0)}
          setParam={(p) => {
            setOptions({
              ...options,
              [selected]: [p],
            });
          }}
        />
      );

    case "mt":
      return (
        <>
          <KeyPicker param={paramOrDefault(params, 1)} setParam={(p) => {}} />
        </>
      );

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
