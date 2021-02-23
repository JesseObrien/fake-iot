"use strict";

const App = () => {
  const userToken = localStorage.getItem("user_token");

  // Set the user login for now based on false
  // @TODO check the cookie and set the state based on the cookie
  const [loggedIn, setLoggedIn] = React.useState(userToken !== null);

  const NullLoginError = { message: null };

  // Using this to store and show the login error
  const [loginError, setLoginError] = React.useState(NullLoginError);

  const [loginRequest, setLoginRequest] = React.useState({
    email: "",
    password: "",
  });

  const handleLogin = async (e) => {
    e.preventDefault();

    // Every attempt reset the error state
    setLoginError(NullLoginError);

    try {
      const response = await axios.post("/login", loginRequest);

      if (response.status == 200 && response.data?.access_token) {
        localStorage.setItem("user_token", response.data.access_token);
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

      localStorage.removeItem("user_token");
      setLoggedIn(false);
    } catch (err) {
      console.log(error);
    }
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

      {loginError.message && (
        <div className="alert is-error">Error: {loginError.message} </div>
      )}

      <button type="submit" className="button block">
        Login to my Dashboard
      </button>
    </form>
  );
};
