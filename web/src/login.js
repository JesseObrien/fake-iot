import React, { useState } from "react";
import { useSession } from "./svc/use_session.js";

const LoginForm = () => {
  const { handleLogin, loginError } = useSession();

  const [loginRequest, setLoginRequest] = useState({
    email: "",
    password: "",
  });

  const onInputChange = (name, value) => {
    setLoginRequest({ ...loginRequest, [name]: value });
  };

  const handleLoginSubmit = (e) => {
    e.preventDefault();
    handleLogin(loginRequest);
  };

  return (
    <form className="login-form" onSubmit={handleLoginSubmit}>
      <h1 data-testid="login-form-title">Sign Into Your Account</h1>

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

export default LoginForm;
