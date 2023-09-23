import { Box, Tab, TabProps, Tabs } from "@mui/material";
import { ComponentType, SyntheticEvent, useState } from "react";
import "./App.css";
import BindingEditor from "./binding-editor";
import Layout from "./layout";
import { Keymap, Binding, Layer } from "./keymap";
import {
  Link,
  useLoaderData,
  useRouteLoaderData,
  LinkProps,
} from "react-router-dom";
import { Keybinding } from "./binding-editor/binding-editor";

const LinkTab: ComponentType<TabProps & LinkProps> = Tab as React.ComponentType<
  TabProps & LinkProps
>;

const DeviceEditor = () => {
  const { keymap: loaderKeymap, name } = useRouteLoaderData("device") as {
    keymap: Keymap;
    name: string;
  };
  const [keymap, setKeymap] = useState(loaderKeymap);

  const { layer: layerIndex } = useLoaderData() as { layer: number };

  const layer = keymap.layers[layerIndex];

  const [editing, setEditing] = useState<Keybinding | undefined>();

  return (
    <Box>
      <Tabs value={layerIndex}>
        {keymap.layers.map((l, i) => (
          <LinkTab
            key={l.name}
            label={l.name}
            LinkComponent={Link}
            to={`/${name}/${i}`}
          />
        ))}
      </Tabs>

      {layer ? (
        <Layout bindings={layer.bindings} startEditing={setEditing} />
      ) : (
        <></>
      )}
      <BindingEditor
        keymap={keymap}
        open={Boolean(editing)}
        target={editing}
        onCancel={() => setEditing(undefined)}
        onConfirm={(newBinding) => {
          console.log("finishedEditing", newBinding.key, newBinding.binding);
          setEditing(undefined);

          const newKeymap: Keymap = {
            ...keymap,
            layers: replace(
              keymap.layers,
              layerIndex,
              replaceKey(layer, newBinding)
            ),
          };

          setKeymap(newKeymap);
        }}
      />
    </Box>
  );
};

export function replace<T>(arr: T[], index: number, replacement: T): T[] {
  if (index < 0 || index >= arr.length) return arr;

  const before = arr.slice(0, index);
  const after = arr.slice(index + 1);

  return before.concat([replacement]).concat(after);
}

const replaceKey = (layer: Layer, replacement: Keybinding): Layer => {
  return {
    ...layer,
    bindings: replace(layer.bindings, replacement.key, replacement.binding),
  };
};

export default DeviceEditor;
