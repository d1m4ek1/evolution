import GetCookie from "../getCookie";
import OpenConnectWebSocket from "../websockets";

const MODULE_CHECK_AUTHORIZE_USER = async () => {
  const token = GetCookie("token");
  const userId = GetCookie("userId");

  if (token !== undefined && userId !== undefined && token !== "" && userId !== "") {
    const response = await fetch(`/api/check_authorization?token=${token}&userId=${userId}`, {
      method: "GET",
    });
    const jsonResponse = await response.json();

    if (jsonResponse !== undefined) {
      if (jsonResponse.isVerify) {
        sessionStorage.setItem("verify", jsonResponse.isVerify);
        OpenConnectWebSocket();
      }
      return true;
    } else {
      return false;
    }
  } else {
    sessionStorage.setItem("verify", false);
    return false;
  }
};

export default MODULE_CHECK_AUTHORIZE_USER;
