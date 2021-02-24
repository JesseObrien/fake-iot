import React, { useEffect, useState } from "react";
import axios from "axios";
import Dashboard from "./dashboard.js";

const App = () => {
  const userToken = localStorage.getItem("user_token");

  const [loggedIn, setLoggedIn] = useState(userToken !== null);

  const NullLoginError = { message: null };

  // Using this to store and show the login error
  const [loginError, setLoginError] = useState(NullLoginError);

  const [loginRequest, setLoginRequest] = useState({
    email: "",
    password: "",
  });

  const handleLogin = async (e) => {
    e.preventDefault();

    // Every attempt reset the error state
    setLoginError(NullLoginError);

    try {
      const response = await axios.post("/login", loginRequest);

      if (response.status == 200) {
        localStorage.setItem("user_token", response.data.access_token);
        localStorage.setItem("user_account_id", response.data.account_id);
        setLoggedIn(true);
      }
    } catch (err) {
      console.log(err);
      if (err.response) {
        setLoginError(err.response.data);
      }
    }
  };

  const handleLogout = async () => {
    try {
      const response = await axios.post("/auth/logout");
    } catch (err) {
      console.log(err);
    }
    localStorage.removeItem("user_token");
    localStorage.removeItem("user_account_id");
    setLoggedIn(false);
  };

  const onInputChange = (name, value) => {
    setLoginRequest({ ...loginRequest, [name]: value });
  };

  // If the user is logged in, render the dashboard instead
  if (loggedIn) {
    return <Dashboard handleLogout={handleLogout} />;
  }

  // If the user is not logged in, render the login form
  return (
    <form className="login-form" onSubmit={handleLogin}>
      <h1>Sign Into Your Account</h1>

      <div>
        <label htmlFor="email">Email Address</label>
        <input
          onChange={(event) => onInputChange("email", event.target.value)}
          type="email"
          data-testid="email"
          id="email"
          className="field"
        />
      </div>

      <div>
        <label htmlFor="password">Password</label>
        <input
          onChange={(event) => onInputChange("password", event.target.value)}
          type="password"
          data-testid="password"
          id="password"
          className="field"
        />
      </div>

      {loginError.message && (
        <div className="alert is-error">Error: {loginError.message} </div>
      )}

      <button data-testid="login-button" type="submit" className="button block">
        Login to my Dashboard
      </button>
    </form>
  );
};

export default App;
