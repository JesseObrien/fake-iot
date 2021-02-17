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
    return (this.loginCount / this.planLimit) * 100;
  }
}

const Dashboard = ({ handleLogout }) => {
  const [accountDetails, setAccountDetails] = React.useState(
    new AccountDetails()
  );

  React.useEffect(() => {
    // @TODO load this from the server
    setAccountDetails(new AccountDetails(100, 0, "Startup", 100));
  }, [accountDetails]);

  const [accountUpgraded, setAccountUpgraded] = React.useState(false);

  const handleAccountUpgrade = () => {
    console.log("account upgraded");
    // @TODO make the HTTP request to upgrade the account for the user
    setAccountUpgraded(true);
  };

  return (
    <>
      <header class="top-nav">
        <h1>User Management Dashboard</h1>
        <button onClick={handleLogout} class="button is-border">
          Logout
        </button>
      </header>

      {accountDetails.limitReached() && (
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
          {accountDetails.planName} Plan - ${accountDetails.planCost}/Month
        </header>

        <div class="plan-content">
          <div class="progress-bar">
            <div
              style={{ width: `${accountDetails.percentOfLimit()}%` }}
              class="progress-bar-usage"
            ></div>
          </div>

          <h3>
            Users: {accountDetails.loginCount}/{accountDetails.planLimit}
          </h3>
        </div>

        <footer>
          <button onClick={handleAccountUpgrade} class="button is-success">
            Upgrade to Enterprise Plan
          </button>
        </footer>
      </div>
    </>
  );
};
