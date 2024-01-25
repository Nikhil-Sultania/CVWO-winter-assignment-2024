import { BrowserRouter, Routes, Route } from "react-router-dom";
import Home from "./Pages/home";
import About from "./Pages/About";
import PageNotFound from "./Pages/PageNotFound";
import React from "react";


function App(){
    return <>
        <h1> Please work</h1>
        <BrowserRouter>
            <Routes>
                <Route path="/" element={<Home />}>
                    <Route path="about" element={<About />} />
                    <Route path="*" element={<PageNotFound />} />
                </Route>
            </Routes>
        </BrowserRouter>
        <div>HI2</div>
    </>
}
export default App;