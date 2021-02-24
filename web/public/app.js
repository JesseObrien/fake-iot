"use strict";

const App = () => {
  const userToken = localStorage.getItem("user_token");
  const [loggedIn, setLoggedIn] = React.useState(userToken !== null);
  const NullLoginError = {
    message: null
  }; // Using this to store and show the login error

  const [loginError, setLoginError] = React.useState(NullLoginError);
  const [loginRequest, setLoginRequest] = React.useState({
    email: "",
    password: ""
  });

  const handleLogin = async e => {
    e.preventDefault(); // Every attempt reset the error state

    setLoginError(NullLoginError);

    try {
      const response = await axios.post("/login", loginRequest);

      if (response.status == 200) {
        localStorage.setItem("user_token", response.data.access_token);
        localStorage.setItem("user_account_id", response.data.account_id);
        setLoggedIn(true);
      }
    } catch (err) {
      if (error.response) {
        setLoginError(error.response.data);
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
    setLoginRequest({ ...loginRequest,
      [name]: value
    });
  }; // If the user is logged in, render the dashboard instead


  if (loggedIn) {
    return /*#__PURE__*/React.createElement(Dashboard, {
      handleLogout: handleLogout
    });
  } // If the user is not logged in, render the login form


  return /*#__PURE__*/React.createElement("form", {
    className: "login-form",
    onSubmit: handleLogin
  }, /*#__PURE__*/React.createElement("h1", null, "Sign Into Your Account"), /*#__PURE__*/React.createElement("div", null, /*#__PURE__*/React.createElement("label", {
    for: "email"
  }, "Email Address"), /*#__PURE__*/React.createElement("input", {
    onChange: event => onInputChange("email", event.target.value),
    type: "email",
    id: "email",
    className: "field"
  })), /*#__PURE__*/React.createElement("div", null, /*#__PURE__*/React.createElement("label", {
    for: "password"
  }, "Password"), /*#__PURE__*/React.createElement("input", {
    onChange: event => onInputChange("password", event.target.value),
    type: "password",
    id: "password",
    className: "field"
  })), loginError.message && /*#__PURE__*/React.createElement("div", {
    className: "alert is-error"
  }, "Error: ", loginError.message, " "), /*#__PURE__*/React.createElement("button", {
    type: "submit",
    className: "button block"
  }, "Login to my Dashboard"));
};