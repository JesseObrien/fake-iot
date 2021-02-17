class AccountDetails {
  constructor(planLimit, loginCount, planName, planCost) {
    this.planLimit = planLimit;
    this.loginCount = loginCount;
    this.planName = planName;
    this.planCost = planCost;
  }

  limitReached() {
    return this.planLimit == this.loginCount;
  }

  percentOfLimit() {
    return this.loginCount / this.planLimit * 100;
  }

}

const Dashboard = ({
  handleLogout
}) => {
  const [accountDetails, setAccountDetails] = React.useState(new AccountDetails());
  React.useEffect(() => {
    // @TODO load this from the server
    setAccountDetails(new AccountDetails(100, 0, "Startup", 100));
  }, [accountDetails]);
  const [accountUpgraded, setAccountUpgraded] = React.useState(false);

  const handleAccountUpgrade = () => {
    console.log("account upgraded"); // @TODO make the HTTP request to upgrade the account for the user

    setAccountUpgraded(true);
  };

  return /*#__PURE__*/React.createElement(React.Fragment, null, /*#__PURE__*/React.createElement("header", {
    class: "top-nav"
  }, /*#__PURE__*/React.createElement("h1", null, "User Management Dashboard"), /*#__PURE__*/React.createElement("button", {
    onClick: handleLogout,
    class: "button is-border"
  }, "Logout")), accountDetails.limitReached() && /*#__PURE__*/React.createElement("div", {
    class: "alert is-error"
  }, "You have exceeded the maximum number of users for your account, please upgrade your plan to increaese the limit."), accountUpgraded && /*#__PURE__*/React.createElement("div", {
    class: "alert is-success"
  }, "Your account has been upgraded successfully!"), /*#__PURE__*/React.createElement("div", {
    class: "plan"
  }, /*#__PURE__*/React.createElement("header", null, accountDetails.planName, " Plan - $", accountDetails.planCost, "/Month"), /*#__PURE__*/React.createElement("div", {
    class: "plan-content"
  }, /*#__PURE__*/React.createElement("div", {
    class: "progress-bar"
  }, /*#__PURE__*/React.createElement("div", {
    style: {
      width: `${accountDetails.percentOfLimit()}%`
    },
    class: "progress-bar-usage"
  })), /*#__PURE__*/React.createElement("h3", null, "Users: ", accountDetails.loginCount, "/", accountDetails.planLimit)), /*#__PURE__*/React.createElement("footer", null, /*#__PURE__*/React.createElement("button", {
    onClick: handleAccountUpgrade,
    class: "button is-success"
  }, "Upgrade to Enterprise Plan"))));
};