"use strict";

const App = () => {
  // Set the user login for now based on false
  // @TODO check the cookie and set the state based on the cookie
  const [loggedIn, setLogin] = React.useState(false); //

  const [loginRequest, setLoginRequest] = React.useState({
    email: "",
    password: ""
  });

  const handleLogin = () => {
    // @TODO Handle the user login, send the loginRequest to the server, set a cookie and log the user in
    console.log("login requested with credentials:", loginRequest);
  };

  const onInputChange = (name, value) => {
    setLoginRequest({ ...loginRequest,
      [name]: value
    });
  }; // If the user is logged in, render the dashboard instead


  if (loggedIn) {
    return /*#__PURE__*/React.createElement(Dashboard, null);
  } // If the user is not logged in, render the login form


  return /*#__PURE__*/React.createElement("div", {
    className: "login-form"
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
  })), /*#__PURE__*/React.createElement("button", {
    onClick: handleLogin,
    className: "button block"
  }, "Login to my Dashboard"));
};