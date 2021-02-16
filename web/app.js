"use strict";

const App = () => {
  // Set the user login for now based on false
  // @TODO check the cookie and set the state based on the cookie
  const [loggedIn, setLogin] = React.useState(false);

  //
  const [loginRequest, setLoginRequest] = React.useState({
    email: "",
    password: "",
  });

  const handleLogin = () => {
    // @TODO Handle the user login, send the loginRequest to the server, set a cookie and log the user in
    console.log("login requested with credentials:", loginRequest);
  };

  const onInputChange = (name, value) => {
    setLoginRequest({ ...loginRequest, [name]: value });
  };

  // If the user is logged in, render the dashboard instead
  if (loggedIn) {
    return <Dashboard />;
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

      <button onClick={handleLogin} className="button block">
        Login to my Dashboard
      </button>
    </div>
  );
};
