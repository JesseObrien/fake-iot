"use strict";

axios.defaults.withCredentials = true;

function getCookie(key) {
  var b = document.cookie.match("(^|;)\\s*" + key + "\\s*=\\s*([^;]+)");
  return b ? b.pop() : "";
}

const App = () => {
  const userToken = getCookie("user_token");

  // Set the user login for now based on false
  // @TODO check the cookie and set the state based on the cookie
  const [loggedIn, setLoggedIn] = React.useState(userToken !== "");

  const NullLoginError = { error: null };

  // Using this to store and show the login error
  const [loginError, setLoginError] = React.useState(NullLoginError);

  const [loginRequest, setLoginRequest] = React.useState({
    email: "",
    password: "",
  });

  const handleLogin = () => {
    // Every attempt reset the error state
    setLoginError(NullLoginError);

    axios
      .post("/login", loginRequest)
      .then((response) => {
        if (response.status == 204) {
          const cookie = getCookie("user_token");
          setLoggedIn(cookie != "");
        }
      })
      .catch((error) => {
        if (error.response) {
          setLoginError(error.response.data);
        }
      });
  };

  const handleLogout = () => {
    axios
      .post("/logout")
      .then((response) => {
        if (response.status === 204) {
          setLoggedIn(false);
        }
      })
      .catch((error) => {
        console.log(error);
      });
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
    <div className="login-form">
      <h1>Sign Into Your Account</h1>

      <div>
        <label for="email">Email Address</label>
        <input
          onChange={(event) => onInputChange("email", event.target.value)}
          type="email"
          id="email"
          className="field"
        />
      </div>

      <div>
        <label for="password">Password</label>
        <input
          onChange={(event) => onInputChange("password", event.target.value)}
          type="password"
          id="password"
          className="field"
        />
      </div>

      {loginError.error && <p>Error: {loginError.error} </p>}

      <button onClick={handleLogin} className="button block">
        Login to my Dashboard
      </button>
    </div>
  );
};
