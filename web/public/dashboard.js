const Dashboard = () => {
  return /*#__PURE__*/React.createElement(React.Fragment, null, /*#__PURE__*/React.createElement("header", {
    class: "top-nav"
  }, /*#__PURE__*/React.createElement("h1", null, /*#__PURE__*/React.createElement("i", {
    class: "material-icons"
  }, "supervised_user_circle"), "User Management Dashboard"), /*#__PURE__*/React.createElement("button", {
    class: "button is-border"
  }, "Logout")), /*#__PURE__*/React.createElement("div", {
    class: "alert is-error"
  }, "You have exceeded the maximum number of users for your account, please upgrade your plan to increaese the limit."), /*#__PURE__*/React.createElement("div", {
    class: "alert is-success"
  }, "Your account has been upgraded successfully!"), /*#__PURE__*/React.createElement("div", {
    class: "plan"
  }, /*#__PURE__*/React.createElement("header", null, "Startup Plan - $100/Month"), /*#__PURE__*/React.createElement("div", {
    class: "plan-content"
  }, /*#__PURE__*/React.createElement("div", {
    class: "progress-bar"
  }, /*#__PURE__*/React.createElement("div", {
    style: {
      width: "35%"
    },
    class: "progress-bar-usage"
  })), /*#__PURE__*/React.createElement("h3", null, "Users: 35/100")), /*#__PURE__*/React.createElement("footer", null, /*#__PURE__*/React.createElement("button", {
    class: "button is-success"
  }, "Upgrade to Enterprise Plan"))));
};