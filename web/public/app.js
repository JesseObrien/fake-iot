"use strict";

const App = () => {
  const [loggedIn, setLogin] = React.useState(true);

  const onLogin = () => {};

  if (loggedIn) {
    return /*#__PURE__*/React.createElement(Dashboard, null);
  }

  return /*#__PURE__*/React.createElement("form", {
    className: "login-form",
    method: "get",
    action: "/dashboard.html"
  }, /*#__PURE__*/React.createElement("h1", null, "Sign Into Your Account"), /*#__PURE__*/React.createElement("div", null, /*#__PURE__*/React.createElement("label", {
    for: "email"
  }, "Email Address"), /*#__PURE__*/React.createElement("input", {
    type: "email",
    id: "email",
    className: "field"
  })), /*#__PURE__*/React.createElement("div", null, /*#__PURE__*/React.createElement("label", {
    for: "password"
  }, "Password"), /*#__PURE__*/React.createElement("input", {
    type: "password",
    id: "password",
    className: "field"
  })), /*#__PURE__*/React.createElement("input", {
    type: "submit",
    value: "Login to my Dashboard",
    className: "button block"
  }));
};