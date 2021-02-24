import React from "react";
import axios from "axios";
import ReactDOM from "react-dom";
import App from "./app.js";

axios.interceptors.request.use((config) => {
  // Authorization Token
  const token = localStorage.getItem("user_token");
  if (token) {
    config.headers.Authorization = `Bearer ${token}`;
  }

  return config;
});

const AppContainer = document.querySelector("#app-container");

ReactDOM.render(<App />, AppContainer);
