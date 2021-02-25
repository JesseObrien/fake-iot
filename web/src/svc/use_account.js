import { useState, useCallback } from "react";
import axios from 'axios';

const DefaultAccountState = () => {
  return {id: '', plan_type: '', plan_limit: 0, login_count: 0};
};

const useAccount = () => {
  const [account, setAccount] = useState(DefaultAccountState());
  const [accountUpgraded, setAccountUpgraded] = useState(false);
  const planLimitReached = account.plan_limit == account.login_count;
  const loginPercentage = (account.login_count / account.plan_limit) * 100;
  const loginCountDisplay = `${account.login_count}/${account.plan_limit}`;

  let planCost = 0;

  if (account.plan_type === "standard") {
    planCost = 100;
  }

  if (account.plan_type === "enterprise") {
    planCost = 1000;
  }

  const updateAccount = useCallback((props) => {
    setAccount((oldAccount) => ({...oldAccount, ...props}));
  }, [account]);

  const handleUpgradeAccount = useCallback(async () => {
    try {
      const accountId = localStorage.getItem("user_account_id");

      const response = await axios.post(`/accounts/${accountId}/upgrade`);

      if (response.status === 200) {
        // Make the popup appear
        setAccountUpgraded(true);

        const data = response.data;
        updateAccount(data)

        // Get rid of the pop up after 4 seconds
        setInterval(() => {
          setAccountUpgraded(false);
        }, 4000);
      }
    } catch (err) {
      console.log(err);
    }
  }, [account]);

return {account, planLimitReached, loginPercentage, planCost, loginCountDisplay, accountUpgraded, updateAccount, handleUpgradeAccount};
};

export default useAccount;
