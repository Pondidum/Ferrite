import { Box, Container, Tab, Tabs } from "@mui/material";
import { SyntheticEvent, useEffect, useState } from "react";
import "./App.css";
import BindingEditor from "./binding-editor";
import Keyboard from "./keyboard";
import { Keymap, KeymapBinding, KeymapLayer } from "./keymap";
import { Zmk } from "./zmk";

interface LayerSelection {
  index: number | false;
  layer: KeymapLayer | undefined;
}

const LayerEditor = ({
  zmk,
  keymap,
  layer,
}: {
  zmk: Zmk;
  keymap: Keymap;
  layer: KeymapLayer | undefined;
}) => {
  if (!layer) {
    return <></>;
  }

  const [binding, editBinding] = useState<KeymapBinding | undefined>();

  return (
    <>
      <Keyboard zmk={zmk} layer={layer} editBinding={editBinding} />
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
  const [zmk, setZmk] = useState<Zmk>({
    layout: [],
    symbols: {},
  });
  const [keymap, setKeymap] = useState<Keymap>({ layers: [] });

  const [layer, setLayer] = useState<LayerSelection>({
    index: false,
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
  }, []);

  useEffect(() => {
    fetch("http://localhost:5656/api/keymap")
      .then((r) => r.json())
      .then((j) => setKeymap(j));
  }, []);

  return (
    <Container>
      <Box>
        <Tabs value={layer.index} onChange={selectLayer}>
          {keymap.layers.map((l) => (
            <Tab key={l.name} label={l.name} />
          ))}
        </Tabs>

        <LayerEditor zmk={zmk} keymap={keymap} layer={layer.layer} />
      </Box>
    </Container>
  );
}

export default App;
