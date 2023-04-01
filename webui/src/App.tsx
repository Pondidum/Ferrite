import { Box, Container, Tab, Tabs } from "@mui/material";
import { SyntheticEvent, useEffect, useState } from "react";
import "./App.css";
import BindingEditor from "./binding-editor/binding-editor";
import Keyboard from "./keyboard";

interface Zmk {
  layout: ZmkLayoutKey[];
}

export interface ZmkLayoutKey {
  Label: string;
  Row: number;
  Col: number;
  X: number;
  Y: number;
  R: number;
  Rx: number;
  Ry: number;
}

export interface Keymap {
  layers: KeymapLayer[];
}

export interface KeymapLayer {
  name: string;
  bindings: KeymapBinding[];
}

export interface KeymapBinding {
  type: string;
  first: string[];
  second: string[];
}

interface LayerSelection {
  index: number;
  layer: KeymapLayer | undefined;
}

const LayerEditor = ({
  layout,
  keymap,
  layer,
}: {
  layout: ZmkLayoutKey[];
  keymap: Keymap;
  layer: KeymapLayer | undefined;
}) => {
  if (!layer) {
    return <></>;
  }

  const [binding, editBinding] = useState<KeymapBinding | undefined>();

  return (
    <>
      <Keyboard layout={layout} layer={layer} editBinding={editBinding} />
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
  const [zmk, setZmk] = useState<Zmk>({ layout: [] });
  const [keymap, setKeymap] = useState<Keymap>({ layers: [] });

  const [layer, setLayer] = useState<LayerSelection>({
    index: 0,
    layer: undefined,
  });

  const selectLayer = (e: SyntheticEvent, newValue: number) => {
    if (keymap) {
      setLayer({
        index: newValue,
        layer: keymap.layers[newValue],
      });
    }
  };

  useEffect(() => {
    fetch("http://localhost:5656/api/zmk")
      .then((r) => r.json())
      .then((j) => setZmk(j));
  });

  useEffect(() => {
    fetch("http://localhost:5656/api/keymap")
      .then((r) => r.json())
      .then((j) => setKeymap(j));
  });

  return (
    <Container>
      <Box>
        <Tabs value={layer.index} onChange={selectLayer}>
          {keymap.layers.map((l) => (
            <Tab key={l.name} label={l.name} />
          ))}
        </Tabs>

        <LayerEditor layout={zmk.layout} keymap={keymap} layer={layer.layer} />
      </Box>
    </Container>
  );
}

export default App;
