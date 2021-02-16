const Dashboard = ({ handleLogout }) => {
  const [accountLimit, setAccountLimit] = React.useState(100);
  const [loginCount, setLoginCount] = React.useState(0);
  const loginPercent = (loginCount / accountLimit) * 100;

  const [accountMaxReached, setAccountMaxReached] = React.useState(false);
  const [accountUpgraded, setAccountUpgraded] = React.useState(false);

  const handleAccountUpgrade = () => {
    console.log("account upgraded");
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

      {accountMaxReached && (
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
        <header>Startup Plan - $100/Month</header>

        <div class="plan-content">
          <div class="progress-bar">
            <div
              style={{ width: `${loginPercent}%` }}
              class="progress-bar-usage"
            ></div>
          </div>

          <h3>
            Users: {loginCount}/{accountLimit}
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
