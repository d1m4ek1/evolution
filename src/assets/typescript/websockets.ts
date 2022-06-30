const OpenConnectWebSocket = (): void => {
  const websocket: WebSocket = new WebSocket('ws://localhost:8000/websocket/connect');

  websocket.addEventListener('open', (): void => {
    console.log('Conection opened');
  });
  websocket.addEventListener('close', (): void => {
    console.log('Conection closed');
  });
};

export default OpenConnectWebSocket;
