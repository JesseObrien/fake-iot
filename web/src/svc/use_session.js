import React, { createContext, useState, useContext, useCallback } from "react";
import axios from "axios";

const NullLoginError = { message: null };
const NullLogoutError = { message: null };

const sessionContext = createContext();

export const useSession = () => useContext(sessionContext);

export const SessionProvider = ({ children }) => {
  const session = useProvideSession();
  return (
    <sessionContext.Provider value={session}>
      {children}
    </sessionContext.Provider>
  );
};


export const useProvideSession = () => {
  const accountId = localStorage.getItem("user_account_id");
  const token = localStorage.getItem("user_token");
  const addr = window.location;

  const [loggedIn, setLoggedIn] = useState(token !== null);
  const [loginError, setLoginError] = useState(NullLoginError);
  const [logoutError, setLogoutError] = useState(NullLogoutError);

  const handleLogin = useCallback(
    async (loginRequest) => {
      // Every attempt reset the error state
      setLoginError(NullLoginError);

      try {
        const response = await axios.post("/login", loginRequest);

        if (response.status == 200) {
          localStorage.setItem("user_token", response.data.access_token);
          localStorage.setItem("user_account_id", response.data.account_id);
          setLoggedIn(true);
        }
      } catch (err) {
        console.log(err);
        if (err.response) {
          setLoginError(err.response.data);
        }
      }
    },
    [loggedIn, setLoggedIn]
  );

  const forceLogout = useCallback(async () => {
    try {
      const response = await axios.post("/auth/logout");
      if (response.status !== 200) {
        setLogoutError({message: "response state was not 200"})
      }
    } catch (err) {
      console.log(err);
      if (err.response) {
        setLogoutError(err.response.data);
      }
    }
    localStorage.removeItem("user_token");
    localStorage.removeItem("user_account_id");
    setLoggedIn(false);
  }, [loggedIn, forceLogout]);

  return {
    loggedIn,
    accountId,
    token,
    addr,
    loginError,
    logoutError,
    handleLogin,
    forceLogout,
  };
};
