import { ThemeProvider } from "@emotion/react";
import { createTheme, CssBaseline } from "@mui/material";
import React from "react";
import ReactDOM from "react-dom/client";
import App from "./App";
import "./index.css";
import { createBrowserRouter, RouterProvider } from "react-router-dom";
import { Keymap } from "./keymap";
import DeviceEditor from "./device";

const theme = createTheme({});
const router = createBrowserRouter([
  {
    element: <App />,
    children: [
      {
        path: "/:layer?",
        element: <DeviceEditor />,
        loader: async ({ params }) => {
          const response = await fetch("http://localhost:5656/api/device");
          const keymap = (await response.json()) as Keymap;

          const layer = Number.parseInt(params.layer as string);

          return { keymap, layer };
        },
      },
    ],
  },
]);
ReactDOM.createRoot(document.getElementById("root") as HTMLElement).render(
  <React.StrictMode>
    <ThemeProvider theme={theme}>
      <CssBaseline />
      <RouterProvider router={router} />
    </ThemeProvider>
  </React.StrictMode>
);
