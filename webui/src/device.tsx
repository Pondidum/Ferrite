import { Box, LinkProps, Tab, TabProps, Tabs } from "@mui/material";
import { ComponentType, SyntheticEvent, useState } from "react";
import "./App.css";
import BindingEditor from "./binding-editor";
import Keyboard from "./keyboard";
import { Keymap, Behavior, Layer, Device } from "./keymap";
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

  const [binding, editBinding] = useState<Behavior | undefined>();

  return (
    <>
      <Keyboard layer={layer} editBinding={editBinding} />
      <BindingEditor
        open={Boolean(binding)}
        keymap={keymap}
        binding={binding}
        onCancel={() => editBinding(undefined)}
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

const DeviceEditor = () => {
  const { device, layer: layerIndex } = useLoaderData() as {
    device: Device;
    layer: number;
  };

  const layer = device.keymap.layers[layerIndex];

  return (
    <Box>
      <Tabs value={layerIndex}>
        {device.keymap.layers.map((l, i) => (
          <LinkTab
            key={l.name}
            label={l.name}
            LinkComponent={Link}
            to={"/" + i}
          />
        ))}
      </Tabs>

      <LayerEditor keymap={device.keymap} layer={layer} />
    </Box>
  );
};

export default DeviceEditor;
