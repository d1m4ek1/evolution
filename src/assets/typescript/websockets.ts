const websocket: WebSocket = new WebSocket('ws://localhost:8000/websocket/connect', 'echo-protocol');

export const EventMessage = (mess): void => {
  websocket.send(mess)
}

const OpenConnectWebSocket = (): void => {

  websocket.addEventListener('open', (): void => {
    console.log('Conection opened');
  });

  websocket.addEventListener('close', (): void => {
    console.log('Conection closed');
  });
};

export default OpenConnectWebSocket;
