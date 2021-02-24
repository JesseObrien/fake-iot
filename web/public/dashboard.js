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
    return this.login_count / this.plan_limit * 100;
  }

}

const Dashboard = ({
  handleLogout
}) => {
  const [account, setAccount] = React.useState(new AccountInfo());
  const [accountUpgraded, setAccountUpgraded] = React.useState(false);

  const handleAccountUpgrade = () => {
    console.log("account upgraded");
    setAccountUpgraded(true);
  };

  React.useEffect(() => {
    let accountId = localStorage.getItem("user_account_id");
    let token = localStorage.getItem("user_token");
    let addr = window.location;
    let uri = `wss://${addr.host}/accounts/${accountId}/updates`;
    let ws = new WebSocket(uri);

    ws.onopen = () => {
      ws.send(JSON.stringify({
        operation: "account_updates_subscribe",
        token: `Bearer ${token}`
      }));
    };

    ws.onmessage = message => {
      const parsedMessage = JSON.parse(message.data);

      if (parsedMessage.operation === "authorization_failure") {
        handleLogout();
        return;
      }

      if (parsedMessage.operation === "account_info_response") {
        const data = JSON.parse(parsedMessage.data);
        setAccount(new AccountInfo(data.id, data.plan_limit, data.login_count, data.plan_type));
      }

      if (parsedMessage.operation === "account_metrics_updated") {
        const data = JSON.parse(parsedMessage.data);
        setAccount(new AccountInfo(account.id, account.plan_limit, data.login_count, account.plan_type));
      }
    };

    return () => {
      ws.close();
    };
  }, []);

  if (!account) {
    return /*#__PURE__*/React.createElement(React.Fragment, null, /*#__PURE__*/React.createElement("div", null, "Loading..."));
  }

  return /*#__PURE__*/React.createElement(React.Fragment, null, /*#__PURE__*/React.createElement("header", {
    class: "top-nav"
  }, /*#__PURE__*/React.createElement("h1", null, "User Management Dashboard"), /*#__PURE__*/React.createElement("button", {
    onClick: handleLogout,
    class: "button is-border"
  }, "Logout")), account.limitReached() && /*#__PURE__*/React.createElement("div", {
    class: "alert is-error"
  }, "You have exceeded the maximum number of users for your account, please upgrade your plan to increaese the limit."), accountUpgraded && /*#__PURE__*/React.createElement("div", {
    class: "alert is-success"
  }, "Your account has been upgraded successfully!"), /*#__PURE__*/React.createElement("div", {
    class: "plan"
  }, /*#__PURE__*/React.createElement("header", null, account.plan_type, " - $100/Month"), /*#__PURE__*/React.createElement("div", {
    class: "plan-content"
  }, /*#__PURE__*/React.createElement("div", {
    class: "progress-bar"
  }, /*#__PURE__*/React.createElement("div", {
    style: {
      width: `${account.loginPercentage()}%`
    },
    class: "progress-bar-usage"
  })), /*#__PURE__*/React.createElement("h3", null, "Users: ", account.login_count, "/", account.plan_limit)), /*#__PURE__*/React.createElement("footer", null, /*#__PURE__*/React.createElement("button", {
    onClick: handleAccountUpgrade,
    class: "button is-success"
  }, "Upgrade to Enterprise Plan"))));
};