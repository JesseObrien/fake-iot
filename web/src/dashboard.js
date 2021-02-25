import React, { useState } from "react";
import useAccount from "./svc/use_account";
import useAccountUpdates from "./svc/use_account_updates";
import { useSession } from './svc/use_session';

const Dashboard = () => {
  const { forceLogout } = useSession();
  const {account, planLimitReached, loginPercentage, planCost, loginCountDisplay, accountUpgraded, updateAccount, handleUpgradeAccount} = useAccount();
  useAccountUpdates(updateAccount, forceLogout);

  if (!account) {
    return (
      <>
        <div>Loading...</div>
      </>
    );
  }

  return (
    <>
      <header className="top-nav">
        <h1 data-testid="dashboard-title">User Management Dashboard</h1>
        <button onClick={forceLogout} className="button is-border">
          Logout
        </button>
      </header>

      {planLimitReached && (
        <div className="alert is-error">
          You have exceeded the maximum number of users for your account, please
          upgrade your plan to increase the limit.
        </div>
      )}

      {accountUpgraded && (
        <div className="alert is-success">
          Your account has been upgraded successfully!
        </div>
      )}

      <div className="plan">
        <header>
          {account.plan_type} - ${planCost}/Month
        </header>

        <div className="plan-content">
          <div className="progress-bar">
            <div
              style={{ width: `${loginPercentage}%` }}
              className="progress-bar-usage"
            ></div>
          </div>

          <h3>
            Users: {loginCountDisplay}
          </h3>
        </div>

        <footer>
          {account.plan_type === "standard" && (
            <button
              onClick={handleUpgradeAccount}
              className="button is-success"
            >
              Upgrade to Enterprise Plan
            </button>
          )}
        </footer>
      </div>
    </>
  );
};

export default Dashboard;
