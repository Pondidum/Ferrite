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
        path: "/:device",
        id: "device",
        loader: async ({ params }) => {
          const name = params.device as string;
          const response = await fetch("http://localhost:5656/api/device");
          const keymap = (await response.json()) as Keymap;

          return { keymap, name };
        },
        children: [
          {
            path: ":layer?",
            element: <DeviceEditor />,
            id: "layer",
            loader: ({ params }) => {
              const layer = Number.parseInt(params.layer as string);
              return { layer };
            },
          },
        ],
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
