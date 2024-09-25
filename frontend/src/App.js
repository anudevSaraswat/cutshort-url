import React from "react";
import CutshortURL from "./components/CutshortURL";
import "./App.css";

function App() {
  return (
    <div className="App">
      <header className="App-header">
        <h1>URL Shortener</h1>
      </header>
      <CutshortURL />
    </div>
  );
}

export default App;
