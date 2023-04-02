import { Box, Container, Tab, Tabs } from "@mui/material";
import { SyntheticEvent, useEffect, useState } from "react";
import "./App.css";
import BindingEditor from "./binding-editor";
import Keyboard from "./keyboard";
import { Keymap, KeymapBinding, KeymapLayer } from "./keymap";
import { ZmkProvider } from "./zmk/context";

interface LayerSelection {
  index: number | false;
  layer: KeymapLayer | undefined;
}

const LayerEditor = ({
  keymap,
  layer,
}: {
  keymap: Keymap;
  layer: KeymapLayer | undefined;
}) => {
  if (!layer) {
    return <></>;
  }

  const [binding, editBinding] = useState<KeymapBinding | undefined>();

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
    fetch("http://localhost:5656/api/keymap")
      .then((r) => r.json())
      .then((j) => setKeymap(j));
  }, []);

  return (
    <ZmkProvider>
      <Container>
        <Box>
          <Tabs value={layer.index} onChange={selectLayer}>
            {keymap.layers.map((l) => (
              <Tab key={l.name} label={l.name} />
            ))}
          </Tabs>

          <LayerEditor keymap={keymap} layer={layer.layer} />
        </Box>
      </Container>
    </ZmkProvider>
  );
}

export default App;
