import App from "./App";
import React from "react";
import {createRoot} from "react-dom/client";
import "./index.css";

const container = document.getElementById("root");
if(container) {
    const r = createRoot(container);
    r.render(<React.StrictMode>
            <App />
        </React.StrictMode>
    );
} else
    console.log("DOM is null");