const Dashboard = ({
  handleLogout
}) => {
  const [accountLimit, setAccountLimit] = React.useState(100);
  const [loginCount, setLoginCount] = React.useState(0);
  const loginPercent = loginCount / accountLimit * 100;
  const [accountMaxReached, setAccountMaxReached] = React.useState(false);
  const [accountUpgraded, setAccountUpgraded] = React.useState(false);

  const handleAccountUpgrade = () => {
    console.log("account upgraded");
    setAccountUpgraded(true);
  };

  return /*#__PURE__*/React.createElement(React.Fragment, null, /*#__PURE__*/React.createElement("header", {
    class: "top-nav"
  }, /*#__PURE__*/React.createElement("h1", null, "User Management Dashboard"), /*#__PURE__*/React.createElement("button", {
    onClick: handleLogout,
    class: "button is-border"
  }, "Logout")), accountMaxReached && /*#__PURE__*/React.createElement("div", {
    class: "alert is-error"
  }, "You have exceeded the maximum number of users for your account, please upgrade your plan to increaese the limit."), accountUpgraded && /*#__PURE__*/React.createElement("div", {
    class: "alert is-success"
  }, "Your account has been upgraded successfully!"), /*#__PURE__*/React.createElement("div", {
    class: "plan"
  }, /*#__PURE__*/React.createElement("header", null, "Startup Plan - $100/Month"), /*#__PURE__*/React.createElement("div", {
    class: "plan-content"
  }, /*#__PURE__*/React.createElement("div", {
    class: "progress-bar"
  }, /*#__PURE__*/React.createElement("div", {
    style: {
      width: `${loginPercent}%`
    },
    class: "progress-bar-usage"
  })), /*#__PURE__*/React.createElement("h3", null, "Users: ", loginCount, "/", accountLimit)), /*#__PURE__*/React.createElement("footer", null, /*#__PURE__*/React.createElement("button", {
    onClick: handleAccountUpgrade,
    class: "button is-success"
  }, "Upgrade to Enterprise Plan"))));
};