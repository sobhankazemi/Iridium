import React from "react";
import ReactDOM from "react-dom/client";
import App from "./App.tsx";
import "./index.css";
import { Providers } from "./providers.tsx";

ReactDOM.createRoot(document.getElementById("root")!).render(
  <Providers>
    <React.StrictMode>
      <App title="Iridium" description="a front-end for iridium project" />
    </React.StrictMode>
  </Providers>
);
