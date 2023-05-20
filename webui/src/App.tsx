import { Box, Container, Tab, Tabs } from "@mui/material";
import { SyntheticEvent, useEffect, useState } from "react";
import "./App.css";
import BindingEditor from "./binding-editor";
import Keyboard from "./keyboard";
import { Keymap, Behavior, Layer, Device } from "./keymap";
import { ZmkProvider } from "./zmk/context";

interface LayerSelection {
  index: number | false;
  layer: Layer | undefined;
}

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

function App() {
  const [device, setDevice] = useState<Device>({ keymap: { layers: [] } });

  const [layer, setLayer] = useState<LayerSelection>({
    index: false,
    layer: undefined,
  });

  const selectLayer = (e: SyntheticEvent, newValue: number) => {
    if (device) {
      setLayer({
        index: newValue,
        layer: device.keymap.layers[newValue],
      });
    }
  };

  useEffect(() => {
    fetch("http://localhost:5656/api/device")
      .then((r) => r.json())
      .then((j) => setDevice(j));
  }, []);

  return (
    <ZmkProvider>
      <Container>
        <Box>
          <Tabs value={layer.index} onChange={selectLayer}>
            {device.keymap.layers.map((l) => (
              <Tab key={l.name} label={l.name} />
            ))}
          </Tabs>

          <LayerEditor keymap={device.keymap} layer={layer.layer} />
        </Box>
      </Container>
    </ZmkProvider>
  );
}

export default App;
