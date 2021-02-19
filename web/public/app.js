"use strict";

axios.defaults.withCredentials = true;

function getCookie(key) {
  var b = document.cookie.match("(^|;)\\s*" + key + "\\s*=\\s*([^;]+)");
  return b ? b.pop() : "";
}

function removeCookie(key) {
  document.cookie = `${key}=;expires=${new Date().toUTCString()};path=/`;
}

const App = () => {
  const userToken = getCookie("user_token"); // Set the user login for now based on false
  // @TODO check the cookie and set the state based on the cookie

  const [loggedIn, setLoggedIn] = React.useState(userToken !== "");
  const NullLoginError = {
    message: null
  }; // Using this to store and show the login error

  const [loginError, setLoginError] = React.useState(NullLoginError);
  const [loginRequest, setLoginRequest] = React.useState({
    email: "",
    password: ""
  });

  const handleLogin = e => {
    e.preventDefault(); // Every attempt reset the error state

    setLoginError(NullLoginError);
    axios.post("/login", loginRequest).then(response => {
      if (response.status == 204) {
        const cookie = getCookie("user_token");
        setLoggedIn(cookie != "");
      }
    }).catch(error => {
      if (error.response) {
        setLoginError(error.response.data);
      }
    });
  };

  const handleLogout = () => {
    axios.post("/auth/logout").then(response => {
      if (response.status === 204) {
        setLoggedIn(false);
      }
    }).catch(error => {
      console.log(error); // If the token is invalid, assume it was revoked and log the user out

      if (error.response.status === 401) {
        removeCookie("user_token");
        setLoggedIn(false);
      }
    });
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