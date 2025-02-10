import axios from "axios";
import {jwtDecode} from "jwt-decode";

const base_url = "/";

export const login = async (username, password) => {
  const login_url = base_url + "login";

  try {
    const response = await axios.post(
      login_url,
      { username, password },
      { headers: { "Content-Type": "application/json" } }
    );
    console.log("RESPONSE RECEIVED: ", response.data);
    let refreshToken = response.data.token;
    localStorage.setItem("refreshToken", refreshToken);

    try {
      const decodedToken = jwtDecode(refreshToken);
      console.log(decodedToken);

      const { username, roles, iss, exp, iat } = decodedToken;
      localStorage.setItem("username", username);
      localStorage.setItem("roles", JSON.stringify(roles));
      localStorage.setItem("iss", iss);
      localStorage.setItem("exp", exp.toString());
      localStorage.setItem("iat", iat.toString());
    } catch (error) {
      console.error("Invalid token:", error);
      return {"success": false, "message": "Invalid Token"}
    }
    return {"success": true};
  } catch (err) {
    if (err.response) {
      console.error("Login Failed:", err.response.data);
      return {"success": false, "message": "Login Failed: " + (err.response.data?.message||"")}
    } else if (err.request) {
      console.error("No response received from server: ", err.request);
      return { success: false, message: "No response received from server"};
    } else {
      console.error("AXIOS ERROR:", err.message);
      return { success: false, message: "Unknown error"};
    }
  }
};

export default login;