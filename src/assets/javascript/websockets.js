export let websocket;

const connetionToWebsocket = setInterval(() => {
  const getVerify = sessionStorage.getItem("verify");
  let isVerify = false;

  if (getVerify !== null) isVerify = JSON.parse(getVerify);

  if (isVerify) {
    const locationHost = window.location.host;
    websocket = new WebSocket(`ws://${locationHost}/websocket/connect`, "contact");

    clearInterval(connetionToWebsocket);
  }
}, 500);

export const EventMessageSend = (mess, websocket) => {
  if (websocket === undefined || websocket === null) return;

  websocket.send(mess);
};

const OpenConnectWebSocket = () => {
  if (websocket === undefined || websocket === null) return;

  websocket.addEventListener("open", () => {
    console.log("Conection opened");
  });

  websocket.addEventListener("close", (event) => {
    console.log("Conection closed");
  });
};

export default OpenConnectWebSocket;
