import { useEffect } from "react";
import { useProvideSession } from "./use_session";

const useAccountUpdates = (handleUpdateAccount, forceLogout) => {
  const { accountId, token, addr } = useProvideSession();

  useEffect(() => {
    if (!token) {
      return;
    }

    let uri = `wss://${addr.host}/accounts/${accountId}/updates`;
    const ws = new WebSocket(uri);

    ws.onopen = () => {
      console.log("websocket connected");
      ws.send(
        JSON.stringify({
          operation: "account_updates_subscribe",
          token: token,
        })
      );
    };

    ws.onclose = () => {
      console.log("websocket closed");
    };

    ws.onmessage = (message) => {
      const parsedMessage = JSON.parse(message.data);

      if (parsedMessage.operation === "authorization_failure") {
        forceLogout();
        return;
      }

      if (parsedMessage.operation === "account_info_response") {
        const data = JSON.parse(parsedMessage.data);
        handleUpdateAccount(data);
      }

      if (parsedMessage.operation === "account_metrics_updated") {
        const data = JSON.parse(parsedMessage.data);

        handleUpdateAccount(data);
      }
    };
    () => {
      ws.close();
    };
  }, []);
};

export default useAccountUpdates;
