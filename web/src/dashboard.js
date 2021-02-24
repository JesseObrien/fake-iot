import React, { createRef, useRef, useEffect, useState } from "react";

class AccountInfo {
  constructor(id, plan_limit, login_count, plan_type) {
    this.id = id;
    this.plan_limit = plan_limit;
    this.login_count = login_count;
    this.plan_type = plan_type;
  }

  limitReached() {
    return this.plan_limit == this.login_count;
  }

  loginPercentage() {
    return (this.login_count / this.plan_limit) * 100;
  }

  planCost() {
    if (this.plan_type === "standard") {
      return 100;
    }

    if (this.plan_type === "enterprise") {
      return 1000;
    }
  }
}

const Dashboard = ({ handleLogout }) => {
  const [account, setAccount] = useState(new AccountInfo());
  const accountRef = useRef();
  accountRef.current = account;

  const [accountUpgraded, setAccountUpgraded] = useState(false);

  const ws = useRef(null);

  const handleAccountUpgrade = async () => {
    try {
      const accountId = localStorage.getItem("user_account_id");

      const response = await axios.post(`/accounts/${accountId}/upgrade`);

      if (response.status === 200) {
        // Make the popup appear
        setAccountUpgraded(true);

        const data = response.data;

        upgradedAccount = new AccountInfo(
          data.id,
          data.plan_limit,
          data.login_count,
          data.plan_type
        );
        setAccount(upgradedAccount);
        accountRef.current = upgradedAccount;

        // Get rid of the pop up after 4 seconds
        setInterval(() => {
          setAccountUpgraded(false);
        }, 4000);
      }
    } catch (err) {
      console.log(err);
    }
  };

  useEffect(() => {
    let accountId = localStorage.getItem("user_account_id");
    let token = localStorage.getItem("user_token");
    let addr = window.location;

    if (!token) {
      return;
    }

    let uri = `wss://${addr.host}/accounts/${accountId}/updates`;
    ws.current = new WebSocket(uri);

    ws.current.onopen = () => {
      console.log("websocket connected");
      ws.current.send(
        JSON.stringify({
          operation: "account_updates_subscribe",
          token: `Bearer ${token}`,
        })
      );
    };

    ws.current.onclose = () => {
      console.log("websocket closed");
    };

    ws.current.onmessage = (message) => {
      const parsedMessage = JSON.parse(message.data);

      if (parsedMessage.operation === "authorization_failure") {
        handleLogout();
        return;
      }

      if (parsedMessage.operation === "account_info_response") {
        const data = JSON.parse(parsedMessage.data);
        const updatedAccount = new AccountInfo(
          data.id,
          data.plan_limit,
          data.login_count,
          data.plan_type
        );
        setAccount(updatedAccount);
        accountRef.current = updatedAccount;
      }

      if (parsedMessage.operation === "account_metrics_updated") {
        const data = JSON.parse(parsedMessage.data);
        const updatedAccount = new AccountInfo(
          accountRef.current.id,
          accountRef.current.plan_limit,
          data.login_count,
          accountRef.current.plan_type
        );
        setAccount(updatedAccount);
        accountRef.current = updatedAccount;
      }
    };
    () => {
      ws.current.close();
    };
  }, []);

  if (!account) {
    return (
      <>
        <div>Loading...</div>
      </>
    );
  }

  return (
    <>
      <header class="top-nav">
        <h1>User Management Dashboard</h1>
        <button onClick={handleLogout} class="button is-border">
          Logout
        </button>
      </header>

      {account.limitReached() && (
        <div class="alert is-error">
          You have exceeded the maximum number of users for your account, please
          upgrade your plan to increaese the limit.
        </div>
      )}

      {accountUpgraded && (
        <div class="alert is-success">
          Your account has been upgraded successfully!
        </div>
      )}

      <div class="plan">
        <header>
          {account.plan_type} - ${account.planCost()}/Month
        </header>

        <div class="plan-content">
          <div class="progress-bar">
            <div
              style={{ width: `${account.loginPercentage()}%` }}
              class="progress-bar-usage"
            ></div>
          </div>

          <h3>
            Users: {account.login_count}/{account.plan_limit}
          </h3>
        </div>

        <footer>
          {account.plan_type === "standard" && (
            <button onClick={handleAccountUpgrade} class="button is-success">
              Upgrade to Enterprise Plan
            </button>
          )}
        </footer>
      </div>
    </>
  );
};

export default Dashboard;
