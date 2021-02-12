"use strict";

const App = () => {
  const [loggedIn, setLogin] = React.useState(true);

  const onLogin = () => {};

  if (loggedIn) {
    return <Dashboard />;
  }

  return (
    <form className="login-form" method="get" action="/dashboard.html">
      <h1>Sign Into Your Account</h1>

      <div>
        <label for="email">Email Address</label>
        <input type="email" id="email" className="field" />
      </div>

      <div>
        <label for="password">Password</label>
        <input type="password" id="password" className="field" />
      </div>

      <input
        type="submit"
        value="Login to my Dashboard"
        className="button block"
      />
    </form>
  );
};
