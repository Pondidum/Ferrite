import { Container } from "@mui/material";
import { Outlet } from "react-router-dom";
import { ZmkProvider } from "./zmk/context";
import "./App.css";

function App() {
  return (
    <ZmkProvider>
      <Container>
        <Outlet />
      </Container>
    </ZmkProvider>
  );
}

export default App;
