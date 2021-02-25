import React from "react";
import { useSession } from "./svc/use_session.js";
import Dashboard from "./dashboard.js";
import LoginForm from "./login.js";

const App = () => {
  const { loggedIn } = useSession();

  return (
    <>
      {!loggedIn && <LoginForm />}
      {loggedIn && <Dashboard />}
    </>
  );
};

export default App;
