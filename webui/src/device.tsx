import { Box, LinkProps, Tab, TabProps, Tabs } from "@mui/material";
import { ComponentType, SyntheticEvent, useState } from "react";
import "./App.css";
import BindingEditor from "./binding-editor";
import Keyboard from "./keyboard";
import { Keymap, Binding, Layer } from "./keymap";
import { Link, useLoaderData } from "react-router-dom";

const LayerEditor = ({
  keymap,
  layer,
}: {
  keymap: Keymap;
  layer: Layer | undefined;
}) => {
  if (!layer) {
    return <></>;
  }

  const [binding, editBinding] = useState<Binding | undefined>();

  return (
    <>
      <Keyboard layer={layer} editBinding={editBinding} />
      <BindingEditor
        open={Boolean(binding)}
        keymap={keymap}
        binding={binding}
        onCancel={() => {
          editBinding(undefined);
        }}
        onConfirm={(newBinding) => {
          console.log("confirm");
          console.log("old", binding);
          console.log("new", newBinding);
          editBinding(undefined);
        }}
      />
    </>
  );
};

const LinkTab: ComponentType<TabProps & LinkProps> = Tab as React.ComponentType<
  TabProps & LinkProps
>;

const KeymapEditor = () => {
  const { keymap, layer: layerIndex } = useLoaderData() as {
    keymap: Keymap;
    layer: number;
  };

  const layer = keymap.layers[layerIndex];

  return (
    <Box>
      <Tabs value={layerIndex}>
        {keymap.layers.map((l, i) => (
          <LinkTab
            key={l.name}
            label={l.name}
            LinkComponent={Link}
            to={"/" + i}
          />
        ))}
      </Tabs>

      <LayerEditor keymap={keymap} layer={layer} />
    </Box>
  );
};

export default KeymapEditor;
