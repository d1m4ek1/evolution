export const websocket: WebSocket = new WebSocket('ws://localhost:8000/websocket/connect', 'contact');

export const EventMessageSend = (mess: string): void => {
  websocket.send(mess)
}

const OpenConnectWebSocket = (): void => {

  websocket.addEventListener('open', (): void => {
    console.log('Conection opened');
  });

  websocket.addEventListener('close', (event): void => {
    console.log('Conection closed');
  });
};

export default OpenConnectWebSocket;
